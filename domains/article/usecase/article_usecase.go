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

func (aus *ArticleUseCase) UpdateArticleData(c echo.Context, pl models.PayloadPost) (interface{}, error) {

	result, err := aus.articleRepo.RGetArticleById(c.Param("id"))
	if err != nil {
		var resp models.ResponseErrorData
		resp.Code = strconv.Itoa(http.StatusBadRequest)
		resp.Title = models.ErrSomethingWrong.Error()
		return resp, err
	}
	//RUpdateArticle
	result.Title = pl.Title
	result.Content = pl.Content
	result.Category = pl.Category
	result.Status = pl.Status

	if err = aus.articleRepo.RUpdateArticle(c.Param("id"), result); err != nil {
		var resp models.ResponseErrorData
		resp.Code = strconv.Itoa(http.StatusBadRequest)
		resp.Title = models.ErrSomethingWrong.Error()
		return resp, err
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusOK)
	resp.Status = models.ResponseSuccess

	return resp, err

}

func (aus *ArticleUseCase) UDestroyArticle(c echo.Context) (interface{}, error) {
	var err error
	if err = aus.articleRepo.RDestroyArticle(c.Param("id")); err != nil {
		var resp models.ResponseErrorData
		resp.Code = strconv.Itoa(http.StatusBadRequest)
		resp.Title = models.ErrSomethingWrong.Error()
		return resp, err
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusOK)
	resp.Status = models.ResponseSuccess

	return resp, err

}

func (aus *ArticleUseCase) UGetArticleByStatus(c echo.Context) (interface{}, error) {

	var post []models.Post
	var err error
	if post, err = aus.articleRepo.RGetArticleByStatus(c.Param("status")); err != nil {
		logger.Make(c, nil).Debug(err)
		return nil, err
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusOK)
	resp.Status = models.ResponseSuccess
	resp.Data = post
	return resp, err

}

func (aus *ArticleUseCase) UChangeStatusArticleById(c echo.Context) (interface{}, error) {
	var err error

	if !isValidStatus(c.Param("status")) {

		var resp models.ResponseErrorData
		resp.Code = strconv.Itoa(http.StatusBadRequest)
		resp.Title = models.ErrSomethingWrong.Error()
		return resp, err

	}

	if err = aus.articleRepo.RUpdateArticleStatusById(c.Param("id"), c.Param("status")); err != nil {
		var resp models.ResponseErrorData
		resp.Code = strconv.Itoa(http.StatusBadRequest)
		resp.Title = models.ErrSomethingWrong.Error()
		return resp, err
	}

	var resp models.Response
	resp.Code = strconv.Itoa(http.StatusOK)
	resp.Status = models.ResponseSuccess
	return resp, err

}

func isValidStatus(input string) bool {

	validStatus := []string{"publish", "draft", "thrash"}
	for _, status := range validStatus {
		if input == status {
			return true
		}
	}
	return false

}
