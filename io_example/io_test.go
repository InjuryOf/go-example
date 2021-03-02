package io_example

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestReadFrom(t *testing.T) {
	//从字符串读取
	data, err := ReadFrom(strings.NewReader("张三你在干啥"), 20)

	// 从文件中读取
	//file ,err := os.Open("test.txt")
	//data ,err := ReadFrom(file, 5)

	if err != io.EOF {
		t.Error(err)
	}
	t.Log(string(data))
}

func TestWriteFrom(t *testing.T) {
	// 将字符串写入到文件
	//file, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_WRONLY, 777)
	file, err := os.Open("test1.txt")
	defer file.Close()
	if err != nil{
		t.Log("file not exists")
	}
	n, _ := WriteFrom(file, "World")
	t.Log(n)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return r.width*2 + r.height*2
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float64 {
	return c.radius * math.Pi * 2
}

func mesure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterface(t *testing.T) {
	r := rect{width: 2, height: 3}
	c := circle{5}

	mesure(r)
	mesure(c)
}

func TestSelect(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string, 2)
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for {
		select {
		case message1 := <-c1:
			fmt.Printf("通道信息：%s\n", message1)
		case message2 := <-c2:
			fmt.Printf("通道信息：%s\n", message2)
			return
		case <-time.After(2 * time.Second):
			fmt.Println("timeout")
			return
		}
	}

}

func TestTimer(t *testing.T) {

	time1 := time.NewTimer(time.Second * 2)
	c1 := make(chan string)
	go func() {
		<-time1.C
		fmt.Println("timer event")
		c1 <- "done"
	}()
	<-c1
}

func TestTickLimit(t *testing.T) {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for request := range requests {
		<-limiter
		fmt.Println("request", request, time.Now())
	}
}

func TestAtomic(t *testing.T) {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddUint64(&ops, 1)
				//ops += 1
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops：",ops)
}
