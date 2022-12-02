package controllers

import (
	"basic-app/database"
	"basic-app/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14) //cost=14, take bytes not string

	//register user
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: pass, //this struct should also take byte therefore change password type in User struct
	}

	database.DB.Create(&user) //insert user into DB
	// Get all records
	result := database.DB.Find(&user)// SELECT * FROM users;

	fmt.Println("HEKKLLLOOADOAOD",result.RowsAffected)// returns found records count, equals `len(users)`
	

	return c.JSON(user)
}
