package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"size:250" json:"username"`
	Href           string `gorm:"size:300" json:"href"`
	NumberPosts    int    `gorm:"size:10" json:"number_posts"`
	NumFollowers   int    `gorm:"size:10" json:"number_followers"`
	NumFollowing   int    `gorm:"size:10" json:"number_following"`
	Private        bool
	BlackList      bool // Perfis que eu já segui e deixei se seguir
	DeadFile       bool // Perfis que não irei mais seguir (que já me seguem ou que já segui muitas vezes)
	FollowMe       bool // O perfil me segue?
	IFollow        bool
	DateIFollowed  time.Time
	TimesIFollowed int // Conta quantas vezes eu já segui esse perfil. Usar isso como limite para seguir ou não
}
