package model

// Reply Model
type Reply struct {
	ID              int64  `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"id"`
	TopicID         int64  `json:"topic_id"`
	Thanks          int64  `json:"thanks"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	MemberID        int64  `json:"member_id"`
	Member          Member `json:"member, omitifempty"`
	Created         int64  `json:"created"`
	LastModified    int64  `json:"last_modified"`
}
