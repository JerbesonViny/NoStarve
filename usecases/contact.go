package usecases

import (
	"github.com/JerbesonViny/nostarve/database"
	"github.com/JerbesonViny/nostarve/entity"
)

func CreateContact(userId int64, Contact entity.Contact) (int64, error){
	query := "INSERT INTO contacts (id, email, cellphone, user_id) VALUES (nextval('contacts_sequence'), $1, $2, $3) RETURNING id;"

	var contactId int64
	connection := database.Connect();
	if err:= connection.QueryRow(query, Contact.Email, Contact.Cellphone, userId).Scan(&contactId); err != nil {
		return 0, err
	}

	return contactId, nil 

}
