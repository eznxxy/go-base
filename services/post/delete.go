package post

import "github.com/eznxxy/go-base/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
