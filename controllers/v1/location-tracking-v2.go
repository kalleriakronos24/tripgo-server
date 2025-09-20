package v1

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kalleriakronos24/khaimal-group/pkg/location"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	ws "github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		// log.Printf("%v >> ", origin)
		return true
	},
} // use default option

var clients []*websocket.Conn

var (
	pingPeriod = (pongPeriod * 9) / 10
	pongPeriod = 60 * time.Second
	writeWait  = 10 * time.Second
)

func PingResponse(ws *websocket.Conn, done chan struct{}) {
	defer close(done)
	// conf := storage.GetConfig()
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongPeriod))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongPeriod)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				// conf.Log.Debugf("Websocket Ping Read Failed: %v", err)
			}
			return
		} else {
			// conf.Log.Debugf("Received message from Websocket client")
		}
	}
}

// done := make(chan struct{})
func AllUserWebsocketWriter(ws *websocket.Conn, done chan struct{}) {
	// conf := storage.GetConfig()
	pingticker := time.NewTicker(pingPeriod)
	defer func() {
		pingticker.Stop()
		ws.Close()
	}()

	ws.SetWriteDeadline(time.Now().Add(writeWait))
	err := ws.WriteJSON("hellow")
	if err != nil {
		log.Printf("Failed to write initial user response: %v", err)
		return
	}
	clients = append(clients, ws)

	for {
		select {
		case <-done:
			for {
				// Read message from browser
				msgType, msg, err := ws.ReadMessage()
				if err != nil {
					log.Printf("failed read message >> %v", err)
				}
				// Print the message to the console
				fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

				for _, client := range clients {
					// log.Printf("clients >> %v", client)
					// Write message back to browser
					if err = client.WriteMessage(msgType, msg); err != nil {
						log.Printf("failed send message >> %v", err)
						// ws.Close()
					}
				}
			}
			// ws.SetWriteDeadline(time.Now().Add(writeWait))
			// err := ws.WriteJSON("hellow")
			// if err != nil {
			// 	log.Printf("Failed to write data to Websocket: %v", err)
			// }
			// return
		// case data, ok := <-datachan:
		// 	if !ok {
		// 		// Done writing, return
		// 		return
		// 	}
		// 	ws.SetWriteDeadline(time.Now().Add(writeWait))
		// 	err := ws.WriteJSON(&data)
		// 	if err != nil {
		// 		log.Printf("Failed to write data to Websocket: %v", err)
		// 		return
		// 	}
		case <-pingticker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			log.Printf("Sending Ping Message to Client")
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("Failed to write ping message to Websocket: %v", err)
				return
			}
		}
	}
}

func LocationTrackingV2(ctx *gin.Context) {
	// done := make(chan struct{})
	w, r := ctx.Writer, ctx.Request
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("upgrader err: %v", err)
		// c.Close()
	}

	//go allUserWebsocketWriter(c, done)
	// go PingResponse(c, done)

	clients = append(clients, c)
	for {
		// Read message from browser
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			log.Printf("failed read message >> %v", err)
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", c.RemoteAddr(), string(msg))

		for _, client := range clients {
			// log.Printf("clients >> %v", client)
			// Write message back to browser
			if err = client.WriteMessage(msgType, msg); err != nil {
				log.Printf("failed send message >> %v", err)
				<-ctx.Done()
				c.Close()
			}
		}
	}
}

func LocationTrackingV3(ctx *gin.Context) {
	// done := make(chan struct{})
	w, r := ctx.Writer, ctx.Request
	_, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("upgrader err: %v", err)
		// c.Close()
	}

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: func(req *http.Request) bool {
					return true
				},
			},
			&ws.Transport{
				CheckOrigin: func(req *http.Request) bool {
					return true
				},
			},
		},
	})
	location.InitChatEndpoints(server)
	go server.Serve()
	defer server.Close()
}

var addr = flag.String("addr", "192.168.176.61:3010", "http service address")

func ListenLocationTracking(gin *gin.Context) {
	id, _ := gin.Params.Get("id")
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: fmt.Sprintf("/api/v1/ws/loc/track/%v", id)}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
