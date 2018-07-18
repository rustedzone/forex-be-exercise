package service

import (
	"errors"
	"forex-be-exercise/conf"
	"log"
	"strings"
)

func NewList(params map[string]interface{}) (string, error) {
	db := conf.Db()
	conf.Db().Close()
	defer db.Close()

	from := strings.ToUpper(params["from"].(string))
	to := strings.ToUpper(params["to"].(string))

	//check
	var count int
	err := db.QueryRow(`select count(0) from exchange_rate where currency_from=$1 and currency_to=$2`, from, to).Scan(&count)
	if err != nil {
		log.Println("Loc : rate")
		log.Println("fn : NewList")
		log.Println("error 1 :", err.Error())
		return "", err
	}

	if count > 0 {
		log.Println("Loc : rate")
		log.Println("fn : NewList")
		log.Println("error on checking ")
		return "", errors.New("list for " + from + "-" + to + " already exist")
	}

	//add new record
	var new string
	err = db.QueryRow(`insert into exchange_rate(currency_from,currency_to) values($1,$2) returning currency_from ||'-'|| currency_to`, from, to).Scan(&new)
	if err != nil {
		log.Println("Loc : rate")
		log.Println("fn : NewList")
		log.Println("error 2 :", err.Error())
		return "", err
	}

	return new, nil
}
