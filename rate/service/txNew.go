package service

import (
	"errors"
	"forex-be-exercise/conf"
	"log"
	"strings"
)

func NewTx(params map[string]interface{}) (int64, error) {

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
		log.Println("fn : NewTx")
		log.Println("error 1 :", err.Error())
		return 0, err
	}

	if count < 1 {
		log.Println("Loc : rate")
		log.Println("fn : NewTx")
		log.Println("error on checking ")
		return 0, errors.New("list for " + from + "-" + to + " is not exist")
	}

	//add new tx
	var new int64
	err = db.QueryRow(`insert into tx_exchange_rate(currency_from,currency_to,"date",rate)
		values ($1,$2,$3,$4)
		on conflict (currency_from,currency_to,"date") do update 
		set rate = $4
		,modi_date = now()::timestamp
		returning seqno_tx`, from, to, params["date"], params["rate"]).Scan(&new)

	if err != nil {
		log.Println("Loc : rate")
		log.Println("fn : NewTx")
		log.Println("error 2 :", err.Error())
		return 0, err
	}

	return new, nil
}
