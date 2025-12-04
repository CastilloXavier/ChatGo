package chatapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/CastilloXavier/ChatGo/app/sdk/errs"
	"github.com/CastilloXavier/ChatGo/foundation/logger"
	"github.com/CastilloXavier/ChatGo/foundation/web"
	"github.com/gorilla/websocket"
)

type app struct {
	log *logger.Logger
	WS  websocket.Upgrader
}

func NewApp(log *logger.Logger) *app {
	return &app{
		log: log,
	}
}

func (a *app) connect(ctx context.Context, r *http.Request) web.Encoder {
	//Web socket implemented here

	c, err := a.WS.Upgrade(web.GetWriter(ctx), r, nil)
	if err != nil {
		return errs.Newf(errs.FailedPrecondition, "unable to upgrade to websocket: %v", err)
	}
	defer c.Close()

	usr, err := a.handshake(c)
	if err != nil {
		return errs.Newf(errs.FailedPrecondition, "unable to perform handshake: %v", err)
	}

	a.log.Info(ctx, "handshake complete", "usr", usr)

	// var wg sync.WaitGroup
	// wg.Add(3)

	// ticker := time.NewTicker(time.Second)

	// for {
	// 	select {
	// 	case msg, wd := <-ch:

	// 		if !wd {
	// 			return web.NewNoResponse()
	// 		}

	// 		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
	// 			return errs.Newf(errs.FailedPrecondition, "unable to write message: %v", err)
	// 		}

	// 	case <-ticker.C:
	// 		if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
	// 			return web.NewNoResponse()
	// 		}
	// 	}
	// }

	return web.NewNoResponse()
}

func (a *app) handshake(c *websocket.Conn) (user, error) {
	if err := c.WriteMessage(websocket.TextMessage, []byte("HELLO")); err != nil {
		return user{}, fmt.Errorf("write message: w%", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	msg, err := a.readMessage(ctx, c)
	if err != nil {
		return user{}, fmt.Errorf("read message: %v", err)
	}

	var usr user

	if err := json.Unmarshal(msg, &usr); err != nil {
		return user{}, fmt.Errorf("unmarshal message: %v", err)
	}

	v := fmt.Sprintf("WELCOME %s", usr.Name)
	if err := c.WriteMessage(websocket.TextMessage, []byte(v)); err != nil {
		return user{}, fmt.Errorf("write message: w%", err)
	}

	return usr, nil
}

func (a *app) readMessage(ctx context.Context, c *websocket.Conn) ([]byte, error) {
	type response struct {
		msg []byte
		err error
	}
	ch := make(chan response, 1)

	go func() {
		a.log.Info(ctx, "starting to handshake read")
		defer a.log.Info(ctx, "finished to handshake read")
		_, msg, err := c.ReadMessage()
		if err != nil {
			ch <- response{msg: nil, err: err}
		}
		ch <- response{msg: msg, err: nil}
	}()

	var resp response

	select {
	case <-ctx.Done():
		c.Close()
		return nil, ctx.Err()
	case resp = <-ch:
		if resp.err != nil {
			return nil, fmt.Errorf("empty message")
		}
	}
	return resp.msg, nil
}
