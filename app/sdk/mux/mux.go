package mux

import (
	"context"
	"net/http"

	"github.com/CastilloXavier/ChatGo/app/domain/chatapp"
	"github.com/CastilloXavier/ChatGo/app/sdk/mid"
	"github.com/CastilloXavier/ChatGo/foundation/logger"
	"github.com/CastilloXavier/ChatGo/foundation/web"
)

type Config struct {
	Log *logger.Logger
}

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(cfg Config) http.Handler {
	logger := func(ctx context.Context, msg string, args ...any) {
		cfg.Log.Info(ctx, msg, args...)
	}

	app := web.NewApp(
		logger,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	chatapp.Routes(app)

	return app
}
