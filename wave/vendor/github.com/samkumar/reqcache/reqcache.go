/*
 * Copyright (c) 2017 Sam Kumar <samkumar@berkeley.edu>
 * Copyright (c) 2017 University of California, Berkeley
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *     * Neither the name of the University of California, Berkeley nor the
 *       names of its contributors may be used to endorse or promote products
 *       derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNERS OR CONTRIBUTORS BE LIABLE FOR
 * ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

// Package reqcache provides an LRU cache with request management. See the
// Github page for more information.
package reqcache

import (
	"container/list"
	"context"
	"sync"
)

// LRUCache represents a Cache with an LRU eviction policy
type LRUCache struct {
	cache     map[interface{}]*LRUCacheEntry
	fetch     func(ctx context.Context, key interface{}) (interface{}, uint64, error)
	onEvict   func(evicted []*LRUCacheEntry)
	lruList   *list.List
	size      uint64
	capacity  uint64
	cacheLock *sync.Mutex
}

// LRUCacheEntry represents an entry in the LRU Cache. The size of this struct
// is the overhead of an entry existing in the cache. You should not have to
// actually use it.
type LRUCacheEntry struct {
	// Key is the key of the key-value pair represented by this entry.
	Key interface{}

	// Value is the value of the key-value pair represented by this entry.
	Value interface{}

	size uint64
	err  error

	/*
	 * Indicates whether or not there is a pending request for the Value. The
	 * cacheLock should be held when reading or writing this variable. If it is
	 * true, then the ready channel is open; otherwise, it is closed (or nil),
	 * if no request was made.
	 */
	pending bool

	/*
	 * Stores the number of goroutines waiting for the result of the pending
	 * fetch. If this number goes to zero, the fetch is cancelled. This is also
	 * protected by the cacheLock.
	 */
	waiting uint64

	/*
	 * Should be called to free the context when the fetch is done.
	 */
	fetchdone func()

	/*
	 * Signal to all waiters that the result is ready by closing this channel.
	 * Make sure you hold be the cacheLock before closing this channel!
	 */
	ready chan struct{}

	/*
	 * Stores the position of this entry in the LRU list.
	 */
	element *list.Element
}

// NewLRUCache returns a new instance of LRUCache.
//
// capacity is the capacity of the cache. If the sum of the sizes of elements in
// the cache exceeds the capacity, the least recently used elements are evicted
// from the cache.
//
// fetch is a function that is called on cache misses to fetch the element that
// is missing in the cache. The key that missed in the cache is passed as the
// argument. The function should return the corresponding value, and the size
// of the result (used to make sure that the total size does not exceed the
// cache's capacity). It can also return an error, in which case the result is
// not cached and the error is propagated to callers of Get(). No locks are
// held when fetch is called, so it is suitable to do blocking operations to
// fetch data. A context is also passed in, that is completely unrelated to the
// context passed to Get(). This context may be cancelled if no pending calls
// to get are interested in the result, which may happen if the contexts of all
// requesting goroutines time out, or if Put() is caled while the request is
// being fetched.
//
// onEvict is a function that is whenever elements are evicted from the cache
// according to the LRU replacement policy. It is called with the key-value
// pairs representing the evicted elements passed as arguments. It is not
// called with locks held, so it can perform blocking operations or even
// interact with this cache. It can be set to nil if the onEvict callback is
// not needed.
func NewLRUCache(capacity uint64, fetch func(ctx context.Context, key interface{}) (interface{}, uint64, error), onEvict func(evicted []*LRUCacheEntry)) *LRUCache {
	return &LRUCache{
		cache:     make(map[interface{}]*LRUCacheEntry),
		fetch:     fetch,
		onEvict:   onEvict,
		lruList:   list.New(),
		capacity:  capacity,
		cacheLock: &sync.Mutex{},
	}
}

// The cacheLock must be held when this function executes.
func (lruc *LRUCache) addEntryToLRU(entry *LRUCacheEntry) []*LRUCacheEntry {
	entry.element = lruc.lruList.PushFront(entry)
	lruc.size += entry.size
	return lruc.evictEntriesIfNecessary()
}

// The cacheLock and lruLock must both be held when this function executes.
func (lruc *LRUCache) evictEntriesIfNecessary() []*LRUCacheEntry {
	pruned := []*LRUCacheEntry{}
	for lruc.size > lruc.capacity {
		element := lruc.lruList.Back()
		lruc.lruList.Remove(element)
		entry := element.Value.(*LRUCacheEntry)
		delete(lruc.cache, entry.Key)
		lruc.size -= entry.size
		pruned = append(pruned, entry)
	}
	return pruned
}

// Calls the onEvict callback for a list of evicted entries. Should be called
// without the cacheLock acquired.
func (lruc *LRUCache) callOnEvict(evicted []*LRUCacheEntry) {
	if lruc.onEvict != nil {
		lruc.onEvict(evicted)
	}
}

// SetCapacity sets the capacity of the cache, evicting elements if necessary.
func (lruc *LRUCache) SetCapacity(capacity uint64) {
	lruc.cacheLock.Lock()
	lruc.capacity = capacity
	evicted := lruc.evictEntriesIfNecessary()
	lruc.cacheLock.Unlock()
	lruc.callOnEvict(evicted)
}

// CacheStatus describes the presence of an item in the cache.
type CacheStatus int

const (
	// Present status indicates that the specified item is in the cache.
	Present CacheStatus = iota

	// Pending status indicates that the specified item is not in the cache,
	// but that there is a pending request for the item.
	Pending CacheStatus = iota

	// Missing status indicates that the specified item is not in the cache,
	// and that there is no outstanding request for the item.
	Missing CacheStatus = iota
)

