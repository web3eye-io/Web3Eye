package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

type IB interface {
	start(cancel *context.CancelFunc)
	index1(ctx context.Context)
	index2(ctx context.Context)
	stop()
}
type A struct {
	cancel *context.CancelFunc
	IB
}

func (a A) StartIndex() {
	ctx, cancel := context.WithCancel(context.Background())
	a.cancel = &cancel
	a.start(&cancel)
	fmt.Println("aaa ", reflect.ValueOf(a.cancel).IsNil())

	go a.index1(ctx)
	go a.index2(ctx)
	fmt.Println("aaa ", reflect.ValueOf(a.cancel).IsNil())

}

func (a A) StopIndex() {
	if a.cancel != nil {
		fmt.Println("A stop")
		(*a.cancel)()
		a.cancel = nil
	}
}

type B struct {
	cancel *context.CancelFunc
}

func (b B) start(cancel *context.CancelFunc) {
	fmt.Println(reflect.ValueOf(b.cancel).IsNil())
	b.cancel = cancel
	fmt.Println(reflect.ValueOf(b.cancel).IsNil())
}

func (b B) stop() {
	fmt.Println(reflect.ValueOf(b.cancel).IsNil())

	if !reflect.ValueOf(b.cancel).IsNil() {
		fmt.Println("B stop")
		(*b.cancel)()
		b.cancel = nil
	}
}

func (b B) index1(ctx context.Context) {
	for i := 0; i < 8; i++ {
		select {
		case <-time.NewTicker(time.Second * 1).C:
			fmt.Println("index1 ", i)
		case <-ctx.Done():
		}
	}
}

func (b B) index2(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-time.NewTicker(time.Second * 1).C:
			fmt.Println("index2 ", i)
		case <-ctx.Done():
		}
	}
	fmt.Println("sssss")
	b.stop()
}

func main() {
	a := A{IB: B{}}
	a.StartIndex()
	time.Sleep(10 * time.Second)
}
