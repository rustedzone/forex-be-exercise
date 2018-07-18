package controller

import (
	"encoding/json"
	"errors"
	service "forex-be-exercise/rate/service"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func NewTx_(c *gin.Context) {

	//params handler
	var request map[string]interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	//params checker
	if request["from"] == nil || len(request["from"].(string)) <= 0 {
		c.Error(errors.New("field from is missing"))
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": "Field from is Missing !"})
		return
	}
	if request["to"] == nil || len(request["to"].(string)) <= 0 {
		c.Error(errors.New("field to is missing"))
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": "Field to is Missing !"})
		return
	}
	if request["date"] == nil || len(request["date"].(string)) <= 0 {
		c.Error(errors.New("field date is missing"))
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": "Field date is Missing"})
		return
	}
	if request["rate"] == nil {
		c.Error(errors.New("field rate is missing"))
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": "Field rate is Missing"})
		return
	}

	//main function
	data, err := service.NewTx(request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "data": data})

}

func TrackDailyTx_(c *gin.Context) {

	//params checker
	if len(c.Query("date")) <= 0 {
		c.Error(errors.New("field date is missing"))
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": "Field date is Missing"})
		return
	}

	//main function
	data, err := service.TrackDailyTx(c.Query("date"))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "data": data})

}
