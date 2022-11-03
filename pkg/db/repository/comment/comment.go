package comment

import (
	"Api/pkg/db/connect"
	"Api/pkg/models/comments"
)

var Repository comments.CommentRepository

func init() {
	Repository = comments.ProvideCommentRepository(connect.DB)
}
