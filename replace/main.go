package main

import (
	"fmt"
	"github.com/expr-lang/expr"
	"log"
)

// 定义结构体
type PetInfo struct {
	Id      int64        `json:"id,omitempty"`
	Name    string       `json:"name,omitempty"`
	Species string       `json:"species,omitempty"`
	Mbti    string       `json:"mbti,omitempty"`
	Extra   PetExtraInfo `json:"extra"`
}

type PetExtraInfo struct {
	WideValue  int    `json:"wide_value,omitempty"`
	Profession string `json:"profession,omitempty"`
	ExtraInfo  string `json:"extra_info,omitempty"`
}

func main() {
	exprStr := "PetInfo.Extra.ExtraInfo"
	// 创建 PetInfo 实例
	et := PetInfo{
		Name:    "Jaken",
		Species: "xsa",
		Extra: PetExtraInfo{
			WideValue:  0,
			Profession: "",
			ExtraInfo:  "ExExtraInfo",
		},
	}
	// 创建表达式语言的环境
	env := map[string]interface{}{
		"PetInfo": et,
	}

	program, err := expr.Compile(exprStr, expr.Env(env))
	if err != nil {
		log.Fatalf("Failed to compile expression %s: %v", exprStr, err)
	}

	result, err := expr.Run(program, env)
	if err != nil {
		log.Fatalf("Failed to execute expression %s: %v", exprStr, err)
	}

	fmt.Println(result)
}
