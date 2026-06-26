package config

import (
	"SecureAdminPanel/Backend/models"
	"log"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Store = sessions.NewCookieStore([]byte("secret-key"))

var Users = map[string]models.User{}

func InitDB() {
	adminHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	userHash, _ := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)

	Users["admin"] = models.User{
		Username: "admin",
		Password: string(adminHash),
		Role:     "admin",
	}

	Users["user"] = models.User{
		Username: "user",
		Password: string(userHash),
		Role:     "user",
	}

	log.Println("Hash password admin:", string(adminHash))
	log.Println("Hash password user :", string(userHash))
}

func GetUserByUsername(username string) (models.User, bool) {
	user, exists := Users[username]
	return user, exists
}
