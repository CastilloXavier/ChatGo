package chatapp

import (
	"context"
	"net/http"

	"github.com/CastilloXavier/ChatGo/foundation/web"
)

type app struct {
}

func NewApp() *app {
	return &app{}
}

func (a *app) test(_ context.Context, _ *http.Request) web.Encoder {
	status := status{
		Status: "ok",
	}

	return status
}
