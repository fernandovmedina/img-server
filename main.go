package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// HTML
	var engine = html.New("./public", ".html")

	// Creamos una nueva app de fiber
	var app = fiber.New(fiber.Config{
		Views: engine,
	})

	// Configuracion de archivos estaticos
	app.Static("/", "./public")

	// Ruta de metodo Post para recibir las imagenes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		// Obtenemos el archivo del form con field = "document"
		if file, err := c.FormFile("upload"); err != nil {
			// Si hubo algun error mostramos un mensaje por pantalla
			return c.Status(fiber.StatusInternalServerError).SendString("Error al cargar el archivo")
		} else {
			// Si no hubo ningun error guardamos el archivo en la carpeta "uploads"
			c.SaveFile(file, "uploads/"+file.Filename)
			return c.Render("index", fiber.Map{})
		}
	})

	log.Fatal(app.Listen(":8080"))
}
