package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/gogo/protobuf/proto"
	"github.com/google/trillian"
	"github.com/google/trillian/types"
	"github.com/immesys/wave/iapi"
	"github.com/immesys/wave/storage/simplehttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const logServer = "localhost:8092"
const mapServer = "localhost:8090"

var PublicKey string
var PrivateKey string
var PrivateKeyUnpacked *ecdsa.PrivateKey
var TreeID_Op int64
var TreeID_Root int64
var TreeID_Map int64

type PromiseObject struct {
	Promise *simplehttp.MergePromise
	Key     []byte
	Value   []byte
}

var promisemu sync.Mutex
var promises map[string]*PromiseObject

func initstorage() {
	optreeid := os.Getenv("VLDM_TREE_OPERATIONS")
	roottreeid := os.Getenv("VLDM_TREE_MAPROOTS")
	maptreeid := os.Getenv("VLDM_TREE_MAP")

	if optreeid == "" || roottreeid == "" || maptreeid == "" {
		fmt.Printf("Missing tree ids\n")
		os.Exit(1)
	}

	v, err := strconv.ParseInt(optreeid, 10, 64)
	if err != nil {
		fmt.Printf("could not parse tree id:%v\n", err)
	}
	TreeID_Op = v

	v, err = strconv.ParseInt(roottreeid, 10, 64)
	if err != nil {
		fmt.Printf("could not parse tree id:%v\n", err)
	}
	TreeID_Root = v

	v, err = strconv.ParseInt(maptreeid, 10, 64)
	if err != nil {
		fmt.Printf("could not parse tree id:%v\n", err)
	}
	TreeID_Map = v

	pub, err := ioutil.ReadFile("vldm_public.pem")
	if err != nil {
		fmt.Printf("could not read public key: %v\n", err)
		os.Exit(1)
	}
	priv, err := ioutil.ReadFile("vldm_private.pem")
	if err != nil {
		fmt.Printf("could not read private key: %v\n", err)
		os.Exit(1)
	}
	PublicKey = string(pub)
	PrivateKey = string(priv)
	pk, err := ParsePrivateKey(priv)
	if err != nil {
		fmt.Printf("could not parse private key\n", err)
	}
	PrivateKeyUnpacked = pk
	queueindexes = make(map[string]int64)
	promises = make(map[string]*PromiseObject)
}

type GetMapKeyResponse struct {
	Unmerged bool
	//If unmerged
	MergePromise *simplehttp.MergePromise
	//If merged
	SignedMapRoot []byte
	MapInclusion  []byte
	SignedLogRoot []byte
	LogInclusion  []byte
	Value         []byte
}

