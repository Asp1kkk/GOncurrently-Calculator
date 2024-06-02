package pool

import (
	"sync"
)

type Task struct {
	Expression string
}

func (t *Task) Execute() {

}

type Pool struct {
	numWorkers int
	tasks      chan Task
	wg         sync.WaitGroup
	isStart    bool
	isStop     bool
}

func (p *Pool) Start() {
	if p.isStart {
		return
	}

	p.wg.Add(p.numWorkers)
	for i := 0; i < p.numWorkers; i++ {
		go func() {
			defer p.wg.Done()

			for task := range p.tasks {
				task.Execute()
			}
		}()
	}

	p.isStart = true
}

func (p *Pool) Stop() {
	if p.isStop {
		return
	}

	go func() {
		close(p.tasks)
		p.wg.Wait()
	}()

	p.isStop = true
}

func (p *Pool) AddWork(task Task) {
	if !p.isStop {
		p.tasks <- task
	}
}

func NewWorkerPool(numWorkers int) *Pool {
	return &Pool{
		numWorkers: numWorkers,
		tasks:      make(chan Task),
	}
}
