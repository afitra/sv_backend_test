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

func (aus *ArticleUseCase) UGetArticleCreated(c echo.Context) (interface{}, error) {
	var count int
	var post []models.Post
	var err error
	if post, err = aus.articleRepo.RGetArticle(c.Param("limit"), c.Param("offset")); err != nil {
		logger.Make(c, nil).Debug(err)
		return nil, err
	}

	if count, err = aus.articleRepo.RGetArticleCount(); err != nil {
		logger.Make(c, nil).Debug(err)
		return nil, err
	}

	result := models.ArticleCreated{
		Count: count,
		Post:  post,
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusOK)
	resp.Status = models.ResponseSuccess
	resp.Data = result
	return resp, err

}

func (aus *ArticleUseCase) UGetArticleDataById(c echo.Context) (interface{}, error) {

	result, err := aus.articleRepo.RGetArticleById(c.Param("id"))
	if err != nil {
		var resp models.ResponseErrorData
		resp.Code = strconv.Itoa(http.StatusBadRequest)
		resp.Title = models.ErrSomethingWrong.Error()
		return resp, err
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusOK)
	resp.Status = models.ResponseSuccess
	resp.Data = result
	return resp, err

}
