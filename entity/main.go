package entity

type Address struct {
	Country 	string `json:"country"`
	State 		string `json:"state"`
	City 		string `json:"city"`
	Street 		string `json:"street"`
    ZipCode     string `json:"zip_code"`
	Number 		int16  `json:"number"`
}

type User struct {
	Name        string  `json:"username"`
    Password    string  `json:"password"`
    Contact     Contact	`json:"contact"`
}

type Contact struct {
    Email       string
    Cellphone   string
}

type Consumer struct {
	User
	Address		Address `json:"address"`
}

type EstablishmentOwner struct {
	User
}

type Product struct {
	ID					    int32
	Name				    string
	Price 				    float64
	Description		        string
	ImageURL			    string
}

type Establishment struct {
	ID 						int32
	Name					string
	Email					string
	Address				    Address
    Products                []Product
	EstablishmentOwner		EstablishmentOwner
}