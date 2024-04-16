package database

import (
	"log"
	"school23/internal/models"
)

func InsertUser(user models.User) error {
	query := `
		insert into USERS(username, password, role, status, lastLoginAt)
		values(?,?,?,?,'');
	`
	statement, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Username, user.Password, user.Role, false)
	return err

}

func UpdateUser(id int, username string, role string, status bool) error {
	smt, err := DB.Prepare(" update USERS set username = ?, role = ?, status = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = smt.Exec(username, role, status, id)
	return err
}

func FindUserById(id int) (models.User, error) {
	smt, err := DB.Prepare("SELECT * FROM USERS where id = ?")
	if err != nil {
		return models.User{}, err
	}
	row := smt.QueryRow(id)
	defer smt.Close()

	var username string
	var password string
	var role string
	var status bool
	var lastLoginAt string
	err = row.Scan(&id, &username, &password, &role, &status, &lastLoginAt)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Id:          id,
		Username:    username,
		Password:    password,
		Role:        role,
		Status:      status,
		LastLoginAt: lastLoginAt,
	}
	return user, nil
}

func GetAllUsers() []models.User {
	row, err := DB.Query("SELECT * FROM USERS ORDER BY id desc")
	if err != nil {
		log.Println(err)
		return []models.User{}
	}
	defer row.Close()
	users := make([]models.User, 0)

	for row.Next() {

		var id int
		var username string
		var password string
		var role string
		var status bool
		var lastLoginAt string

		err = row.Scan(&id, &username, &password, &role, &status, &lastLoginAt)

		if err == nil {
			users = append(users, models.User{
				Id:          id,
				Username:    username,
				Password:    password,
				Role:        role,
				Status:      status,
				LastLoginAt: lastLoginAt,
			})
		}
	}
	return users
}

func DeleteUserById(id int) error {
	smt, err := DB.Prepare("delete from USERS where id = ?")
	if err != nil {
		return err
	}
	defer smt.Close()
	_, err = smt.Exec(id)
	return err
}

func UserExistsByUsername(username string) (bool, error) {
	var count int = 0
	err := DB.QueryRow("select count(*) from USERS where LOWER(username) = LOWER(?);").Scan(&count)
	return count != 0, err
}

func FindUserByUsername(uname string) (models.User, error) {
	smt, err := DB.Prepare("SELECT * FROM USERS where LOWER(username) = LOWER(?);")
	if err != nil {
		return models.User{}, err
	}
	row := smt.QueryRow(uname)
	defer smt.Close()

	var id int
	var username string
	var password string
	var role string
	var status bool
	var lastLoginAt string
	err = row.Scan(&id, &username, &password, &role, &status, &lastLoginAt)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Id:          id,
		Username:    username,
		Password:    password,
		Role:        role,
		Status:      status,
		LastLoginAt: lastLoginAt,
	}, nil
}