func GetMapKeyAtRev(key []byte, rev int64) (bool, *GetMapKeyResponse) {
	resp2, err := vmap.GetLeavesByRevision(context.Background(), &trillian.GetMapLeavesByRevisionRequest{
		MapId:    TreeID_Map,
		Index:    [][]byte{key},
		Revision: rev,
	})
	if err != nil {
		panic(err)
	}
	if resp2.MapLeafInclusion[0].Leaf.LeafValue == nil {
		//check for promise rather
		promisemu.Lock()
		pv, ok := promises[string(key)]
		promisemu.Unlock()
		if ok {
			return true, &GetMapKeyResponse{
				Unmerged:     true,
				MergePromise: pv.Promise,
				Value:        pv.Value,
			}
		}
	}

	smrbytes, err := proto.Marshal(resp2.MapRoot)
	if err != nil {
		panic(err)
	}
	h := sha256.New()
	h.Write([]byte{0})
	h.Write(smrbytes)
	r := h.Sum(nil)
	var rootloginclusion *trillian.GetInclusionProofByHashResponse

	llr, err := logclient.GetLatestSignedLogRoot(context.Background(), &trillian.GetLatestSignedLogRootRequest{
		LogId: TreeID_Root,
	})
	if err != nil {
		panic(err)
	}
	rootloginclusion, err = logclient.GetInclusionProofByHash(context.Background(), &trillian.GetInclusionProofByHashRequest{
		LogId:    TreeID_Root,
		LeafHash: r,
		TreeSize: llr.SignedLogRoot.TreeSize,
	})
	if err != nil {
		return false, nil
		/*
			if grpc.Code(err) == codes.NotFound {
				return false, nil
			}
			panic(err)*/
	}
	slrbytes, err := proto.Marshal(rootloginclusion.SignedLogRoot)
	if err != nil {
		panic(err)
	}
	slrinc, err := proto.Marshal(rootloginclusion.Proof[0])
	if err != nil {
		panic(err)
	}
	inclusion, err := proto.Marshal(resp2.MapLeafInclusion[0])
	if err != nil {
		panic(err)
	}
	smr, err := proto.Marshal(resp2.MapRoot)
	if err != nil {
		panic(err)
	}

	if resp2.MapLeafInclusion[0].Leaf.LeafValue == nil {
		return true, &GetMapKeyResponse{
			SignedMapRoot: smr,
			MapInclusion:  inclusion,
			SignedLogRoot: slrbytes,
			LogInclusion:  slrinc,
		}
	} else {
		// pv := &PromiseObject{}
		// fmt.Printf("leaf value is: %q", resp2.MapLeafInclusion[0].Leaf.LeafValue)
		// err := json.Unmarshal(resp2.MapLeafInclusion[0].Leaf.LeafValue, pv)
		// if err != nil {
		// 	panic(err)
		// }
		return true, &GetMapKeyResponse{
			SignedMapRoot: smr,
			MapInclusion:  inclusion,
			Value:         resp2.MapLeafInclusion[0].Leaf.LeafValue,
		}
	}
}

var goodrev int64

func GetMapKeyValue(key []byte) *GetMapKeyResponse {
	resp2, err := vmap.GetLeaves(context.Background(), &trillian.GetMapLeavesRequest{
		MapId: TreeID_Map,
		Index: [][]byte{key},
	})
	if err != nil {
		panic(err)
	}
	if resp2.MapLeafInclusion[0].Leaf.LeafValue == nil {
		//check for promise rather
		promisemu.Lock()
		pv, ok := promises[string(key)]
		promisemu.Unlock()
		if ok {
			return &GetMapKeyResponse{
				Unmerged:     true,
				MergePromise: pv.Promise,
				Value:        pv.Value,
			}
		}
	}

	smrbytes, err := proto.Marshal(resp2.MapRoot)
	if err != nil {
		panic(err)
	}
	h := sha256.New()
	h.Write([]byte{0})
	h.Write(smrbytes)
	r := h.Sum(nil)
	var rootloginclusion *trillian.GetInclusionProofByHashResponse

	llr, err := logclient.GetLatestSignedLogRoot(context.Background(), &trillian.GetLatestSignedLogRootRequest{
		LogId: TreeID_Root,
	})
	if err != nil {
		panic(err)
	}
	rootloginclusion, err = logclient.GetInclusionProofByHash(context.Background(), &trillian.GetInclusionProofByHashRequest{
		LogId:    TreeID_Root,
		LeafHash: r,
		TreeSize: llr.SignedLogRoot.TreeSize,
	})
	if err != nil {
		if grpc.Code(err) == codes.NotFound || grpc.Code(err) == codes.InvalidArgument {
			//fmt.Printf("finding previous map: map smr is not in log yet\n")
			var mr types.MapRootV1
			err := mr.UnmarshalBinary(resp2.MapRoot.MapRoot)
			if err != nil {
				panic(err)
			}
			rev := atomic.LoadInt64(&goodrev)
			for {
				rootok, resp := GetMapKeyAtRev(key, int64(rev))
				if rootok {
					return resp
				}
				rev--
				if rev < 0 {
					panic("could not find good log\n")
				}
			}
		} else {
			panic(err)
		}
	} else {
		var mr types.MapRootV1
		err := mr.UnmarshalBinary(resp2.MapRoot.MapRoot)
		if err != nil {
			panic(err)
		}
		atomic.StoreInt64(&goodrev, int64(mr.Revision))
	}

	slrbytes, err := proto.Marshal(rootloginclusion.SignedLogRoot)
	if err != nil {
		panic(err)
	}
	slrinc, err := proto.Marshal(rootloginclusion.Proof[0])
	if err != nil {
		panic(err)
	}
	inclusion, err := proto.Marshal(resp2.MapLeafInclusion[0])
	if err != nil {
		panic(err)
	}
	smr, err := proto.Marshal(resp2.MapRoot)
	if err != nil {
		panic(err)
	}

	if resp2.MapLeafInclusion[0].Leaf.LeafValue == nil {
		return &GetMapKeyResponse{
			SignedMapRoot: smr,
			MapInclusion:  inclusion,
			SignedLogRoot: slrbytes,
			LogInclusion:  slrinc,
		}
	} else {
		// pv := &PromiseObject{}
		// fmt.Printf("leaf value is: %q", resp2.MapLeafInclusion[0].Leaf.LeafValue)
		// err := json.Unmarshal(resp2.MapLeafInclusion[0].Leaf.LeafValue, pv)
		// if err != nil {
		// 	panic(err)
		// }
		return &GetMapKeyResponse{
			SignedMapRoot: smr,
			MapInclusion:  inclusion,
			Value:         resp2.MapLeafInclusion[0].Leaf.LeafValue,
		}
	}
}

