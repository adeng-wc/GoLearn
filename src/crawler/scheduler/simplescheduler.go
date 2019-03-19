package scheduler

import "GoLearn/src/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	// 通过 开 go routine 来解决循环登陆
	go func() {
		s.workerChan <- r
	}()
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}
