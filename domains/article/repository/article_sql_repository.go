package repository

import (
	"github.com/jmoiron/sqlx"
	"sv_backend_test/domains/article"
	"sv_backend_test/models"
)

type PsqlArticle struct {
	sqlx *sqlx.DB
}

func NewPsqlArticle(sqlx *sqlx.DB) article.Repository {
	return &PsqlArticle{sqlx}
}

func (pa *PsqlArticle) RCreateArticle(payload models.PayloadPost) error {

	query := "INSERT INTO posts (title,content,category,status) VALUES (?,?,?,?);"
	var err error
	if _, err = pa.sqlx.Exec(query, payload.Title, payload.Content, payload.Category, payload.Status); err != nil {
		return err
	}
	return err

}
