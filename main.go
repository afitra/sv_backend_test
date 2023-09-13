package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	migrate "github.com/rubenv/sql-migrate"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"os"
	_articleHandler "sv_backend_test/domains/article/delivery/http"
	_articleRepository "sv_backend_test/domains/article/repository"
	_articleUseCase "sv_backend_test/domains/article/usecase"
	"sv_backend_test/logger"
	"sv_backend_test/models"
	"time"
)

var ech *echo.Echo

type CustomValidator struct {
	validator *validator.Validate
}

func main() {

	sqlConn, sqlxConn := getDBConn()
	_ = dataMigrations(sqlxConn)
	defer sqlConn.Close()

	echoGroup := models.EchoGroup{
		API: ech.Group(""),
	}

	customValidator := validator.New()

	// Register validator dengan Echo
	ech.Validator = &CustomValidator{validator: customValidator}
	articleRepository := _articleRepository.NewPsqlArticle(sqlxConn)
	articleUsecase := _articleUseCase.NewArticleUseCase(articleRepository)
	_articleHandler.NewArticleHandler(echoGroup, articleUsecase)

	ech.GET("/ping", ping)

	err := ech.Start(":" + os.Getenv(`PORT`))

	if err != nil {
		logger.Make(nil, nil).Debug(err)
	}
}

func init() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	ech = echo.New()
	ech.Debug = true
	loadEnv()
	logger.Init()
}
func loadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return
	}

	err := godotenv.Load()

	if err != nil {
		logger.Make(nil, nil).Fatal("Error loading .env file")
	}
}

func ping(echTx echo.Context) error {
	response := models.Response{}
	response.Status = models.StatusSuccess
	response.Message = "Server Actived!!"

	return echTx.JSON(http.StatusOK, response)
}

func getDBConn() (*sql.DB, *sqlx.DB) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbDialect := os.Getenv("DB_DIALECT")

	dbConfig := mysql.Config{
		User:                 dbUser,
		Passwd:               dbPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName:               dbName,
		ParseTime:            true, // Ini adalah opsi yang memungkinkan parsing tanggal dan waktu
		AllowNativePasswords: true,
	}

	// Format DSN (Data Source Name) dari konfigurasi
	dsn := dbConfig.FormatDSN()

	sqlxConn, err := sqlx.Connect(dbDialect, dsn)

	if err != nil {
		return nil, nil
	}
	// Connection for sql/migrations
	sqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	sqlConn, err := sql.Open("mysql", sqlConnection)
	if err != nil {
		return nil, nil
	}

	err = sqlxConn.Ping()
	if err != nil {
		return nil, nil
	}

	err = sqlConn.Ping()
	if err != nil {
		return nil, nil
	}

	return sqlConn, sqlxConn
}

func dataMigrations(dbConn *sqlx.DB) error {
	migrations := &migrate.FileMigrationSource{
		Dir: "db/migrations/", // Ganti dengan direktori tempat file migrasi Anda
	}
	n, err := migrate.Exec(dbConn.DB, "mysql", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Jumlah migrasi yang diterapkan: %d\n", n)

	return nil
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}
