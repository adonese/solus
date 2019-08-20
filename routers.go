package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

func index(c *gin.Context) {
	// Just serve a template
	c.HTML(200, "base.html", gin.H{"Title": "My First template",
		"Content": "GO is for the win!"})
}

// IPin used to Bind request body against it
type IPin struct {
	Ipin      string `form:"ipin" binding:"required"`
	UUID      string `form:"uuid" binding:"required"`
	PublicKey string `form:"pubkey" binding:"required"`
}

func pin(c *gin.Context) {
	if c.Request.Method == "POST" {
		var ipin IPin
		err := c.ShouldBindWith(&ipin, binding.Form)
		if err != nil {
			log.Printf("The error i: %v\n", err)
			c.HTML(400, "submit.html", gin.H{"error": err.Error()})
			return
		}

		// Now, let's go to the func
		pinBlock, err := rsaEncrypt(ipin.PublicKey, ipin.Ipin, ipin.UUID)
		if err != nil {
			log.Printf("The error i: %v\n", err)
			c.HTML(400, "submit.html", gin.H{"error": err.Error()})
			return
		}
		c.HTML(200, "submit.html", gin.H{"pin": pinBlock})
	} else {
		c.HTML(200, "submit.html", gin.H{"data": "Fill the form"})
	}
}
