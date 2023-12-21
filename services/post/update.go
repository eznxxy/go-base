package post

import (
	"github.com/eznxxy/go-base/models"
	"github.com/eznxxy/go-base/requests"
)

func (postService *Service) Update(post *models.Post, updatePostRequest *requests.UpdatePostRequest) {
	post.Content = updatePostRequest.Content
	post.Title = updatePostRequest.Title
	postService.DB.Save(post)
}
