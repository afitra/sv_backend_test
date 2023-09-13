package article

import "sv_backend_test/models"

type Repository interface {
	RCreateArticle(payload models.PayloadPost) error
	RGetArticle(limit string, offset string) ([]models.Post, error)
	RGetArticleCount() (int, error)
	RGetArticleById(id string) (models.Post, error)
}
