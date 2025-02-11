package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
)

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

func ChatComplete() {
	apiKey := ""
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = "https://dashscope.aliyuncs.com/compatible-mode/v1"

	client := openai.NewClientWithConfig(cfg)
	ctx := context.TODO()

	prompt := `
	xxxx
`
	replace := map[string]string{
		"${pet.name}":       "野芙",
		"${pet.species}":    "兔子",
		"${pet.profession}": "荒野探险家",
		"${pet.extrainfo}":  "外向、大胆、无畏，喜欢自嘲，常说出其不意的话，打破常规，注重体验，不在意成败。有时冲动和疯狂，但常能带来有趣的结果。不仅打破常规，还通过与人辩论和互动获得新想法。她是一个社交的冒险家，既享受挑战大自然的艰辛，也从人与人之间的思想碰撞中寻找灵感。",
		"${pet.mbti}":       "ENTP",
		"${pet.wide_value}": "8",
	}

	//replace := map[string]string{
	//	"${pet.name}":       "PetInfo.Name",
	//	"${pet.species}":    "PetInfo.Species",
	//	"${pet.profession}": "PetInfo.Extra.Profession",
	//	"${pet.extrainfo}":  "PetInfo.Extra.ExtraInfo",
	//	"${pet.mbti}":       "PetInfo.Mbti",
	//	"${pet.wide_value}": "PetInfo.Extra.WideValue",
	//}

	for k, v := range replace {
		prompt = strings.ReplaceAll(prompt, k, v)
	}

	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "qwen-plus",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a creative story writer.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})
	if err != nil {
		var apiErr *openai.APIError
		ok := errors.As(err, &apiErr)
		if ok {
			fmt.Println("error as ok: ", apiErr.HTTPStatusCode, apiErr.Code)
		}
		fmt.Printf("failed to create chat completion, err: %v\n", err)
		return
	}

	fmt.Println(MarshalToString(resp))
	for idx, choice := range resp.Choices {
		fmt.Printf("index: %d, choice msg: \n%s\n", idx, choice.Message.Content)
	}
}
