package service

import (
	"errors"
	"forex-be-exercise/conf"
	"log"
	"strings"
)

func RemoveFromList(params map[string]interface{}) (string, error) {
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
		log.Println("fn : RemoveFromList")
		log.Println("error 1 :", err.Error())
		return "", err
	}

	if count < 1 {
		log.Println("Loc : rate")
		log.Println("fn : RemoveFromList")
		log.Println("error on checking ")
		return "", errors.New("list for " + from + "-" + to + " is not exist")
	}

	//remove the record
	stmt, err := db.Prepare(`delete from exchange_rate where currency_from=$1 and currency_to=$2`)
	if err != nil {
		log.Println("Loc : rate")
		log.Println("fn : RemoveFromList")
		log.Println("error 2 :", err.Error())
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(from, to)
	if err != nil {
		log.Println("Loc : rate")
		log.Println("fn : RemoveFromList")
		log.Println("error 3 :", err.Error())
		return "", err
	}

	return "", nil
}
