package article

import (
	"github.com/labstack/echo"
	"sv_backend_test/models"
)

type Usecase interface {
	UGenerateArticle(c echo.Context, pl models.PayloadPost) (interface{}, error)
}
