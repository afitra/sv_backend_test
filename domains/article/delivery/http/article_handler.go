package http

import (
	"net/http"
	"sv_backend_test/domains/article"
	"sv_backend_test/models"

	"github.com/labstack/echo"
)

type ArticleHandler struct {
	response       models.Response
	respErrors     models.ResponseErrors
	ArticleUseCase article.Usecase
}

func NewArticleHandler(echoGroup models.EchoGroup, auc article.Usecase) {
	handler := &ArticleHandler{
		ArticleUseCase: auc,
	}
	echoGroup.API.POST("/article", handler.createArticle)
}

func (aha *ArticleHandler) createArticle(c echo.Context) error {

	var request models.PayloadPost
	aha.response, aha.respErrors = models.NewResponse()
	if err := c.Bind(&request); err != nil {
		aha.respErrors.SetTitle(models.MessageUnprocessableEntity)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	if err := c.Validate(request); err != nil {
		aha.respErrors.SetTitle(models.ErrSomethingWrong.Error())
		aha.respErrors.AddError(err.Error())
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	resp, err := aha.ArticleUseCase.UGenerateArticle(c, request)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusCreated, resp)
}
