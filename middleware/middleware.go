package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
	"sv_backend_test/logger"
	"sv_backend_test/models"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type customMiddleware struct {
	e *echo.Echo
}

var echGroup models.EchoGroup

// InitMiddleware to generate all middleware that domains need
func InitMiddleware(ech *echo.Echo, echoGroup models.EchoGroup) {
	cm := &customMiddleware{ech}
	echGroup = echoGroup

	ech.Use(middleware.RequestIDWithConfig(middleware.DefaultRequestIDConfig))
	cm.customLogging()
	cm.customBodyDump()
	ech.Use(middleware.Recover())
	cm.cors()

	cm.customValidation()
}

func (cm *customMiddleware) customBodyDump() {
	cm.e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, req, resp []byte) {
			bodyParser(c, &req)
			reqBody := c.Request()

			logger.MakeWithoutReportCaller(c, req).Info("Request payload for endpoint " + reqBody.Method + " " + reqBody.URL.Path)
			logger.MakeWithoutReportCaller(c, resp).Info("Response payload for endpoint " + reqBody.Method + " " + reqBody.URL.Path)
		},
	}))
}

func (cm *customMiddleware) customLogging() {
	cm.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logrus.SetReportCaller(false)
			req := c.Request()
			res := c.Response()
			reqID := req.Header.Get(echo.HeaderXRequestID)

			if reqID == "" {
				reqID = res.Header().Get(echo.HeaderXRequestID)
			}

			logrus.WithFields(logrus.Fields{
				"requestID":  reqID,
				"method":     req.Method,
				"status":     res.Status,
				"host":       req.Host,
				"user_agent": req.UserAgent(),
				"uri":        req.URL.String(),
				"ip":         c.RealIP(),
			}).Info("Incoming request")
			return next(c)
		}
	})
}

func (cm *customMiddleware) customValidation() {
	validate := validator.New()
	customValidator := customValidator{}
	_ = validate.RegisterValidation("isRequiredWith", customValidator.isRequiredWith)
	_ = validate.RegisterValidation("base64", customValidator.base64)
	_ = validate.RegisterValidation("dateString", customValidator.dateString)
	_ = validate.RegisterValidation("requiredSlice", customValidator.requiredSlice)
	customValidator.validator = validate
	cm.e.Validator = &customValidator
}

func (cm customMiddleware) cors() {
	cm.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"Access-Control-Allow-Origin"},
		AllowMethods: []string{"*"},
	}))
}

// begin custom validator
func (cv *customValidator) isRequiredWith(fl validator.FieldLevel) bool {
	field := fl.Field()
	otherField, _, _, _ := fl.GetStructFieldOK2()

	if otherField.IsValid() && otherField.Interface() != reflect.Zero(otherField.Type()).Interface() {
		if field.IsValid() && field.Interface() == reflect.Zero(field.Type()).Interface() {
			return false
		}
	}

	return true
}
