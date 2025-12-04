package chatapp

import (
	"net/http"

	"github.com/CastilloXavier/ChatGo/foundation/logger"
	"github.com/CastilloXavier/ChatGo/foundation/web"
)

// Routes adds specific routes for the chatapp.
func Routes(app *web.App, log *logger.Logger) {
	api := NewApp(log)

	app.HandlerFunc(http.MethodGet, "", "/connect", api.connect)
}
