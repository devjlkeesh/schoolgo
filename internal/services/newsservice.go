package services

import (
	"errors"
	"log"
	"school23/internal/database"
	"school23/internal/models"
	"school23/internal/utils"
)

func GetAllNews(page int, isNews bool) models.NewsPage {
	totalCount := database.GetNewsCount(isNews)
	newsSlice := database.GetAllNews(isNews, page)
	pagination := utils.GetPagination(totalCount)
	return models.NewsPage{
		NewsArray:  newsSlice,
		Pagination: pagination,
		Page:       page,
	}
}

func GetAllAcceptedNews(page int, isNews bool, accepted bool) models.NewsPage {
	totalCount := database.GetNewsCountByAccepted(isNews, accepted)
	newsSlice := database.GetAllNewsByAccepted(isNews, page, accepted)
	pagination := utils.GetPagination(totalCount)
	return models.NewsPage{
		NewsArray:  newsSlice,
		Pagination: pagination,
		Page:       page,
	}
}

func NewsCreate(title string, overview string, body string, isNews bool) error {
	// TODO bu yerda session dagi user gi set bo'lishi kerak
	err := database.InsertNews(title, overview, body, isNews, "Umida Elmurodova")
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func NewsUpdate(id int, title string, overview string, body string, isNews bool, accepted bool) error {
	err := database.UpdateNews(id, title, overview, body, isNews, accepted)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func DeleteNewsById(id int) error {
	err := database.DeleteNewsById(id)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func GetNews(id int) (models.News, error) {
	user, err := database.FindNewsById(id)
	if err != nil {
		log.Println(err)
		return models.News{}, errors.New(utils.ServerError)
	}
	return user, nil
}
