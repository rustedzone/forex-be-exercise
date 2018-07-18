package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"forex-be-exercise/route"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddDailyExchangeRate(t *testing.T) {
	testRouter := route.Router()

	form := make(map[string]interface{})

	form["from"] = "FROM"
	form["to"] = "TO"
	form["rate"] = 0.75799
	form["date"] = "2018-07-01"
	byt, err := json.Marshal(form)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest("POST", "/api/daily", bytes.NewBuffer(byt))
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	if resp.Code != 200 {
		err := errors.New("SOMETHING WENT WRONG")
		t.Errorf("Error %s", err.Error())
		t.Fail()
		return
	}

}
