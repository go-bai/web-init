package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed web
var f embed.FS

func setWeb(route *gin.Engine) {
	// 设置 web
	subFS, err := fs.Sub(f, "web")
	if err != nil {
		panic(err)
	}
	templ := template.Must(template.New("").ParseFS(subFS, "*.html"))
	route.SetHTMLTemplate(templ)
	route.StaticFS("/web/", http.FS(subFS))
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}
