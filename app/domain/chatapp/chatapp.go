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
	//Web socket implemented here

	return status{
		Status: "ok",
	}

}
