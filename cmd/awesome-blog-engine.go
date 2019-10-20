package main

import (
	"github.com/galikpali/awesome-blog-engine/internal/comment"
	"github.com/galikpali/awesome-blog-engine/internal/db"
	"github.com/galikpali/awesome-blog-engine/internal/entry"
	"github.com/galikpali/awesome-blog-engine/internal/user"
)

func init() {
	db.GetInstance().AutoMigrate(&user.User{}, &entry.Entry{}, &comment.Comment{})
}

func main() {

}
