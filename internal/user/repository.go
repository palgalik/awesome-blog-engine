package user

import "github.com/galikpali/awesome-blog-engine/internal/db"

func findByEmail(email string) User {
	var user User
	db.GetInstance().First(&user, "email = ?", email)
	return user
}
