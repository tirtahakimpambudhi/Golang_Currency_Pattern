package schedule

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Schedule struct {
	workers int
	msgC chan struct{}
	signalC chan os.Signal
	wg sync.WaitGroup
}

func NewScheduler(workers, buffers int) *Schedule {
	return &Schedule{
		workers: workers,
		msgC: make(chan struct{},buffers),
		signalC: make(chan os.Signal,1),
	}
}

func (s *Schedule) ListenAndWork()  {
		go func() {
			signal.Notify(s.signalC,syscall.SIGTERM)
		for {
			<-s.signalC
			s.msgC <- struct{}{}
			}
	}()
	s.wg.Add(s.workers)
	for i := 0; i < s.workers; i++ {
		i := i
		go func() {
			for {
				select {
				case _,open := <- s.msgC:
					if !open {
						fmt.Printf("worker %d <- closing \n",i)
						s.wg.Done()
						return
					}
					fmt.Printf("worker %d <- processing \n",i)
				}
			}
		}()
	}
}

func (s *Schedule) Exit()  {
	close(s.msgC)
	s.wg.Wait()
}

func WaitToExit() <-chan struct{} {
	runC := make(chan struct{})
	sc := make(chan os.Signal,1)
	signal.Notify(sc,os.Interrupt)
	go func() {
		defer close(runC)
		<- sc
	}()
	return runC
}