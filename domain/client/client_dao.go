package client

import (
	"errors"
	"log"
	"mitchvillanueva.com/pulseid/datasources/mysql/pulseid_db"
	"mitchvillanueva.com/pulseid/utils"
)

const (
	queryCreateToken = "INSERT INTO client_tokens(token, status, created_at) values(?,?,NOW())"
	queryFindByToken    = `SELECT 
							id, status, created_at 
						FROM 
							client_tokens
						WHERE
							token = ? AND
							status = 'Active' AND
							created_at >= (now() - INTERVAL 8 DAY)`
	queryGetAllToken = "SELECT id, token, status, created_at FROM client_tokens"
)

func (client *Client) CreateToken() error {
	stmt, err := pulseid_db.Client.Prepare(queryCreateToken)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(client.Token, client.Status)
	if saveErr != nil {
		log.Println(saveErr)
		return errors.New("error when tying to find admin")
	}

	client.Id, err = insertResult.LastInsertId()
	if saveErr != nil {
		log.Println(err)
		return errors.New("error when tying to save")
	}

	return nil
}


func (client *Client) FindByToken() error {
	stmt, err := pulseid_db.Client.Prepare(queryFindByToken)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	token := utils.MD5hash(client.Token)

	result := stmt.QueryRow(token)
	if getErr := result.Scan(&client.Id, &client.Status, &client.Created_at); getErr != nil {
		log.Println(getErr)
		return errors.New("error when tying to find token")
	}

	return nil
}

func (client *Client) GetAllToken() ([]Client, error) {

	clients := []Client{}

	stmt, err := pulseid_db.Client.Prepare(queryGetAllToken)
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	results, err := stmt.Query()

	if err != nil {
		log.Println(err)
	}

	for results.Next(){
		client := Client{}
		results.Scan(&client.Id, &client.Token, &client.Status, &client.Created_at)
		clients = append(clients, client)
	}

	return clients, nil
}
