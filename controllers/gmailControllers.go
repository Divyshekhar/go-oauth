package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/Divyshekhar/go-oauth/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var tok *oauth2.Token

func LoginHandler(c *gin.Context) {
	url := utils.GetURL()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallBackHandler(c *gin.Context) {
	code := c.Query("code")
	token := utils.Exchange(code)
	tok = token
	c.JSON(200, gin.H{"token generated": token, "refresh token": token.RefreshToken})
}

type EmailResult struct {
	Subject string
	Err     error
}

func fetchSubject(gmailsrvc *gmail.Service, user string, m *gmail.Message, ch chan<- EmailResult) {
	msg, err := gmailsrvc.Users.Messages.Get(user, m.Id).
		Format("metadata").
		MetadataHeaders("Subject").
		Do()

	if err != nil {
		ch <- EmailResult{"", err}
		return
	}

	for _, header := range msg.Payload.Headers {
		if header.Name == "Subject" {
			ch <- EmailResult{Subject: header.Value}
			return
		}
	}
	ch <- EmailResult{"(no subject)", nil}
}

func EmailHandler(c *gin.Context) {
	if tok == nil {
		c.JSON(402, gin.H{"error": "No token found"})
	}
	client := utils.OauthConfig.Client(context.Background(), tok)
	gmailsrvc, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		c.JSON(400, gin.H{"message": "Error creating a mail client"})
	}
	var subjects []string
	user := "me"
	msgs, err := gmailsrvc.Users.Messages.List(user).MaxResults(10).Do()
	if err != nil {
		c.JSON(200, gin.H{"message": "Error getting the messages"})
	}
	ch := make(chan EmailResult)
	for _, m := range msgs.Messages {
		go fetchSubject(gmailsrvc, user, m, ch)
	}
	for i := 0; i < len(msgs.Messages); i++ {
		result := <-ch
		if result.Err != nil {
			log.Printf("Error fetching subject: %v", result.Err)
			continue
		}
		subjects = append(subjects, result.Subject)
	}
	c.JSON(200, gin.H{"subjects": subjects})

}
