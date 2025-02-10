package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

const (
	QWenApiKey  = ""
	QWenBaseUrl = "https://dashscope.aliyuncs.com/compatible-mode/v1"
)

const (
	StoryWriterPrompt = "You are a creative story writer."
)

type ModelControlOption struct {
	Temperature float32 `json:"temperature,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	OutputJson  bool    `json:"output_json"`
}

func GetQWenConfig() openai.ClientConfig {
	cfg := openai.DefaultConfig(QWenApiKey)
	cfg.BaseURL = QWenBaseUrl

	return cfg
}

func QuickToMessages(systemPrompt string, prompt string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	}
}

func HandleError(err error) {
	var apiErr *openai.APIError
	if errors.As(err, &apiErr) {
		metricsStr := fmt.Sprintf("%d_%v", apiErr.HTTPStatusCode, apiErr.Code)
		println(metricsStr)
		return
	}

}

func QWenChatComplete(ctx context.Context, prompt string, opt *ModelControlOption) (string, error) {
	messages := QuickToMessages(StoryWriterPrompt, prompt)
	req := openai.ChatCompletionRequest{Messages: messages}
	if opt != nil {
		req.Temperature = opt.Temperature
		req.MaxTokens = opt.MaxTokens
		if opt.OutputJson {
			req.ResponseFormat = &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			}
		}
	}

	cli := openai.NewClientWithConfig(GetQWenConfig())
	var err error
	defer func() {
		if err != nil {
			HandleError(err)
		}
	}()

	// log
	resp, err := cli.CreateChatCompletion(ctx, req)
	// log
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		err = errors.New("choices length is zero")
		return "", err
	}

	// metrics token use
	return resp.Choices[0].Message.Content, nil
}

//func CallChatComplete(ctx context.Context, req openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error) {
//	// log
//
//	cli := GetClient(GetQWenConfig())
//	resp, err = cli.CreateChatCompletion(ctx, req)
//	// log
//	return
//}
