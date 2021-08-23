package admin

import (
	"errors"
	"log"
	"mitchvillanueva.com/pulseid/datasources/mysql/pulseid_db"
	"mitchvillanueva.com/pulseid/utils"
)

const (
	queryFindByUsernameAndPassword = "SELECT id, username FROM admin WHERE username=? AND password=?"
)

func (admin *Admin) FindByUsernameAndPassword() error {
	stmt, err := pulseid_db.Client.Prepare(queryFindByUsernameAndPassword	)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	password := utils.MD5hash(admin.Password)

	result := stmt.QueryRow(admin.Username, password)
	if getErr := result.Scan(&admin.Id, &admin.Username); getErr != nil {
		log.Println(getErr)
		return errors.New("error when tying to find admin")
	}

	return nil
}

