package app

import (
	"encoding/json"
	"flink/internal/data/model"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) updateLocationHistory(w http.ResponseWriter, r *http.Request) {

	locationData := new(model.LocationData)
	if err := json.NewDecoder(r.Body).Decode(locationData); err != nil {
		log.Println("updateLocationHistory: failed to decode json:", err)
		return
	}
	orderId := chi.URLParam(r, "order_id")
	_, ok := h.Data[orderId]
	if ok {
		h.Data[orderId] = append(h.Data[orderId], model.LocationData{
			Latitude:  locationData.Latitude,
			Longitude: locationData.Longitude,
		})
	} else {
		h.Data[orderId] = []model.LocationData{
			{
				Latitude:  locationData.Latitude,
				Longitude: locationData.Longitude,
			},
		}
	}

	SendResponse(w, http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

func (h *Handler) getLocationHistory(w http.ResponseWriter, r *http.Request) {

	orderId := chi.URLParam(r, "order_id")
	var (
		response []model.LocationData
		max      int
		err      error
	)

	maxString := r.URL.Query().Get("max")
	if len(maxString) == 0 {
		max = 0
	} else {
		max, err = strconv.Atoi(maxString)
		if err != nil {
			log.Println("getLocationHistory: invalid max value:", err)
			SendResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	locationHistory := h.Data[orderId]
	if locationHistory == nil {
		log.Println("getLocationHistory: Location not found!")
		SendResponse(w, http.StatusNotFound, "Location not found!")
		return
	}

	if max == 0 {
		response = locationHistory
	} else {
		if max > len(locationHistory) {
			max = len(locationHistory)
		}
		response = locationHistory[len(locationHistory)-max:]
	}

	SendResponse(w, http.StatusOK, map[string]interface{}{
		"order_id": orderId,
		"history":  response,
	})
}

func (h *Handler) deleteLocationHistory(w http.ResponseWriter, r *http.Request) {

	orderId := chi.URLParam(r, "order_id")
	delete(h.Data, orderId)

	SendResponse(w, http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
