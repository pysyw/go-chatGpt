package controller

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"

	"go-gpt/internal/model/gpt"
	"go-gpt/pkg"
)

func WhisperGptHandler(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{
			"message": "获取数据失败",
		})
		return
	}
	fileTypes := map[string]struct{}{
		"audio/ogg":  {},
		"audio/mpeg": {},
		"audio/wav":  {},
	}
	_, ok := fileTypes[f.Header.Get("Content-Type")]
	if !ok {
		c.AbortWithStatusJSON(200, gin.H{
			"message": "文件类型不匹配,请上传audio类型文件",
		})
		return
	}
	ext := path.Ext(f.Filename)

	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	fileName := fileNameStr + ext
	p, err := pkg.MkdirIfNotExists("upload")
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	date := time.Now().Format("2006-01-02")
	fp := filepath.Join(p, "/", date, fileName)
	err = c.SaveUploadedFile(f, fp)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer os.Remove(fp)
	client := gpt.NewClient()
	res, err := client.CreateTranscription(c.Request.Context(),
		openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: fp,
		})
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, map[string]interface{}{
		"text": res.Text,
	})

}
