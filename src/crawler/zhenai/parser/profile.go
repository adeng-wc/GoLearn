package parser

import (
	"GoLearn/src/crawler/engine"
	"GoLearn/src/crawler/model"
	"regexp"
)

var re = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)

func ParseProfile(contents []byte, user string) engine.ParseResult {
	profile := extractString(contents, re)
	profile.Name = user
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, r *regexp.Regexp) model.Profile {
	matches := r.FindAllSubmatch(contents, -1)

	profile := model.Profile{}

	for i, m := range matches {
		switch i {
		case 0:
			profile.Marriage = string(m[1])
		case 1:
			profile.Age = string(m[1])
		case 2:
			profile.Xinzuo = string(m[1])
		case 3:
			profile.Height = string(m[1])
		case 4:
			profile.Hukou = string(m[1])
		case 5:
			profile.Income = string(m[1])
		case 6:
			profile.Occupation = string(m[1])
		case 7:
			profile.Education = string(m[1])
		}
	}

	return profile
}
