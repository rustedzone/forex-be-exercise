package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	service "forex-be-exercise/rate/service"

	"gopkg.in/gin-gonic/gin.v1"
)

func NewList_(c *gin.Context) {

	//params handler
	var request map[string]interface{}
	err := json.NewDecoder(c.Request.Body).Decode(&request)
	if err != nil {
		log.Println(c.Request)
		log.Println("error on parsing request")
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

	//main function
	data, err := service.NewList(request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "data": data})
}

func RemoveFromList_(c *gin.Context) { //params handler
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

	//main function
	data, err := service.RemoveFromList(request)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"result": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": true, "data": data})

}