func InsertKeyValue(key []byte, value []byte) *simplehttp.MergePromise {
	//TODO check for existing merge promise for this value
	hi := iapi.KECCAK256.Instance(value)
	hasharr := hi.Value()
	mp, err := MakeMergePromise(key, hasharr, PrivateKeyUnpacked)
	if err != nil {
		fmt.Printf("failed to make promise: %v\n", err)
		os.Exit(1)
	}
	po := &PromiseObject{
		Promise: mp,
		Key:     key,
		Value:   value,
	}
	promisemu.Lock()
	promises[string(key)] = po
	promisemu.Unlock()
	poj, err := json.Marshal(po)
	if err != nil {
		fmt.Printf("failed to marshal promise object: %v\n", err)
	}
	ctx := context.Background()
	llf := &trillian.LogLeaf{
		LeafValue: poj,
	}
	_, err = logclient.QueueLeaf(ctx, &trillian.QueueLeafRequest{
		LogId: TreeID_Op,
		Leaf:  llf,
	})
	if err != nil {
		panic(err)
	}
	return mp
}

func Exists(queueid []byte, index int64) bool {
	tohash := make([]byte, 40)
	copy(tohash[:32], queueid)
	binary.LittleEndian.PutUint64(tohash[32:], uint64(index+1))
	hi := iapi.KECCAK256.Instance(tohash)
	hiarr := hi.Value()

	resp2, err := vmap.GetLeaves(context.Background(), &trillian.GetMapLeavesRequest{
		MapId: TreeID_Map,
		Index: [][]byte{hiarr},
	})
	if err != nil {
		panic(err)
	}
	return resp2.MapLeafInclusion[0].Leaf.LeafValue != nil
}

var queueindexes map[string]int64
var queuemu sync.Mutex

func GetIndex(queueid []byte) int64 {
	queuemu.Lock()
	idx, ok := queueindexes[string(queueid)]
	queuemu.Unlock()
	if ok {
		return idx
	}
	if !Exists(queueid, 0) {
		return -1
	}
	start := int64(128)
	for Exists(queueid, start) {
		start <<= 4
	}
	pivot := start / 2
	interval := start / 4
	for {
		if Exists(queueid, pivot) {
			if interval == 0 {
				break
			}
			pivot += interval
			interval /= 2
		} else {
			if interval == 0 {
				pivot--
				break
			}
			pivot -= interval
			interval /= 2
		}
	}
	queuemu.Lock()
	queueindexes[string(queueid)] = pivot
	queuemu.Unlock()
	return pivot
}
func SetIndex(queueid []byte, id int64) {
	queuemu.Lock()
	queueindexes[string(queueid)] = id
	queuemu.Unlock()
}
