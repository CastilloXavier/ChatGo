package chatapp

import (
	"net/http"

	"github.com/CastilloXavier/ChatGo/foundation/web"
)

// Routes adds specific routes for the chatapp.
func Routes(app *web.App) {
	api := NewApp()

	app.HandlerFunc(http.MethodGet, "", "/test", api.test)
}
