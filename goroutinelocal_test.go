package goroutinelocal

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

var gllocal = NewGoroutineLocal(func() string {
	return "default"
})

func TestGoroutineLocal(t *testing.T) {
	gllocal.Set("test0")
	fmt.Println(runtime.GetGoroutineId(), gllocal.Get())

	go func() {
		gllocal.Set("test1")
		fmt.Println(runtime.GetGoroutineId(), gllocal.Get())
		gllocal.Remove()
		fmt.Println(runtime.GetGoroutineId(), gllocal.Get())
	}()
	time.Sleep(2 * time.Second)
	gllocal.Remove()
	fmt.Println("end", gllocal.GetMap())
}

var glruntime = runtime.NewGoroutineLocal(func() []byte {
	return make([]byte, 10*1024*1024)
})

var glruntime1 = runtime.NewGoroutineLocal(func() int {
	return 1985
})

func TestRuntimeGoroutineLocal(t *testing.T) {
	var gl = runtime.NewGoroutineLocal(func() string {
		return "default"
	})
	gl.Set("test0")
	fmt.Println(runtime.GetGoroutineId(), gl.Get())

	go func() {
		gl.Set("test1")
		fmt.Println(runtime.GetGoroutineId(), gl.Get())
		gl.Remove()
		fmt.Println(runtime.GetGoroutineId(), gl.Get())
	}()

	time.Sleep(2 * time.Second)
	gl.Remove()
	// fmt.Println("end", gl.GetMap())

}

func startAllocRuntime() {
	for i := 0; i < 1000; i++ {
		runtime.GC()
		go func() {
			glruntime1.Set(i)
			glruntime.Set(make([]byte, 10*1024*1024))
			// glruntime1.Remove()
			fmt.Println("GetGoroutineId", runtime.GetGoroutineId(), glruntime1.Get())
			time.Sleep(1 * time.Microsecond * 10)
		}()
		// time.Sleep(1 * time.Second)
	}
	fmt.Println("done")
}

func TestRuntimeGoroutineLocalLeak(t *testing.T) {

	var stats runtime.MemStats
	go func() {
		for {
			runtime.GC()
			debug.FreeOSMemory()
			runtime.ReadMemStats(&stats)
			fmt.Printf("HeapAlloc    = %d\n", stats.HeapAlloc)
			fmt.Printf("NumGoroutine = %d\n", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	startAllocRuntime()

	time.Sleep(10000 * time.Second)
}
