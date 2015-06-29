package model

// Node Model
type Node struct {
	ID               int64  `gorm:"primary_key" json:"id" sql:"AUTO_INCREMENT"`
	Name             string `gorm:"primary_key" json:"name"`
	Title            string `json:"title"`
	TitleAlternative string `json:"title_alternative"`
	URL              string `json:"url"`
	Topics           int64  `json:"topics"`
	AvatarMini       string `json:"avatar_mini"`
	AvatarNormal     string `json:"avatar_normal"`
	AvatarLarge      string `json:"avatar_large"`
	Created          int64  `json:"created"`
	Stars            int64  `json:"stars"`
	Header           string `json:"header"`
	Footer           string `json:"footer"`
}
