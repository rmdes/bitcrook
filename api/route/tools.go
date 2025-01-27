package route

import (
	"github.com/audioo/bitcrook/api/handlers"
	"github.com/gofiber/fiber/v2"
)

// Tools route
func Tools(app *fiber.App) {
	var h handlers.ToolsHandler
	r := app.Group("/tools")
	r.Get("/", h.Index)
}
