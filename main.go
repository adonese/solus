package main

import (
	"html/template"
)


func main(){
	r := getEngine()
	temp := template.Must(template.ParseFiles("./static/base.html", "./static/submit.html"))
	r.SetHTMLTemplate(temp)

	r.Static("/static", "./static")
	r.GET("/index", index)
	r.GET("/", pin)
	r.POST("/submit", pin)

	r.Run(":8001")
}
