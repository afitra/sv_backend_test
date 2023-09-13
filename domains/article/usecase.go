package article

import (
	"github.com/labstack/echo"
	"sv_backend_test/models"
)

type Usecase interface {
	UGenerateArticle(c echo.Context, pl models.PayloadPost) (interface{}, error)
	UGetArticleCreated(c echo.Context) (interface{}, error)
	UGetArticleDataById(c echo.Context) (interface{}, error)
	UpdateArticleData(c echo.Context, pl models.PayloadPost) (interface{}, error)
	UDestroyArticle(c echo.Context) (interface{}, error)
}
