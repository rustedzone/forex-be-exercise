package service

import (
	"database/sql"
	"forex-be-exercise/conf"
	"log"
	"strconv"
	"time"

	"github.com/akbarkn/aknstruct"
)

func TrackDailyTx(date string) ([]map[string]interface{}, error) {
	db := conf.Db()
	conf.Db().Close()
	defer db.Close()

	type List struct {
		From string `json="from" map="from"`
		To   string `json="to" map="to"`
		Rate string `json="rate" map="rate"`
		Avg  string `json="avg" map="avg"`
	}

	var result []map[string]interface{}
	errResult := make([]map[string]interface{}, 0)

	//get list of currency
	rows, err := db.Query(`select currency_from,currency_to from exchange_rate`)
	if err != nil {
		log.Println("Loc : rate")
		log.Println("fn : ListTransaction")
		log.Println("error 1 :", err.Error())
		return errResult, err
	}
	defer rows.Close()

	for rows.Next() {
		var l List
		err := rows.Scan(&l.From, &l.To)
		if err != nil {
			log.Println("Loc : rate")
			log.Println("fn : ListTransaction")
			log.Println("error 2 :", err.Error())
			return errResult, err
		}

		//get  7-days-before date from requested date
		toDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Println("Loc : rate")
			log.Println("fn : ListTransaction")
			log.Println("error 2 :", err.Error())
			return errResult, err
		}
		fromDate := toDate.AddDate(0, 0, -7)
		fromDateString := fromDate.Format("2006-01-02")

		log.Println("fromDateString :", fromDateString)

		//get currencies daily rate and 7-days-average
		var rate, avg sql.NullFloat64
		err = db.QueryRow(`select ( 
		select coalesce(rate,0)
		from tx_exchange_rate
		where "date"=$2
		and currency_from=$3
		and currency_to=$4
		) as rate, coalesce (avg(rate),0) as average
		from tx_exchange_rate a
		where a."date" >= $1
		and a."date"<=$2
		and currency_from=$3
		and currency_to=$4
		`, fromDateString, date, l.From, l.To).Scan(&rate, &avg)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println("Loc : rate")
				log.Println("fn : ListTransaction")
				log.Println("error 3 :", err.Error())
				return errResult, err
			}
		}

		l.Avg = "insufficient data"
		if avg.Valid {
			l.Avg = strconv.FormatFloat(avg.Float64, 'f', -1, 64)
		}

		l.Rate = "insufficient data "
		if rate.Valid {
			l.Rate = strconv.FormatFloat(rate.Float64, 'f', -1, 64)
		}

		maps := aknstruct.Map(l)
		result = append(result, maps)

	}

	if len(result) < 1 {
		result = errResult
	}

	return result, nil
}
