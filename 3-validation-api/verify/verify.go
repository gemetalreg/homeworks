package verify

import (
	"crypto/sha256"
	"email/api/dto"
	"email/api/storage"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
)

func generateEmailHash(email string) string {
	// 1. Normalize the email (lowercase and trim whitespace)
	cleanEmail := strings.ToLower(strings.TrimSpace(email))

	// 2. Generate the SHA-256 hash
	hash := sha256.Sum256([]byte(cleanEmail))

	// 3. Encode as a hexadecimal string
	return hex.EncodeToString(hash[:])
}

func Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var sendDto dto.SendDto
		err := json.NewDecoder(r.Body).Decode(&sendDto)
		if err != nil {
			http.Error(w, "Некорректный JSON", http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(sendDto)
		if err != nil {
			http.Error(w, "Некорректный email", http.StatusBadRequest)
			return
		}

		hash := generateEmailHash(sendDto.Email)

		storage.Save(dto.VerifyDto{
			Email: sendDto.Email,
			Hash:  hash,
		})

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{"test@example.com"}
		e.Subject = "Awesome Subject"
		e.HTML = []byte(fmt.Sprintf("<h1>Fancy HTML is supported, too!</h1><a href='http://localhost:8081/verify/{%s}'>verification link</a>", hash))
		e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))

	}
}

func Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		verifyDto := storage.Load()
		if hash != verifyDto.Hash {
			fmt.Println(false)
		} else {
			fmt.Println(true)
		}
		storage.Clear()

	}
}
