package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/mailru/easyjson"
)

var db *sql.DB

func main() {
	var err error
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/contacts", func(c *fiber.Ctx) error {
		rows, err := db.Query("SELECT id, name, phone FROM contacts")
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		defer rows.Close()

		var list Contacts
		for rows.Next() {
			var contact Contact
			if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err == nil {
				list = append(list, contact)
			}
		}

		if list == nil {
			list = Contacts{}
		}

		data, _ := easyjson.Marshal(list)
		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})

	app.Get("/contacts/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var contact Contact

		err := db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id).Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			return c.Status(404).SendString("Contact not found")
		}

		data, _ := easyjson.Marshal(contact)
		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})

	app.Post("/contacts", func(c *fiber.Ctx) error {
		var contact Contact
		if err := easyjson.Unmarshal(c.Body(), &contact); err != nil {
			return c.Status(400).SendString("Invalid JSON")
		}

		err := db.QueryRow("INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id", contact.Name, contact.Phone).Scan(&contact.ID)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		data, _ := easyjson.Marshal(contact)
		c.Set("Content-Type", "application/json")
		return c.Status(201).Send(data)
	})

	app.Put("/contacts/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		var contact Contact

		if err := easyjson.Unmarshal(c.Body(), &contact); err != nil {
			return c.Status(400).SendString("Invalid JSON")
		}

		_, err := db.Exec("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3", contact.Name, contact.Phone, id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		contact.ID = id
		data, _ := easyjson.Marshal(contact)
		c.Set("Content-Type", "application/json")
		return c.Send(data)
	})

	app.Delete("/contacts/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		_, err := db.Exec("DELETE FROM contacts WHERE id = $1", id)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.SendStatus(204)
	})

	app.Listen(":8080")
}
