package main

import (
	"github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func onConnection(so socketio.Socket) {
	log.Info("client connected...")

	so.Join("echo")

	so.On("echo", func(msg string) {
		log.Debug("echo")
		log.Infof("echo \"%s\"", msg)

		err := so.Emit("echo", msg, func(so socketio.Socket, data string) {
			log.Debugf("Client ACK with data: ", data)
		})

		if err != nil {
			log.Errorf("error: %s", err)
		}
	})

	so.On("disconnection", func() {
		log.Debug("client disconnected...")
	})
}

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", onConnection)
	server.On("error", func(so socketio.Socket, err error) {
		log.Errorf("error: %s", err)
	})

	http.Handle("/", server)
	log.Println("Serving at localhost:1773...")
	log.Fatal(http.ListenAndServe(":1773", nil))
}
