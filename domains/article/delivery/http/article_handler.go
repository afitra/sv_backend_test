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
	echoGroup.API.GET("/article/:limit/:offset", handler.getArticle)
	echoGroup.API.GET("/article/:id", handler.getArticleDataById)
	echoGroup.API.PUT("/article/:id", handler.updateArticleData)
	echoGroup.API.DELETE("/article/:id", handler.destroyArticleData)
	echoGroup.API.GET("/article/status/:status", handler.getArticleByStatus)
	echoGroup.API.PUT("/article/status/:id/:status", handler.changeStatusArticleById)

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

func (aha *ArticleHandler) getArticle(c echo.Context) error {

	resp, err := aha.ArticleUseCase.UGetArticleCreated(c)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusOK, resp)
}

func (aha *ArticleHandler) getArticleDataById(c echo.Context) error {

	resp, err := aha.ArticleUseCase.UGetArticleDataById(c)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusOK, resp)
}

func (aha *ArticleHandler) updateArticleData(c echo.Context) error {

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

	resp, err := aha.ArticleUseCase.UpdateArticleData(c, request)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusOK, resp)
}

func (aha *ArticleHandler) destroyArticleData(c echo.Context) error {

	resp, err := aha.ArticleUseCase.UDestroyArticle(c)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusOK, resp)
}

func (aha *ArticleHandler) getArticleByStatus(c echo.Context) error {

	resp, err := aha.ArticleUseCase.UGetArticleByStatus(c)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusOK, resp)
}

func (aha *ArticleHandler) changeStatusArticleById(c echo.Context) error {

	resp, err := aha.ArticleUseCase.UChangeStatusArticleById(c)
	if err != nil {
		errMap := resp.(models.ResponseErrorData)
		aha.respErrors.SetTitleCode(errMap.Code, errMap.Title, errMap.Description)
		aha.response.SetResponse("", &aha.respErrors)
		return aha.response.Body(c, err)
	}

	aha.response.SetResponse(&resp, &aha.respErrors)
	return c.JSON(http.StatusOK, resp)
}
