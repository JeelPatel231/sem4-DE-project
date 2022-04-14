package main

import (
	"api/database"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Err error `json:"error"`
}

func main() {
	// localDemo()

	db := database.Connect(os.Getenv("POSTGRES_URL"))

	app := fiber.New()
	app.Get("/getlayout/:uuid", func(c *fiber.Ctx) error {
		data, err := database.GetLayout(db, c.Params("uuid"))
		if err != nil {
			return c.JSON(Error{err})
		}
		return c.JSON(data)
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
