package serverutil_test

import (
	"context"
	"statusbay/serverutil"
	"sync"
	"testing"
	"time"
)

type ServeStruct struct {
	Request chan string
	isStop  bool
	init    bool
}

func MockServeStruct() *ServeStruct {
	return &ServeStruct{
		Request: make(chan string),
		isStop:  false,
		init:    false,
	}

}

func (m *ServeStruct) Serve(ctx context.Context, wg *sync.WaitGroup) {

	go func() {
		m.init = true
		for {
			select {
			case <-ctx.Done():
				m.isStop = true
				wg.Done()
				return
			}
		}
	}()

}
func TestServe(t *testing.T) {

	serverStruct := MockServeStruct()

	if serverStruct.init {
		t.Fatalf("unexpected serve init, got %t expected %t", serverStruct.isStop, false)
	}

	servers := []serverutil.Server{
		serverStruct,
	}
	ctx := context.Background()
	serverutil.RunAll(ctx, servers)
	time.Sleep(time.Second)
	if !serverStruct.init {
		t.Fatalf("unexpected init serve funtion, got %t expected %t", serverStruct.isStop, true)
	}

}
