package main

import (
	"GoLearn/src/crawler/engine"
	"GoLearn/src/crawler/scheduler"
	"GoLearn/src/crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		WorkerCount: 100,
		Scheduler: &scheduler.SimpleScheduler{

		},
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}
