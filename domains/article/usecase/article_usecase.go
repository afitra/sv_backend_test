package usecase

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"sv_backend_test/domains/article"
	"sv_backend_test/logger"
	"sv_backend_test/models"
)

type ArticleUseCase struct {
	articleRepo article.Repository
}

func NewArticleUseCase(arp article.Repository) article.Usecase {
	return &ArticleUseCase{
		articleRepo: arp,
	}
}

func (aus *ArticleUseCase) UGenerateArticle(c echo.Context, pl models.PayloadPost) (interface{}, error) {

	err := aus.articleRepo.RCreateArticle(pl)

	if err != nil {
		logger.Make(c, nil).Debug(err)
		return nil, err
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusCreated)
	resp.Status = models.ResponseSuccess
	return resp, err

}
