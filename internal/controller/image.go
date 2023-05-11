package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"

	"go-gpt/internal/model/gpt"
)

func ImageBase64ChatGptHandler(c *gin.Context) {
	client := gpt.NewClient()
	reqBase64 := openai.ImageRequest{
		Prompt:         "Portrait of a humanoid parrot in a classic costume, high detail, realistic light, unreal engine",
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}
	respBase64, err := client.CreateImage(c.Request.Context(), reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return
	}
	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return
	}

	file, err := os.Create("image.png")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
		return
	}

	fmt.Println("The image was saved as example.png")
	c.JSON(200, nil)
}

func GetImageUrlChatGptHandler(c *gin.Context) {
	description := c.Query("description")
	if description == "" {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"message": "query description is not allow empty",
		})
	}
	client := gpt.NewClient()
	imgReq := openai.ImageRequest{
		Prompt:         description,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
		N:              1,
	}
	res, err := client.CreateImage(c.Request.Context(), imgReq)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(500, err)
		return
	}

	url := res.Data[0].URL
	c.JSON(200, map[string]interface{}{
		"url": url,
	})

}
