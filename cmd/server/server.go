package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wneessen/go-mail"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("web/src/templates/*")

	router.Static("/assets", "./web/out")

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Test uwu",
		})
	})

	router.GET("/test/mail", func(ctx *gin.Context) {
		m := mail.NewMsg()
		if err := m.From("test.tester@example.com"); err != nil {
			log.Fatalf("failed to set From address: %s", err)
		}
		if err := m.To("tina.tester@example.com"); err != nil {
			log.Fatalf("failed to set To address: %s", err)
		}
		m.Subject("Testing")
		m.SetBodyString(mail.TypeTextPlain, "test email")

		c, err := mail.NewClient(
			"127.0.0.1",
			mail.WithPort(1025),
			mail.WithSMTPAuth(mail.SMTPAuthPlain),
			mail.WithUsername("My_test"),
			mail.WithTLSConfig(&tls.Config{
				InsecureSkipVerify: true,
			}),
			mail.WithPassword("my_password"),
		)

		if err != nil {
			log.Fatalf("failed to create client: %s", err)
		}

		if err := c.DialAndSend(m); err != nil {
			log.Fatalf("failed to send mail: %s", err)
		}

		log.Printf("test")
		ctx.Status(http.StatusOK)
	})

	router.Run(":8080")
}
