package main

import (
	"fmt"
	"runtime"
	"time"
)

// alloc big mem
// fast random method for mem[i]

func main() {
	var a, b, c, a_new, b_new, c_new, result, i uint64
	var data [536870912]uint64
	t0 := time.Now()
	for i = 0; i < 536870912; i++ {
		data[i] = i
	}
	fmt.Printf("4GB  Data Initialization Done in %v\n", time.Since(t0))
	a = 1
	b = 3
	c = 7
	runtime.GC()
	t := time.Now()
	for i = 0; i < 536870912; i++ {
		result = a + b*c + i%32
		//_ = data[result%8]
		//data[result%8] = data[result%8] + i
		//data[result%8] = irr
		result += data[result%8]
		a_new = b - a
		b_new = c - b
		c_new = c + a
		a = a_new
		b = b_new
		c = c_new
	}
	fmt.Printf("4GB Cache Random Read Done in %v -> latency*10 is %v\n", time.Since(t), time.Since(t)*10/536870912)
	a = 1
	b = 3
	c = 7
	runtime.GC()
	t2 := time.Now()
	//tmax := time.Since(time.Now())
	//t4 := tmax
	for i = 0; i < 536870912; i++ {
		result = a + b*c + i%32
		//_ = data[result%536870912]
		// data[result%536870912] = i
		// data[result%536870912] = data[result%536870912] + i
		//t3 := time.Now()
		result += data[result%536870912]
		//t4 = time.Since(t3)
		//if t4 > tmax {
		//	tmax = t4
		//}
		a_new = b - a
		b_new = c - b
		c_new = c + a
		a = a_new
		b = b_new
		c = c_new
	}
	fmt.Printf("4GB  DRAM Random Read Done in %v -> latency    is %v\n", time.Since(t2), time.Since(t2)/536870912)
	//fmt.Println(tmax, t4, result)
	fmt.Println(result)
}
