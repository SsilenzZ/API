package post

import (
	"Api/pkg/db/connect"
	"Api/pkg/models/posts"
)

var Repository posts.PostRepository

func init() {
	Repository = posts.ProvidePostRepository(connect.DB)
}
