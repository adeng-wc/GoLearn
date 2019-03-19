package parser

import (
	"GoLearn/src/crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

const age = "41岁"
const Height = "155cm"
const Income = "月收入:3-5千"
const Marriage = "离异"

func TestParseProfile(t *testing.T) {
	bytes, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseProfile(bytes, "你的幸福是我的")

	//{[] [{  41岁 155cm  月收入:3-5千 离异 高中及以下 销售专员 工作地:阿坝松潘 魔羯座(12.22-01.19)  }]}

	fmt.Println(result)

	for _, m := range result.Items {
		profile := m.(model.Profile)

		if profile.Age != age {
			t.Errorf("profile.Age is %s, expected is %s", profile.Age, age)
		}

		if profile.Height != Height {
			t.Errorf("profile.Height is %s, expected is %s", profile.Height, age)
		}

		if profile.Income != Income {
			t.Errorf("profile.Income is %s, expected is %s", profile.Income, age)
		}


		if profile.Marriage != Marriage {
			t.Errorf("profile.Marriage is %s, expected is %s", profile.Marriage, age)
		}

	}
}
