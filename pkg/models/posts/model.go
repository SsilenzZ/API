package posts

type Posts struct {
	ID    int
	User  uint `json:"User"`
	Title string
	Body  string
}