// TryGet checks if an element is present in the cache. The second return
// value describes the presence of the item in the cache; if it is Present,
// then provides the value corresponding to the provided key in the first
// return value.
func (lruc *LRUCache) TryGet(key interface{}) (interface{}, CacheStatus) {
	lruc.cacheLock.Lock()
	defer lruc.cacheLock.Unlock()

	value, ok := lruc.cache[key]
	if !ok {
		return value, Missing
	} else if value.pending {
		return value, Pending
	} else {
		return value, Present
	}

}

// Get returns the value corresponding to the specialized key, caching the
// result. Returns an error if and only if there was a cache miss and the
// provided fetch() function returned an error. If Put() is called while
// a fetch is blocking, then the result of the fetch is thrown away and the
// value specified by Put() is returned.
func (lruc *LRUCache) Get(ctx context.Context, key interface{}) (interface{}, error) {
	lruc.cacheLock.Lock()
	entry, ok := lruc.cache[key]
	if ok {
		/* Wait for the result if it's still pending. */
		if entry.pending {
			var err error
			entry.waiting++
			lruc.cacheLock.Unlock()
			select {
			case <-entry.ready:
				err = entry.err
			case <-ctx.Done():
				err = ctx.Err()
			}
			lruc.cacheLock.Lock()
			if entry.waiting--; entry.waiting == 0 {
				entry.fetchdone()
			}
			if err != nil {
				/* There was an error fetching this value. */
				lruc.cacheLock.Unlock()
				return nil, err
			}
		}

		/* Cache hit. */
		lruc.lruList.MoveToFront(entry.element)
		value := entry.Value
		lruc.cacheLock.Unlock()
		return value, nil
	}

	/* Cache miss. Create placeholder. */
	entry = &LRUCacheEntry{
		Key:     key,
		pending: true,
		waiting: 1,
		ready:   make(chan struct{}),
	}
	lruc.cache[key] = entry
	fetchctx, fetchcancel := context.WithCancel(context.Background())
	entry.fetchdone = fetchcancel
	lruc.cacheLock.Unlock()

	/* Fetch the value. */
	go func() {
		value, size, err := lruc.fetch(fetchctx, key)

		lruc.cacheLock.Lock()
		/*
		 * If the pending flag is no longer set, then someone called Put()
		 * meanwhile. We don't want to touch the cache or use the new value we
		 * got; instead, just use the value that was put there.
		 */
		if entry.pending {
			/* Check for and handle any error in fetching the value. */
			if err != nil {
				delete(lruc.cache, key)
				entry.err = err
				entry.pending = false
				close(entry.ready)
				lruc.cacheLock.Unlock()
				return
			}
			/* Store the result in the cache. */
			entry.Value = value
			entry.size = size
			entry.pending = false
			close(entry.ready)
			evicted := lruc.addEntryToLRU(entry)
			lruc.cacheLock.Unlock()

			lruc.callOnEvict(evicted)
		} else {
			lruc.cacheLock.Unlock()
		}
	}()

	select {
	case <-entry.ready:
		/*
		 * There are two cases. Someone may have called Put(), in which case
		 * we can go ahead and cancel the fetch. Or, the fetch could have
		 * completed. Either way, we return entry.Value.
		 */
		lruc.cacheLock.Lock()
		if entry.waiting--; entry.waiting == 0 {
			entry.fetchdone()
		}
		value := entry.Value
		lruc.cacheLock.Unlock()
		return value, nil
	case <-ctx.Done():
		/*
		 * This context expired, but other goroutines may be waiting for this
		 * result to appear. If none are, we can cancel the fetch.
		 */
		lruc.cacheLock.Lock()
		if entry.waiting--; entry.waiting == 0 {
			entry.fetchdone()
		}
		lruc.cacheLock.Unlock()
		return nil, ctx.Err()
	}
}

// Put an entry with a known value into the cache.
func (lruc *LRUCache) Put(key interface{}, value interface{}, size uint64) bool {
	lruc.cacheLock.Lock()
	entry, ok := lruc.cache[key]

	/* Check for case where it's already in the cache. */
	if ok {
		var evicted []*LRUCacheEntry
		entry.Value = value

		/* If the entry is still pending, wake up any waiting threads. */
		if entry.pending {
			entry.pending = false
			close(entry.ready)

			/* Add to the LRU list. */
			evicted = lruc.addEntryToLRU(entry)
		}
		lruc.cacheLock.Unlock()

		if evicted != nil {
			lruc.callOnEvict(evicted)
		}

		return true
	}

	/* Put it in the cache. */
	entry = &LRUCacheEntry{
		Key:     key,
		Value:   value,
		size:    size,
		pending: false,
		err:     nil,
		ready:   nil,
	}
	lruc.cache[key] = entry
	evicted := lruc.addEntryToLRU(entry)
	lruc.cacheLock.Unlock()

	lruc.callOnEvict(evicted)

	return false
}

// Evict an entry from the cache.
func (lruc *LRUCache) Evict(key interface{}) bool {
	lruc.cacheLock.Lock()
	defer lruc.cacheLock.Unlock()

	entry, ok := lruc.cache[key]
	if !ok {
		return false
	}
	lruc.lruList.Remove(entry.element)
	delete(lruc.cache, entry.Key)
	lruc.size -= entry.size
	return true
}

// Invalidate empties the cache, calling the onEvict callback as appropriate.
func (lruc *LRUCache) Invalidate() {
	lruc.cacheLock.Lock()
	oldCapacity := lruc.capacity

	lruc.capacity = 0
	entries := lruc.evictEntriesIfNecessary()

	lruc.capacity = oldCapacity
	lruc.cacheLock.Unlock()

	lruc.callOnEvict(entries)
}
