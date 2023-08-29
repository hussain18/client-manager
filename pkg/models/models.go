package models

type ChangeHistory struct {
	Field     string `bson:"field"`
	From      string `bson:"from"`
	To        string `bson:"to"`
	Timestamp string `bson:"timstamp"`
}

type Client struct {
	FirstName       string          `bson:"first_name"`
	LastName        string          `bson:"last_name"`
	TelephoneNumber string          `bson:"telephone_number"`
	EmailAddress    string          `bson:"email_address"`
	Street          string          `bson:"street"`
	City            string          `bson:"city"`
	Country         string          `bson:"country"`
	PostalCode      string          `bson:"postal_code"`
	changeHistory   []ChangeHistory `bson:"change_history"`
}

type Admin struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}
