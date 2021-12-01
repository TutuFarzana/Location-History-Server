package app

import (
	"encoding/json"
	"flink/internal/data/model"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Handler is a top level structure
type Handler struct {
	Port string
	Data map[string][]model.LocationData
}

// NewHandler is a factory method
func NewHandler(port string) *Handler {

	return &Handler{
		Port: port,
		Data: make(map[string][]model.LocationData),
	}
}

// Serve starts a new http server.
func (h *Handler) StartServer(handler http.Handler) {

	if !strings.HasPrefix(h.Port, ":") {
		h.Port = ":" + h.Port
	}
	server := http.Server{
		Addr:    h.Port,
		Handler: handler,
	}

	go func() {
		log.Printf("Listens to %s...\n", server.Addr)
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal("Server failed: ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}

// SendResponse sends json API responses.
func SendResponse(w http.ResponseWriter, status int, result interface{}) {

	j, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
