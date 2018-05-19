package ants_test

import (
	"testing"
	"github.com/panjf2000/ants"
	"fmt"
	"runtime"
)

var n = 10

func demoFunc() {
	var n int
	for i := 0; i < 10000; i++ {
		n += i
	}
	fmt.Printf("finish task with result:%d\n", n)
}

func TestDefaultPool(t *testing.T) {
	for i := 0; i < n; i++ {
		ants.Push(demoFunc)
	}

	t.Logf("pool capacity:%d", ants.Cap())
	t.Logf("running workers number:%d", ants.Running())
	t.Logf("free workers number:%d", ants.Free())

	ants.Wait()

	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println("memory usage:", mem.TotalAlloc/1024)
}

//func TestCustomPool(t *testing.T) {
//	p := ants.NewPool(1000)
//	for i := 0; i < n; i++ {
//		p.Push(demoFunc)
//	}
//
//	t.Logf("pool capacity:%d", p.Cap())
//	t.Logf("running workers number:%d", p.Running())
//	t.Logf("free workers number:%d", p.Free())
//
//	mem := runtime.MemStats{}
//	runtime.ReadMemStats(&mem)
//
//}

func TestNoPool(t *testing.T) {
	for i := 0; i < n; i++ {
		go demoFunc()
	}
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Println("memory usage:", mem.TotalAlloc/1024)

}