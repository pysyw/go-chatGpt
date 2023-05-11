package gpt

import (
	"github.com/sashabaranov/go-openai"

	"go-gpt/config"
)

type client *openai.Client

func NewClient() *openai.Client {
	return openai.NewClient(config.GLOBAL_CONFIG.API_KEY)
}
