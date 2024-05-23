package mailer

import (
  "net/smtp"
  "fmt"
)

type Mailer struct {
  FromMail string
  ToMail string
  Password string
}

func (mailer Mailer) SendResults(resultBuffer chan string){
  for {
    result := <- resultBuffer    
    fmt.Println(result)
    auth := smtp.PlainAuth("", "marcelomspessoto@gmail.com", mailer.Password, "smtp.gmail.com")
    to := []string{"marcelomspessoto@gmail.com"}
    msg := []byte(result)
    err := smtp.SendMail("smtp.gmail.com:587", auth, "marcelomspessoto@gmail.com", to, msg)
    if err != nil {
      fmt.Println(err)
    }
  }
}
