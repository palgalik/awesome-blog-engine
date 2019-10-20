package comment

import (
	"github.com/galikpali/awesome-blog-engine/internal/entry"
	"github.com/galikpali/awesome-blog-engine/internal/user"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content string
	User    user.User
	Entry   entry.Entry
}
