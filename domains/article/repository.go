package article

import "sv_backend_test/models"

type Repository interface {
	RCreateArticle(payload models.PayloadPost) error
}
