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

func TestRemoveFromList(t *testing.T) {
	testRouter := route.Router()

	form := make(map[string]interface{})

	form["from"] = "FROM"
	form["to"] = "to"
	byt, err := json.Marshal(form)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	req, err := http.NewRequest("DELETE", "/api/list", bytes.NewBuffer(byt))
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
