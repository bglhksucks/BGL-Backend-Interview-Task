package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type errorResp struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func getLastPrice(w http.ResponseWriter, r *http.Request) {
	resp, err := lastPrice()
	responseToClient(resp, err, w)
}

func getPriceAtTime(w http.ResponseWriter, r *http.Request) {
	timeStamp := chi.URLParam(r, "time")

	resp, err := priceAtGivenTime(timeStamp)
	responseToClient(resp, err, w)
}

func getAveragePrice(w http.ResponseWriter, r *http.Request) {
	from := chi.URLParam(r, "from")
	to := chi.URLParam(r, "to")

	resp, err := averagePrice(from, to)
	responseToClient(resp, err, w)
}

func responseToClient(resp interface{}, err error, w http.ResponseWriter) {
	if err != nil {
		responseWithError(err, w)
	} else {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			responseWithError(err, w)
		}

		w.Write(respJSON)
	}
}

func responseWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(500)

	respError := errorResp{
		http.StatusText(http.StatusInternalServerError),
		err.Error(),
	}
	respErrorJSON, _ := json.Marshal(respError)

	w.Write(respErrorJSON)
}
