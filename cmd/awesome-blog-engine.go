package main

import (
	"github.com/galikpali/awesome-blog-engine/internal/auth"
	"github.com/galikpali/awesome-blog-engine/internal/comment"
	"github.com/galikpali/awesome-blog-engine/internal/db"
	"github.com/galikpali/awesome-blog-engine/internal/entry"
	"github.com/galikpali/awesome-blog-engine/internal/user"
	"log"
	"net/http"
)

func init() {
	db.GetInstance().AutoMigrate(&user.User{}, &entry.Entry{}, &comment.Comment{})
}

func main() {
	http.HandleFunc("/auth", auth.Auth)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
