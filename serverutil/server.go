package serverutil

import (
	"context"
	"sync"
)

// StopFunc is a server stop function, typically returned from Serve()
type StopFunc func()

// Server represents a component that serves stuff, and can be gracefully stopped.
type Server interface {
	Serve(ctx context.Context, wg *sync.WaitGroup)
}

// Runner is used to stop/start multiple servers
type Runner struct {
	stoppers []StopFunc
	wg       sync.WaitGroup
}

// RunAll will run all given servers and return a Runner instance
func RunAll(ctx context.Context, servers []Server) *Runner {
	r := &Runner{}
	for _, server := range servers {
		if server == nil {
			continue
		}
		server.Serve(ctx, &r.wg)
		r.wg.Add(1)
	}

	return r

}

// StopFunc will stop all registered servers in reverse order from how they were registered
func (r *Runner) StopFunc(cancelFn context.CancelFunc) {
	cancelFn()
	r.wg.Wait()
}
