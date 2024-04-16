package database

import (
	"log"
	"school23/internal/models"
)

func InsertTeacher(fullname string, birthdate string, phone string, subject string, category string, uploadedFilePath string, isImgPublic bool) error {
	query := `
			insert into TEACHERS(fullname, birthdate, subject, category, phone, img, isImgPublic)
			values(?, ?, ?, ?, ? ,?, ?);
		`
	statement, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(fullname, birthdate, subject, category, phone, uploadedFilePath, isImgPublic)
	return err

}

func UpdateTeacher(id int, fullname string, birthdate string, phone string, subject string, category string, uploadedFilePath string, isImgPublic bool) error {
	query := `
			update TEACHERS set fullname = ?, birthdate = ?, subject = ?, category = ?, phone = ?, img = ?, isImgPublic = ?
			where id = ?;
		`
	statement, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(fullname, birthdate, subject, category, phone, uploadedFilePath, isImgPublic, id)
	return err

}

func GetAllTeachers() []models.Teacher {
	row, err := DB.Query("SELECT * FROM TEACHERS ORDER BY id desc")
	if err != nil {
		log.Println(err)
		return []models.Teacher{}
	}
	defer row.Close()

	teachers := make([]models.Teacher, 0)

	for row.Next() {
		var id int
		var fullname string
		var birthdate string
		var subject string
		var category string
		var phone string
		var uploadedFilePath string
		var isImgPublic bool

		err = row.Scan(&id,
			&fullname,
			&birthdate,
			&subject,
			&category,
			&phone,
			&uploadedFilePath,
			&isImgPublic,
		)
		if err == nil {
			teacher := models.Teacher{
				Id:          id,
				Fullname:    fullname,
				Birthdate:   birthdate,
				Subject:     subject,
				Category:    category,
				Phone:       phone,
				Img:         uploadedFilePath,
				IsImgPublic: isImgPublic,
			}
			teachers = append(teachers, teacher)
		}

	}
	return teachers
}

func GetTeachersPage(page int) []models.Teacher {
	row, err := DB.Query("SELECT * FROM TEACHERS ORDER BY id desc limit 5 offset ?", page * 5)
	if err != nil {
		log.Println(err)
		return []models.Teacher{}
	}
	defer row.Close()

	teachers := make([]models.Teacher, 0)

	for row.Next() {
		var id int
		var fullname string
		var birthdate string
		var subject string
		var category string
		var phone string
		var uploadedFilePath string
		var isImgPublic bool

		err = row.Scan(&id,
			&fullname,
			&birthdate,
			&subject,
			&category,
			&phone,
			&uploadedFilePath,
			&isImgPublic,
		)
		if err == nil {
			teacher := models.Teacher{
				Id:          id,
				Fullname:    fullname,
				Birthdate:   birthdate,
				Subject:     subject,
				Category:    category,
				Phone:       phone,
				Img:         uploadedFilePath,
				IsImgPublic: isImgPublic,
			}
			teachers = append(teachers, teacher)
		}
	}
	return teachers
}

func GetTeachersCount() int {
	var count int = 0
	err := DB.QueryRow("SELECT count(*) FROM TEACHERS;").Scan(&count)
	if err != nil {
		log.Println(err)
	}
	return count
}

func DeleteTeacherById(id int) error {
	smt, err := DB.Prepare("delete from TEACHERS where id = ?")
	if err != nil {
		return err
	}
	defer smt.Close()
	_, err = smt.Exec(id)
	return err
}

func FindTeacherById(id int) (models.Teacher, error) {
	smt, err := DB.Prepare("SELECT * FROM TEACHERS where id = ?")
	if err != nil {
		return models.Teacher{}, err
	}
	row := smt.QueryRow(id)
	defer smt.Close()

	var fullname string
	var birthdate string
	var subject string
	var category string
	var phone string
	var uploadedFilePath string
	var isImgPublic bool

	err = row.Scan(&id,
		&fullname,
		&birthdate,
		&subject,
		&category,
		&phone,
		&uploadedFilePath,
		&isImgPublic,
	)
	if err == nil {
		return models.Teacher{
			Id:          id,
			Fullname:    fullname,
			Birthdate:   birthdate,
			Subject:     subject,
			Category:    category,
			Phone:       phone,
			Img:         uploadedFilePath,
			IsImgPublic: isImgPublic,
		}, nil
	}
	return models.Teacher{}, err

}

// ---TEACHER OPERATIONS END
