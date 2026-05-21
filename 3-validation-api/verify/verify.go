package verify

import (
	"email/api/configs"
	"encoding/json"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var cfg configs.Config
		err := json.NewDecoder(r.Body).Decode(&cfg)
		if err != nil {
			http.Error(w, "Некорректный JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{"test@example.com"}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))

	}
}

func Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
