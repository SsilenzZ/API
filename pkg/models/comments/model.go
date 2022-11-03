package comments

type Comments struct {
	PostID uint `json:"PostID"`
	ID     int
	Name   string
	Email  string
	Body   string
}
