package controller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sashabaranov/go-openai"

	"go-gpt/internal/model/gpt"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SocketHandler(c *gin.Context) {
	promtHistory := []openai.ChatCompletionMessage{}
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	var client = gpt.NewClient()

	for {
		var lock sync.Mutex

		//读取ws中的数据
		lock.Lock()
		defer lock.Unlock()
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		res, err := ChatGPT(client, string(message), promtHistory)
		if err != nil {
			fmt.Println(err)
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt, []byte(res.Choices[0].Message.Content))
		if err != nil {
			fmt.Println(err)
			break
		}
		promtHistory = append(promtHistory,
			openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: string(message),
			},
			openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: res.Choices[0].Message.Content,
			})

	}
}
