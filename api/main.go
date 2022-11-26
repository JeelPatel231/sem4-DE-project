package main

import (
	"api/database"
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// localDemo()

	db := database.Connect(os.Getenv("POSTGRES_URL"))

	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "*",
		},
	))

	app.Get("/getlayout/:uuid", func(c *fiber.Ctx) error {
		data, err := database.GetLayout(db, c.Params("uuid"))
		if err != nil {
			return err
		}
		return c.JSON(data)
	})

	app.Post("/setlayout/:uuid", func(c *fiber.Ctx) error {
		var postBody [][]int
		json.Unmarshal(c.Body(), &postBody)
		err := database.UpdateLayout(db, c.Params("uuid"), postBody)
		if err != nil {
			return err
		}

		return c.SendStatus(200)
	})

	app.Post("/register/:name", func(c *fiber.Ctx) error {
		resp, err := database.InsertRestaurant(db, c.Params("name"))
		if err != nil {
			return err
		}
		err = database.InsertLayout(db, *resp, [][]int{})
		if err != nil {
			return err
		}
		return c.Send([]byte(*resp))
	})

	app.Post("/deleteres/:uuid", func(c *fiber.Ctx) error {
		err := database.DeleteRestaurant(db, c.Params("uuid"))
		if err != nil {
			return err
		}
		return c.SendStatus(200)
	})

	app.Post("/delete/:uuid", func(c *fiber.Ctx) error {
		err := database.DeleteLayout(db, c.Params("uuid"))
		if err != nil {
			return err
		}
		return c.SendStatus(200)
	})

	app.Listen(":3000")

}

func localDemo() {
	// empty = 0
	// seat = 1
	// seat occupied = 2
	// table = 3

	asciiBlocks := []string{"  ", "⬤ ", "⚫", "██"}

	str := [][]int{
		{2, 3, 3, 1, 0, 2, 3, 3, 2},
		{1, 3, 3, 2, 0, 2, 3, 3, 2},
		{2, 3, 3, 1, 0, 1, 3, 3, 1},
	}

	str2 := [][]int{
		{0, 1, 1, 0, 0, 1, 3, 1, 0, 1, 3, 1},
		{1, 3, 3, 1, 0, 1, 3, 1, 1, 1, 3, 1},
		{1, 3, 3, 1, 0, 1, 3, 3, 3, 3, 3, 1},
		{0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0},
	}

	for _, matrices := range [][][]int{str, str2} {
		for _, k := range matrices {
			for _, j := range k {
				print(asciiBlocks[j])
			}
			print("\n")
		}
		print("\n\n")
	}
}
