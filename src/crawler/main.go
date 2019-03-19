package main

import (
	"GoLearn/src/crawler/engine"
	"GoLearn/src/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}
