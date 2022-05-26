package helper

import (
	"errors"
	"net/mail"
	"os"

	"github.com/JerbesonViny/nostarve/entity"
)

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func Validateaddress(address entity.Address) bool {
	if len(address.City) > 0  &&  len(address.Country) > 0 && len(address.Street) > 0 && len(address.State) > 0  && len(address.ZipCode) > 0 &&
	address.Number != 0 {
		return true;
	}
	return false
}

func ValidateContact(Contact entity.Contact) (bool, error) {
	if len(Contact.Cellphone) != 11  {
		return false, errors.New("invalid cellphone");
	}

	if _, err := mail.ParseAddress(Contact.Email); err != nil {
		return false, err;
	}
	
	return true, nil;
}

