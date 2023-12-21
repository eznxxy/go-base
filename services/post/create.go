package post

import "github.com/eznxxy/go-base/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
