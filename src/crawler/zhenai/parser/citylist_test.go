package parser

import (
	"io/ioutil"
	"testing"
)

/**
	城市列表
 */
func TestParseCityList(t *testing.T) {
	bytes, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(bytes)

	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result.Requests len is %d, %d", len(result.Requests), resultSize)
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result.Items len is %d, %d", len(result.Items), resultSize)
	}

	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s;but was %s", i, city, result.Items[i].(string))
		}
	}

}
