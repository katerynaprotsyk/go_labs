package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mailru/easyjson"
)

var notesDB = make(map[int]Note)
var nextID = 1

func main() {
	app := fiber.New()

	app.Get("/notes", func(c *fiber.Ctx) error {
		var list Notes
		for _, note := range notesDB {
			list = append(list, note)
		}

		data, _ := easyjson.Marshal(list)
		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})

	app.Get("/notes/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			errResp := ErrorResponse{Error: "Invalid ID format"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		note, exists := notesDB[id]
		if !exists {
			errResp := ErrorResponse{Error: "Note not found"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(404).Send(data)
		}

		data, _ := easyjson.Marshal(note)
		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})

	app.Post("/notes", func(c *fiber.Ctx) error {
		var note Note
		if err := easyjson.Unmarshal(c.Body(), &note); err != nil {
			errResp := ErrorResponse{Error: "Invalid JSON body"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		if note.Title == "" {
			errResp := ErrorResponse{Error: "Title is required"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		note.ID = nextID
		notesDB[nextID] = note
		nextID++

		data, _ := easyjson.Marshal(note)
		c.Set("Content-Type", "application/json")
		return c.Status(201).Send(data)
	})

	app.Put("/notes/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			errResp := ErrorResponse{Error: "Invalid ID format"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		if _, exists := notesDB[id]; !exists {
			errResp := ErrorResponse{Error: "Note not found"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(404).Send(data)
		}

		var updatedNote Note
		if err := easyjson.Unmarshal(c.Body(), &updatedNote); err != nil {
			errResp := ErrorResponse{Error: "Invalid JSON body"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		if updatedNote.Title == "" {
			errResp := ErrorResponse{Error: "Title is required"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		updatedNote.ID = id
		notesDB[id] = updatedNote

		data, _ := easyjson.Marshal(updatedNote)
		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})

	app.Delete("/notes/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			errResp := ErrorResponse{Error: "Invalid ID format"}
			data, _ := easyjson.Marshal(errResp)
			c.Set("Content-Type", "application/json")
			return c.Status(400).Send(data)
		}

		delete(notesDB, id)
		return c.SendStatus(204)
	})

	app.Listen(":8080")
}
