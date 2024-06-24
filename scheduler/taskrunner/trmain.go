package taskrunner

import (
	"errors"
	"log"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}
func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func Start() {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent:%v", i)
		}
		return nil
	}
	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received:%v", d)
			default:
				break forloop
			}
		}
		return errors.New("Excutor")
	}

	runner := NewRunner(30, false, d, e)

	worker := NewWorker(3, runner)
	worker.startWorker()
}
