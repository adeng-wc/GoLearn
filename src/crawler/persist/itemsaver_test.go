package persist

import (
	"GoLearn/src/crawler/model"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"testing"
)

func TestSave(t *testing.T) {

	expected := model.Profile{
		Name:       "1",
		Gender:     "2",
		Age:        "3",
		Height:     "4",
		Weight:     "5",
		Income:     "6",
		Marriage:   "7",  // 婚姻状况
		Education:  "8",  // 教育
		Occupation: "9",  // 工作
		Hukou:      "10", // 户口
		Xinzuo:     "11", // 星座
		House:      "12", // 房子
		Car:        "13", // 车子
	}

	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	// TODO : try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%+v \n", resp)
	t.Logf("%s \n", *resp.Source)

	/**
		A *json.RawMesasge is not a []byte. It's a pointer.
		Dereference the pointer to convert to a slice of bytes:
		byteSlice := []byte(*pointerToRawMessage)
	 */
	fmt.Println([]byte(*resp.Source))

	var actual model.Profile
	err = json.Unmarshal([]byte(*resp.Source), &actual)

	if err != nil {
		panic(err)
	}

	if actual != expected {
		t.Errorf("get %s not eaqueal %s", actual, expected)
	}
}
