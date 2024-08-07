package mailer

import (
	"fmt"
	"net/smtp"
	"time"
	"lore-fetcher/internal/types"
)

type Mailer struct {
  FromMail string
  ToMail string
  Password string
}

func (mailer Mailer) SendResults(resultBuffer chan types.Patch){
  for {
    patch := <- resultBuffer    
    auth := smtp.PlainAuth("", mailer.FromMail, mailer.Password , "smtp.gmail.com")
    to := []string{mailer.ToMail}
    msg := []byte("Subject: " + "Tests for patch '" + patch.Title + "' completed.\r\n" +
      "\r\n" +
      patch.ResultString +
      "\r\n")
    err := smtp.SendMail("smtp.gmail.com:587", auth, mailer.FromMail, to, msg)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("[", time.Now(), "]: Sent result to", mailer.FromMail)
  }
}
