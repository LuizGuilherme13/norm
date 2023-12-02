package test

import (
	"testing"
)

func TestCreate(t *testing.T) {
	// unidb := unodatabase.New(models.DBConn{
	// 	Driver:   "postgres",
	// 	Host:     "localhost",
	// 	Port:     "5432",
	// 	User:     "postgres",
	// 	Password: "postgres",
	// 	DBName:   "unidb",
	// 	SSLMode:  "disable",
	// })

	// account := tables.Account{
	// 	UserName:  faker.Name(),
	// 	Password:  faker.Password(),
	// 	Email:     faker.Email(),
	// 	CreatedOn: time.Now(),
	// }

	// unidb.Migrate(&account)

	// err := unidb.Model(&account).Create()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// fmt.Println("Result =>", account.UserName)
}
