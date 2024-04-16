package database

import (
	"database/sql"
	"log"
	"school23/internal/models"
	"time"
)

// ---NEWS OPERATIONS BEGIN

func InsertNews(title string, overview string, body string, isNews bool, createdBy string) error {
	query := `
			insert into NEWS(title, overview, body, isNews, createdAt, createdAtMilli, createdBy, accepted)
			values(?, ?, ?, ?, ?, ? ,?, ?);
		`
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	now := time.Now()
	createdAt := now.Format("2006-01-02 15:04:05")
	createdAtMilli := now.Unix()

	_, err = stmt.Exec(title, overview, body, isNews, createdAt, createdAtMilli, createdBy, false)
	return err

}

func UpdateNews(id int, title string, overview string, body string, isNews bool, accepted bool) error {
	query := `
			update NEWS set title = ?, overview = ?, body = ?, isNews = ?, accepted = ?
			where id = ?;
		`
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(title, overview, body, isNews, accepted, id)
	return err
}

func GetAllNews(news bool, page int) []models.News {
	query := "SELECT * FROM NEWS where isNews = ? ORDER BY createdAtMilli desc limit 5 offset ?"
	params := []any{news, page * 5}
	return getAllNews(query, params...)
}

func GetAllNewsByAccepted(news bool, page int, accepted bool) []models.News {
	query := "SELECT * FROM NEWS where isNews = ? and accepted = ? ORDER BY createdAtMilli desc limit 5 offset ?"
	params := []any{news, accepted, page * 5}
	return getAllNews(query, params...)
}

func GetNewsCount(news bool) int {
	query := "SELECT count(*) FROM NEWS where isNews = ?;"
	params := []any{news}
	return getNewsTotalCount(query, params...)
}

func GetNewsCountByAccepted(news bool, accepted bool) int {
	query := "SELECT count(*) FROM NEWS where isNews = ? and accepted = ?;"
	params := []any{news, accepted}
	return getNewsTotalCount(query, params...)
}

func DeleteNewsById(id int) error {
	smt, err := DB.Prepare("delete from NEWS where id = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	defer smt.Close()
	_, err = smt.Exec(id)
	return err
}

func FindNewsById(id int) (models.News, error) {
	smt, err := DB.Prepare("SELECT * FROM NEWS where id = ?")
	if err != nil {
		return models.News{}, err
	}
	row := smt.QueryRow(id)
	defer smt.Close()
	return fromRowToNews(row)
}

// ---TEACHER OPERATIONS END

func getAllNews(query string, params ...any) []models.News {
	newsSlice := []models.News{}
	row, err := DB.Query(query, params...)
	if err != nil {
		return newsSlice
	}
	defer row.Close()

	for row.Next() {
		var id int
		var title string
		var overview string
		var body string
		var isNews bool
		var createdAt string
		var createdAtMilli int64
		var createdBy string
		var accepted bool

		row.Scan(&id,
			&title,
			&overview,
			&body,
			&isNews,
			&createdAt,
			&createdAtMilli,
			&createdBy,
			&accepted,
		)

		newsStruct := models.News{
			Id:             id,
			Title:          title,
			Overview:       overview,
			Body:           body,
			IsNews:         isNews,
			CreatedAt:      createdAt,
			CreatedAtMilli: createdAtMilli,
			CreatedBy:      createdBy,
			Accepted:       accepted,
		}
		newsSlice = append(newsSlice, newsStruct)
	}
	return newsSlice
}

func getNewsTotalCount(query string, params ...any) int {
	var count int = 0
	err := DB.QueryRow(query, params...).Scan(&count)
	if err != nil {
		log.Println(err)
	}
	return count
}

func fromRowToNews(row *sql.Row) (models.News, error) {
	var id int
	var title string
	var overview string
	var body string
	var isNews bool
	var createdAt string
	var createdAtMilli int64
	var createdBy string
	var accepted bool

	err := row.Scan(&id,
		&title,
		&overview,
		&body,
		&isNews,
		&createdAt,
		&createdAtMilli,
		&createdBy,
		&accepted,
	)
	if err != nil {
		return models.News{}, err
	}

	return models.News{
		Id:             id,
		Title:          title,
		Overview:       overview,
		Body:           body,
		IsNews:         isNews,
		CreatedAt:      createdAt,
		CreatedAtMilli: createdAtMilli,
		CreatedBy:      createdBy,
		Accepted:       accepted,
	}, nil
}
