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
	#背景#
    你是一个创意十足的故事创作者，现在在为一个虚拟的动物世界构建角色。这是一个名为“Zaisland”的动物世界，生活在这里的动物叫“崽崽“。崽崽们过着与人类相似的生活，拥有自己的职业、人际关系和社会活动，但不受人类世界规律的限制，有奇幻色彩。这里没有人类的存在。
    #目的#
    你需要根据崽崽的基础信息，展开想象，补充以下内容，让这个角色更加立体和有趣：
    工作职责：具体描述宠物的职业，包括它的职位名称、日常任务、主要责任以及必备技能。
    特殊能力：描述一种与宠物的种类或职业相匹配的独特或神奇能力。
    兴趣爱好：设想一个类似人类的兴趣爱好，描述宠物的偏好及习惯。
    趣味小知识：提供两个独特且不常见的趣味事实，这些可以让宠物的形象更加鲜明和亲切。
    性格弱点：指出两个能让宠物显得更为真实和有吸引力的性格缺陷或古怪之处。
    #崽崽基础信息#
    名称：${pet.name}
    种类：${pet.species}
    职业：${pet.profession}
    MBTI类型：${pet.mbti}
    天生反骨值：${pet.wide_value}
    其他信息：${pet.extrainfo}
    
    #创作指南#
    发挥创意：尽情发挥你的想象力，创造一个丰富多彩的幻想世界。
    保持一致：确保所有添加的信息与宠物的基础信息和MBTI性格类型相吻合。
    避免直接提及：在创作过程中，不要直接提到MBTI类型，而是通过具体的细节展现其特征。
    新增细节：加入一些基础信息中未曾提及的新元素。
    语言风格：保持语言的逻辑性和精炼性，避免冗余和不必要的描述。
    字数限制：最终输出的内容不得超过400字。
    #输出#
    工作职责:
    特殊能力:
    兴趣爱好:
    趣味小知识:
    性格弱点:
    一句话打招呼:
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
