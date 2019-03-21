package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d:%v", itemCount, item)

			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver error: saving item %v: %v", item, err)
			}
		}
	}()

	return out
}

func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		return "", err
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	// %+v 会把结构体字段输出
	fmt.Printf("%+v", resp)

	return resp.Id, nil
}
