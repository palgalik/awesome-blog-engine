package user

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/galikpali/awesome-blog-engine/internal/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtKey = []byte(config.GetConfig().JWT.Secret)

func Auth(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userEntity := findByEmail(credentials.Email)

	if !checkPasswordHash(credentials.Password, userEntity.PasswordHash) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	addTokenToResponseHeader(credentials, err, w)
}

func addTokenToResponseHeader(credentials Credentials, err error, w http.ResponseWriter) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
