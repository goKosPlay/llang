package internel

import (
	"github.com/bregydoc/gtranslate"
	"github.com/gin-gonic/gin"
	"log"
)

type Message struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ConvertLang struct {
	ZhCn  string `json:"zh_cn"`
	En    string `json:"en"`
	Khmer string `json:"khmer"`
}

//启动
func Run() {
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	r.POST("/get_lang", func(ctx *gin.Context) {
		var langData ConvertLang
		var status int
		var message string

		keyword := ctx.PostForm("keyword")
		status = -1
		message = "not found data!!"
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered %s\n", err)
			}
		}()
		//多一个协程当有异常会捕捉 panic 资讯
		if keyword != "" {
			enLang, err := gtranslate.TranslateWithParams(
				keyword,
				gtranslate.TranslationParams{
					From: "zh-CN",
					To:   "en",
				},
			)
			if err != nil {
				panic(err)
			}
			khmerLang, err := gtranslate.TranslateWithParams(
				enLang,
				gtranslate.TranslationParams{
					From: "en",
					To:   "km",
				},
			)
			if err != nil {
				panic(err)
			}
			langData = ConvertLang{ZhCn: keyword, En: enLang, Khmer: khmerLang}
			status = 0
			message = "found current data!!"
		}
		ctx.JSON(200, Message{Status: status, Message: message, Data: langData})
	})
	r.Run(":8080")
}
