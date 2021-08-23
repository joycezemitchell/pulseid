package pulseid_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	host     = os.Getenv(mysqlUsersHost)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	log.Println("THIS IS MY VARIABLE")
	log.Println(os.Getenv(mysqlUsersUsername))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err)
	}
	if err = Client.Ping(); err != nil {
		log.Println(err)
	}

	log.Println("database successfully configured")
}