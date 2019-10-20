package entry

import (
	"github.com/galikpali/awesome-blog-engine/internal/user"
	"github.com/jinzhu/gorm"
)

type Entry struct {
	gorm.Model
	Content string
	User    user.User
}
