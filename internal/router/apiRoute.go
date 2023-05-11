package router

import (
	"github.com/gin-gonic/gin"

	"go-gpt/internal/controller"
)

func initRoute(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/socket", controller.SocketHandler)
		v1.POST("/chat", controller.ChatController)
		// // 生成图片保存在本地, 目前暂时写死 Prompt
		v1.GET("/base/img", controller.ImageBase64ChatGptHandler)
		// // 生成在线图片返回url
		v1.GET("/img", controller.GetImageUrlChatGptHandler)
		// // 语音识别
		v1.POST("/whisper", controller.WhisperGptHandler)
	}
}
