// Package run implements an actor-runner with deterministic teardown. It is
// somewhat similar to package errgroup, except it does not require actor
// goroutines to understand context semantics. This makes it suitable for use in
// more circumstances; for example, goroutines which are handling connections
// from net.Listeners, or scanning input from a closable io.Reader.
package run

// Group collects actors (functions) and runs them concurrently.
// When one actor (function) returns, all actors are interrupted.
// The zero value of a Group is useful.
type Group struct {
	actors []actor
}

// Add an actor (function) to the group. Each actor must be pre-emptable by an
// interrupt function. That is, if interrupt is invoked, execute should return.
// Also, it must be safe to call interrupt even after execute has returned.
//
// The first actor (function) to return interrupts all running actors.
// The error is passed to the interrupt functions, and is returned by Run.
// 将各服务组件的启动入口通过Add方法添加到actors中。
// execute方法为所对应服务组件的启动入口方法；
// interrupt方法在组件退出时被调用，通常用于资源释放。
func (g *Group) Add(execute func() error, interrupt func(error)) {
	g.actors = append(g.actors, actor{execute, interrupt})
}

// Run all actors (functions) concurrently.
// When the first actor returns, all others are interrupted.
// Run only returns when all actors have exited.
// Run returns the error returned by the first exiting actor.
func (g *Group) Run() error {
	if len(g.actors) == 0 {
		return nil
	}

	// Run each actor.
	// 启动所有的actor
	errors := make(chan error, len(g.actors))
	for _, a := range g.actors {
		go func(a actor) {
			errors <- a.execute()
		}(a)
	}

	// Wait for the first actor to stop.
	// 等待第一个actor停止
	err := <-errors

	// Signal all actors to stop.
	// 向所有actor发送停止信号
	for _, a := range g.actors {
		a.interrupt(err)
	}

	// Wait for all actors to stop.
	// 等待所有的actor停止
	for i := 1; i < cap(errors); i++ {
		<-errors
	}

	// Return the original error.
	return err
}

type actor struct {
	execute   func() error
	interrupt func(error)
}
