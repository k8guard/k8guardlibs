package rmq

import (
	"fmt"
	"log"
	"net/http"

	armq "github.com/adjust/rmq"
	libs "github.com/k8guard/k8guardlibs"
)

func initStatsHandler() {
	broker := libs.Cfg.RmqBroker
	connection := armq.OpenConnection("k8guard-handler", "tcp", broker, 1)
	http.Handle("/", NewHandler(connection))
	fmt.Printf("Handler listening on http://localhost:3002\n")
	http.ListenAndServe(":3002", nil)
}

type Handler struct {
	connection armq.Connection
}

func NewHandler(connection armq.Connection) *Handler {
	return &Handler{connection: connection}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	layout := request.FormValue("layout")
	refresh := request.FormValue("refresh")

	queues := handler.connection.GetOpenQueues()
	stats := handler.connection.CollectStats(queues)
	log.Printf("queue stats\n%s", stats)
	fmt.Fprint(writer, stats.GetHtml(layout, refresh))
}
