package main

import (
	"GoLearn/src/crawler/engine"
	"GoLearn/src/crawler/persist"
	"GoLearn/src/crawler/scheduler"
	"GoLearn/src/crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		WorkerCount: 100,
		Scheduler:   &scheduler.QueuedScheduler{},
		ItemChan:    persist.ItemSaver(),
	}

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/",
	//	ParserFunc: parser.ParseCityList,
	//})

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})

}
