package database

import (
	"log"
	"team-project/services/models"
)

//AddUser adds info about new user to the database
func AddUser(user models.User) int {
	db := OpenDatabase()
	defer db.Close()
	//insert values to the database
	sqlStatement := `
        INSERT INTO users(name,surname,login, password,role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

//GetUserPassword gets user's password and returns password
func GetUserPassword(login string) string {
	var password string
	db := OpenDatabase()
	defer db.Close()
	//get user's password for given login
	sqlStatement := `SELECT password FROM users WHERE login=$1;`
	err := db.QueryRow(sqlStatement, login).Scan(&password)
	//if there's no matches for login return empty value
	if err != nil {
		return ""
	}
	//else return password
	return password
}

//UpdateUserInfo updates user's personal information
func UpdateUser(user models.User, id int) {
	db := OpenDatabase()
	defer db.Close()
	sqlStatement := `UPDATE users
	SET name = $2, surname = $3, login=$4, password=$5, role=$6
	WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id, user.Name, user.Surname, user.Signin.Login, user.Signin.Password, user.Role)
	if err != nil {
		log.Fatal(err)
	}
}

//DeleteUser deletes user's page from db
func DeleteUser(id int) {
	db := OpenDatabase()
	defer db.Close()
	sqlStatement := `DELETE FROM users
	WHERE id = $1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}
}
