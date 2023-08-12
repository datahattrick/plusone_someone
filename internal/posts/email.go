package posts

import (
	"bytes"
	"context"
	"net/smtp"

	"github.com/datahattrick/plusone_someone/internal/utils"
	"github.com/gofiber/fiber/v2/log"
)

func SendEmailToPostsRecipient(post Post) {
	cfg := &utils.Config{}
	cl, err := smtp.Dial(cfg.SMTP.Host)
	if err != nil {
		log.Errorf("Failed to connect to %s: %s", cfg.SMTP.Host, err)
	}
	c := context.Background()
	author, _ := utils.Database.DB.GetUserById(c, post.AuthorID)
	user, _ := utils.Database.DB.GetUserById(c, post.UserID)

	if cl != nil {
		defer cl.Close()
		cl.Mail(author.Email)
		cl.Rcpt(user.Email)

		wc, err := cl.Data()
		if err != nil {
			log.Errorf("Something went wrong: %s", err)
		}
		defer wc.Close()
		buf := bytes.NewBufferString(post.Message)
		if _, err = buf.WriteTo(wc); err != nil {
			log.Errorf("Something went wrong: %s", err)
		}
	}

}
