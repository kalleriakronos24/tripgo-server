package v1

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	Name string `json:"full_name"`
	Age  int    `json:"age,omitempty"`
}

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

	err := wsjson.Write(ctx, conn, map[string]interface{}{"host": "www.google.de", "port": "80"})
	if err != nil {

	}

	<-ctx.Done()
	conn.Close(websocket.StatusNormalClosure, "")
}

func WSTrackLocation(c *gin.Context) {
	// xx, _ := c.Params.Get("id")

	conn, wsErr := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if wsErr != nil {
		c.AbortWithError(http.StatusInternalServerError, wsErr)
		return
	}
	// conn.CloseNow()

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	defer cancel()

	var v interface{}
	err := wsjson.Read(ctx, conn, &v)
	if err != nil {
		log.Printf("ERROR READ MSG >> %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	// err = wsjson.Write(ctx, conn, v)
	// if err != nil {
	// 	log.Printf("ERROR SEND MSG >> %v", err)
	// 	c.AbortWithError(http.StatusInternalServerError, err)
	// 	return
	// }
	<-ctx.Done()
	conn.Close(websocket.StatusNormalClosure, "")
}

func WSSendLocation(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Minute)
	defer cancel()
	conn, _, err := websocket.Dial(ctx, "ws://192.168.68.116:3009/api/v1/ws/loc/track/qweqwe123123213", nil)
	if err != nil {
		log.Printf("ERR DIAL WRITe >> %v", err)
	}
	defer conn.CloseNow()

	err = wsjson.Write(ctx, conn, map[string]interface{}{"host": "www.google.de", "port": "80"})
	if err != nil {
		log.Printf("ERROR SEND WRITE >> %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	conn.Close(websocket.StatusNormalClosure, "")
	c.JSON(101, nil)
}
