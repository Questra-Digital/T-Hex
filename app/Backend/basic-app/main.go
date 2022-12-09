package main

import (
	"basic-app/database"
	"basic-app/routes"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()
	routes.Setup(app)

	fmt.Println("Go Redis ")
	// // another Popular method for installation
	// opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	// if err != nil {
	// 	panic(err)
	// }
	// rdb := redis.NewClient(opt)

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping(client.Context()).Result() //client.Context()
	fmt.Println(pong, err)

	app.Listen(":8000")
}
