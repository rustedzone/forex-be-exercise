package test

import (
	"errors"
	"forex-be-exercise/route"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTrackDailyExchangeRate(t *testing.T) {
	testRouter := route.Router()

	req, err := http.NewRequest("GET", "/api/daily?date=2018-07-01", nil)
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
