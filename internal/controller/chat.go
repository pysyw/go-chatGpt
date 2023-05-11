package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"

	"go-gpt/config"
)

type chatQuery struct {
	Message string `binding:"required" json:"message"`
}

func ChatController(g *gin.Context) {
	var m chatQuery
	err := g.ShouldBindJSON(&m)
	if err != nil {
		g.AbortWithError(400, err)
		return
	}
	client := openai.NewClient(config.GLOBAL_CONFIG.API_KEY)

	res, err := ChatGPT(client, m.Message, []openai.ChatCompletionMessage{})
	if err != nil {
		g.AbortWithError(500, err)
		return
	}
	g.JSON(200, res.Choices[0].Message.Content)
}

func ChatGPT(client *openai.Client, s string, h []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: s,
		},
	}

	messages = append(messages, h...)

	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}
	return res, nil
}
