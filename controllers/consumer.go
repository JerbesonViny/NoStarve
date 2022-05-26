package controllers

import (
	"fmt"
	"net/http"

	"github.com/JerbesonViny/nostarve/entity"
	"github.com/JerbesonViny/nostarve/helper"
	"github.com/JerbesonViny/nostarve/usecases"
	"github.com/gin-gonic/gin"
)

func CreateConsumer(c *gin.Context) {
	var newUser entity.Consumer;

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		});

		return
	};

	if exist, err := usecases.VerifyUserExists(newUser.Contact.Email); exist {
		c.JSON(http.StatusConflict, gin.H {
			"message": "Email is already exist.",
		});

		return
	} else {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"error": err.Error(),
			});
			
			return
		}
	}

	if _, err := helper.ValidateContact(newUser.Contact); err != nil {
		fmt.Println("Entrei no validate contact")
		return
	}
	//Valida o endereço

	if !helper.Validateaddress(newUser.Address) {
		fmt.Println("Entrei no validate address")
		return
	}

	// Creating a consumer
	userId, err := usecases.CreateUser(newUser.User); 
	
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		
		return 
	}

	// Creating a contact for this user
	if _, err := usecases.CreateContact(userId, newUser.Contact); err != nil{
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	if err := usecases.CreateAddress(userId, newUser.Address); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// Case this user is created successfully
	c.JSON(201, gin.H{
		"user_id": userId,
	})
}

