package parser

import (
	"GoLearn/src/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

/**
	城市列表
 */
func ParseCityList(bytes []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}

	// 控制 10个
	limit := 3
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
