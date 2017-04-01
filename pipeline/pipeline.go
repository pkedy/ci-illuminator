package pipeline

import "sync"

type Signal int

const (
	Next Signal = iota
	Finished
)

type Pipe interface {
	Process(chan Signal)
}

type Pipeline struct {
	repeat bool
	pipes  []Pipe
}

type Execution struct {
	sig      chan Signal
	wait     *sync.WaitGroup
	running  bool
	position int
}

func New(repeat bool, pipes ...Pipe) *Pipeline {
	return &Pipeline{
		repeat: repeat,
		pipes:  pipes,
	}
}

func (p *Pipeline) Start() *Execution {
	signal := make(chan Signal, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	e := &Execution{
		sig:      signal,
		running:  true,
		position: 0,
		wait:     &wg,
	}
	go func() {
		for e.running {
			pipe := p.pipes[e.position]
			pipe.Process(signal)

			s := <-e.sig
			switch s {
			case Finished:
				e.running = false
			case Next:
				e.position = (e.position + 1) % len(p.pipes)
				if !p.repeat && e.position == 0 {
					e.running = false
				}
			}
		}
		e.wait.Done()
	}()

	return e
}

func (e *Execution) Stop(wait bool) {
	e.sig <- Finished
	if wait {
		e.wait.Done()
	}
}
