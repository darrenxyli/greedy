package model

// Topic Model
type Topic struct {
	ID              int64  `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"id"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	Content         string `json:"content"`
	ContentRendered string `json:"content_rendered"`
	Replies         int64  `json:"replies"`
	MemberID        int64  `json:"member_id"`
	Member          Member `json:"member, omitifempty"`
	NodeID          int64  `json:"node_id"`
	Node            Node   `json:"node, omitifempty"`
	Created         int64  `json:"created"`
	LastedModified  int64  `json:"lasted_modified"`

	ReplyItems []Reply `json:"-"`
}
