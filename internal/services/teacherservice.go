package services

import (
	"errors"
	"log"
	"mime/multipart"
	"school23/internal/database"
	"school23/internal/models"
	"school23/internal/utils"
)

func GetTeachers() []models.Teacher {
	return database.GetAllTeachers()
}
func GetTeachersPage(page int) models.TeachersPage {
	teacherPage := database.GetTeachersPage(page)
	for i := 0; i < len(teacherPage); i++ {
		teacher := teacherPage[i]
		if !teacher.IsImgPublic {
			teacher.Img = "uploads/img_avatar.png"
			teacherPage[i] = teacher
		}
	}
	totalCount := database.GetTeachersCount()
	pagination := utils.GetPagination(totalCount)
	return models.TeachersPage{Page: page, Teachers: teacherPage, Pagination: pagination}
}

func TeacherCreate(fullname string, birthdate string, phone string, subject string, category string,
	file multipart.File, header *multipart.FileHeader, isImgPublic bool) error {

	uploadedFilePath, err := utils.UploadFile(file, header)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}

	err = database.InsertTeacher(fullname, birthdate, phone, subject, category, uploadedFilePath, isImgPublic)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}

	return nil
}

func TeacherUpdate(id int, fullname string, birthdate string, phone string, subject string, category string,
	file multipart.File, header *multipart.FileHeader, isImgPublic bool) error {

	uploadedFilePath, err := utils.UploadFile(file, header)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}

	err = database.UpdateTeacher(id, fullname, birthdate, phone, subject, category, uploadedFilePath, isImgPublic) //consider error
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func DeleteTeacherById(id int) error {
	err := database.DeleteTeacherById(id)
	if err != nil {
		log.Println(err)
		return errors.New(utils.ServerError)
	}
	return nil
}

func GetTeacher(id int) (models.Teacher, error) {
	user, err := database.FindTeacherById(id)
	if err != nil {
		log.Println(err)
		return models.Teacher{}, errors.New(utils.ServerError)
	}
	return user, nil
}
