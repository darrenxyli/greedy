package model

// Member Model
type Member struct {
	ID            int64  `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"id"`
	URL           string `json:"url"`
	Username      string `json:"username"`
	Password      string `json:"-"`
	Email         string `json:"email"`
	Website       string `json:"website"`
	Twitter       string `json:"twitter"`
	Facebook      string `json:"facebook"`
	FacebookToken string `json:"-"`
	Location      string `json:"location"`
	Tagline       string `json:"tagline"`
	Bio           string `json:"bio"`
	AvatarMini    string `json:"avatar_mini"`
	AvatarNormal  string `json:"avatar_normal"`
	AvatarLarge   string `json:"avatar_large"`
	Created       int64  `json:"created"`

	TopicItems []Topic `json:"-"`
	ReplyItems []Reply `json:"-"`
}
