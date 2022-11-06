package posts

import "gorm.io/gorm"

type PostRepositoryI interface {
	GetAllPosts() (post []Posts)
	GetPost(id string) (post Posts, err error)
	CreatePost(reqBody map[string]interface{})
	UpdatePost(reqBody map[string]interface{}) (err error)
	DeletePost(id string)
}

type PostRepository struct {
	DB *gorm.DB
}

func ProvidePostRepository(DB *gorm.DB) PostRepository {
	return PostRepository{DB: DB}
}

func (p PostRepository) GetAllPosts() (post []Posts) {
	p.DB.Find(&post)
	return post
}

func (p PostRepository) GetPost(id string) (post Posts, err error) {
	err = p.DB.First(&post, id).Error
	return
}

func (p PostRepository) CreatePost(reqBody map[string]interface{}) {
	p.DB.Model(&Posts{}).Create(&reqBody)
}

func (p PostRepository) UpdatePost(reqBody map[string]interface{}) (err error) {
	var post Posts
	post.ID = reqBody["id"].(int)
	delete(reqBody, "id")
	err = p.DB.Model(&post).Updates(reqBody).Error
	return
}

func (p PostRepository) DeletePost(id string) {
	p.DB.Delete(&Posts{}, id)
}
