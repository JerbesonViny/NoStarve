package usecases

import (
	"database/sql"

	"github.com/JerbesonViny/nostarve/database"
	"github.com/JerbesonViny/nostarve/entity"
)

func CreateUser(user entity.User) (int64, error) {
	query := "INSERT INTO users (id, name, password) VALUES (nextval('users_sequence'), $1, $2) RETURNING id;";
	lastInsertId := 0;

	connection := database.Connect();

	if err := connection.QueryRow(query, user.Name, user.Password).Scan(&lastInsertId); err != nil && err != sql.ErrNoRows {
		return 0, err;
	}

	return int64(lastInsertId), nil;
}

func VerifyUserExists(email string) (bool, error) {
	query := "select id from contacts where email = $1;";

	connection := database.Connect();

	var identification int; 

	if err := connection.QueryRow(query, email).Scan(&identification); err != nil {

		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
		
	}
	return true, nil
}

