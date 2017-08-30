package rmq

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adjust/rmq"
	libs "github.com/k8guard/k8guardlibs"
)

func initHandler() {
	broker := libs.Cfg.RmqBroker
	connection := rmq.OpenConnection("k8guard-handler", "tcp", broker, 1)
	http.Handle("/rmq", NewHandler(connection))
	fmt.Printf("Handler listening on http://localhost:3002/rmq\n")
	http.ListenAndServe(":3002", nil)
}

type Handler struct {
	connection rmq.Connection
}

func NewHandler(connection rmq.Connection) *Handler {
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
