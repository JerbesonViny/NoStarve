package usecases

import (
	"github.com/JerbesonViny/nostarve/database"
	"github.com/JerbesonViny/nostarve/entity"
)

func CreateAddress(userId int64, Address entity.Address) error {
	query := "INSERT INTO address (id, country, state, city, street, zip_code, number, user_id) VALUES (nextval('address_sequence'),$1, $2, $3, $4, $5, $6, $7);"
	
	connection := database.Connect()
	if err := connection.QueryRow(query, Address.Country, Address.State, Address.City, Address.Street, Address.ZipCode, Address.Number, userId); err != nil {
		return err.Err()
	}
	return nil
}	
