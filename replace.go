package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/expr-lang/expr"
	"strings"
)

//type PetInfo struct {
//	Id       int64        `json:"id,omitempty"`
//	Name     string       `json:"name,omitempty"`
//	Species  string       `json:"species,omitempty"`
//	Mbti     string       `json:"mbti,omitempty"`
//	Extra    PetExtraInfo `json:"extra"`
//}
//
//type PetExtraInfo struct {
//	WideValue  int    `json:"wide_value,omitempty"`
//	Profession string `json:"profession,omitempty"`
//	ExtraInfo  string `json:"extra_info,omitempty"`
//}

func MarshalToString(v any) string {
	bs, _ := json.Marshal(v)
	return string(bs)
}

func UnMarshal(s string) map[string]string {
	var res map[string]string
	_ = json.Unmarshal([]byte(s), &res)
	return res
}

func Replace() {
	pet := PetInfo{
		Id:      123,
		Name:    "Buddy",
		Species: "Dog",
		Mbti:    "ENFJ",
		Extra: PetExtraInfo{
			WideValue:  85,
			Profession: "Guard",
			ExtraInfo:  "Loyal and Friendly",
		},
	}

	tmpl := `#崽崽基础信息#
	名称：${pet.name}
	种类：${pet.species}
	职业：${pet.profession}
	MBTI类型：${pet.mbti}
	天生反骨值：${pet.wide_value}
	其他信息：${pet.extrainfo}`

	replace := map[string]string{
		"${pet.name}":       "pet.Name",
		"${pet.species}":    "pet.Species",
		"${pet.profession}": "pet.Extra.Profession",
		"${pet.mbti}":       "pet.Mbti",
		"${pet.wide_value}": "pet.Extra.WideValue",
		"${pet.extrainfo}":  "pet.Extra.ExtraInfo",
	}

	replaceStr := MarshalToString(replace)
	fmt.Println(replaceStr)
	program, err := expr.Compile(replaceStr, expr.Env(PetInfo{}))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, pet)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
	fmt.Println(reflect.TypeOf(output))
	replace = UnMarshal(output.(string))

	for k, v := range replace {
		tmpl = strings.ReplaceAll(tmpl, k, v)
	}

	//fmt.Println(templateString)

	fmt.Println(tmpl)
}
