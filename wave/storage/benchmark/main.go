package main

import (
  "fmt"
  "os"
  "strconv"
  "context"
  "crypto/rand"
  mathrand "math/rand"
  "time"

  "github.com/immesys/wave/iapi"
	"github.com/immesys/wave/storage/overlay"
	//"github.com/immesys/wave/storage/simplehttp"
  "wave/storage/simplehttp"
)

type Benchmark struct {
  nThreads  int64
  rps int64
  duration  int64
  updatePercent int64 
  loadSize  int64 // number of items to pre-load
  reqSize int64 // how may bytes per update request
  store []iapi.HashSchemeInstance
}
var storageconfig map[string]string

func NewBenchmark(threads, rps, duration, update, ls, rs int64) *Benchmark{
  bench := Benchmark{nThreads: threads,
                rps: rps,
                duration: duration,
                updatePercent: update,
                loadSize: ls, 
                reqSize: rs}
  return &bench
}

func init() {
  storageconfig = make(map[string]string)
  storageconfig["provider"] = "http_v1"
  storageconfig["url"] = "http://127.0.0.1:8080/v1"
  storageconfig["v1key"] = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEdgl41jCTp+Elxv1UzSoonFQxdd4d
bjaWQL5EYcghyIv3KRyR3PV2C87LHmk3FRWNbbfHlQZmJv7cWpJKNsjy3Q==
-----END PUBLIC KEY-----`
  storageconfig["v1auditors"] = "127.0.0.1:5001"
  cfg := make(map[string]map[string]string)
  cfg["default"] = storageconfig
  ov, err := overlay.NewOverlay(cfg)
  if err != nil {
    panic(err)
  }
  iapi.InjectStorageInterface(ov)
}

func getInstance() iapi.StorageDriverInterface {
  sh := &simplehttp.SimpleHTTPStorage{}
  sh.Initialize(context.Background(), "simplehttp", storageconfig)
  return sh

}

func (b *Benchmark) load () {
  b.store = make([]iapi.HashSchemeInstance, b.loadSize)
  in := getInstance()
  ctx := context.Background()

  for i:=0; i<int(b.loadSize); i++ {
    content := make([]byte, b.reqSize)
    rand.Read(content)
    hi, err := in.Put(ctx, content)
    if err != nil {
      panic(err)
    }
    b.store[i] = hi
  }
}

/*
func (b *Benchmark) load_test () {
  in := getInstance()
  ctx := context.Background()

  for i:=int(b.loadSize)-1; i>=0; i-- {
    _, err := in.Get(ctx, b.store[i])
    if err != nil {
      panic(err)
    }
  }
}
*/

func (b *Benchmark) worker(workChan, done, end chan bool) {
  in := getInstance()
  ctx := context.Background()

  for {
    select {
    case <- end:
      done <- true
      break
    case <- workChan:
      x := mathrand.Intn(100)
      st := time.Now()
      if x < int(b.updatePercent) {
        // put
        content := make([]byte, b.reqSize)
        rand.Read(content)
        if _, err := in.Put(ctx, content); err != nil {
          panic(err)
        }
        if (x % 20 == 0) {
          fmt.Printf("PUT latency %v\n", time.Since(st))
        }
      } else {
        // get
        if _, err := in.Get(ctx, b.store[mathrand.Intn(int(b.loadSize))]); err !=nil {
          panic(err)
        }
        if (x % 20 == 0) {
          fmt.Printf("GET latency %v\n", time.Since(st))
        }
      }
    }
  }
}

// args: threads rps duration update loadSize reqSize
func main() {
  t, _ := strconv.ParseInt(os.Args[1], 10, 64)
  r, _ := strconv.ParseInt(os.Args[2], 10, 64)
  d, _ := strconv.ParseInt(os.Args[3], 10, 64)
  u, _ := strconv.ParseInt(os.Args[4], 10, 64)
  l, _ := strconv.ParseInt(os.Args[5], 10, 64)
  rs, _ := strconv.ParseInt(os.Args[6], 10, 64)
  b := NewBenchmark(t, r, d, u, l, rs)
  fmt.Printf("== Workload %v ==\n", b)
  fmt.Printf("== Loading ==\n")
  start := time.Now()
  b.load()
  time.Sleep(2*time.Second)
  duration := time.Since(start)
  fmt.Printf("== Loaded %v items in %v ==\n", b.loadSize, duration)
  workChan := make(chan bool, t)
  done := make(chan bool, t)
  end := make(chan bool, t)
  for i := 0 ; i< int(t); i++ {
    go b.worker(workChan, done, end)
  }

  mathrand.Seed(time.Now().UnixNano())
  endTime := time.Now().Add(time.Duration(b.duration*1000*1000*1000))
  iv := 1000*1000*1000/b.rps
  next := time.Now()
  count := 0
  for time.Now().Before(endTime) {
    if (time.Now().Before(next)) {
      continue
    }
    workChan <- true
    count++
    next = time.Now().Add(time.Duration(iv))
  }

  fmt.Printf("== Completed %v requests in %v ==\n", count, b.duration)
  for i:=0; i<int(t); i++ {
    end <- true
  }
}
