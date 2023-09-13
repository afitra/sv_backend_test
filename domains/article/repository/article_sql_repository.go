package repository

import (
	"database/sql"
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

func (pa *PsqlArticle) RGetArticle(limit string, offset string) ([]models.Post, error) {

	var result []models.Post
	var err error
	query := `select  title,content,category,status from posts LIMIT ? OFFSET ?`
	if err = pa.sqlx.Select(&result, query, limit, offset); err != nil {
		return result, err
	}
	return result, nil

}

func (pa *PsqlArticle) RGetArticleCount() (int, error) {

	var result []models.Post
	var err error
	query := `select  title from posts`
	if err = pa.sqlx.Select(&result, query); err != nil {
		return 0, err
	}
	return len(result), nil

}

func (pa *PsqlArticle) RGetArticleById(id string) (models.Post, error) {

	var result models.Post
	var err error
	query := `SELECT title,content,category,created_date,updated_date,status FROM posts WHERE id = ?`

	if err := pa.sqlx.Get(&result, query, id); err != nil {
		return result, err
	}
	return result, err

}

func (pa *PsqlArticle) RUpdateArticle(id string, payload models.Post) error {

	query := "UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?"
	var err error
	if _, err = pa.sqlx.Exec(query, payload.Title, payload.Content, payload.Category, payload.Status, id); err != nil {
		return err
	}
	return nil

}

func (pa *PsqlArticle) RDestroyArticle(id string) error {

	query := "DELETE FROM posts WHERE id = ?"
	var result sql.Result
	var err error
	if result, err = pa.sqlx.Exec(query, id); err != nil {
		return err
	}

	// Periksa jumlah baris yang terpengaruh oleh operasi DELETE
	numRows, _ := result.RowsAffected()
	if numRows == 0 {
		return sql.ErrNoRows // ID tidak ada dalam tabel
	}

	return nil

}

func (pa *PsqlArticle) RGetArticleByStatus(status string) ([]models.Post, error) {

	var result []models.Post
	var err error
	query := `select  id,title,content,category,status from posts where status = ?;`
	if err = pa.sqlx.Select(&result, query, status); err != nil {
		return result, err
	}
	return result, nil

}

func (pa *PsqlArticle) RUpdateArticleStatusById(id string, status string) error {
	query := "UPDATE posts SET  status = ? WHERE id = ?"
	var err error
	if _, err = pa.sqlx.Exec(query, status, id); err != nil {
		return err
	}
	return nil
}
