package v1

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
)

func GETWSConn(c *gin.Context) {
	xx, _ := c.Params.Get("id")

	log.Printf("%v", xx)
	conn, wsErr := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if wsErr != nil {
		c.AbortWithError(http.StatusInternalServerError, wsErr)
		return
	}

	defer conn.Close(websocket.StatusInternalError, "Closed unexepetedly")

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)

	defer cancel()
	ctx = conn.CloseRead(ctx)

	conn.Write(ctx, websocket.MessageText, []byte("Hello, WebSocket!"))

	<-ctx.Done()
	conn.Close(websocket.StatusNormalClosure, "")
}
