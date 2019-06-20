package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kargo-api/processor"
)

type Handler struct {
	proc processor.Processor
}

func NewHandler() Handler {
	return Handler{
		proc: processor.NewProcessor(),
	}
}

func (handler *Handler) GetSortedJobsHandler(w http.ResponseWriter, r *http.Request) {
	sortedJobs, err := handler.proc.GetSortedJobs()
	if err != nil {
		//response err
		errorMsg := err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorJSON, _ := json.Marshal(map[string]interface{}{"errorMsg": errorMsg})
		w.Write([]byte(errorJSON))
		return
	}
	bodyJSON, _ := json.Marshal(sortedJobs)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bodyJSON))
}

func (handler *Handler) GetSortedBidsHandler(w http.ResponseWriter, r *http.Request) {
	sortedBids, err := handler.proc.GetSortedBids(int64(1))
	if err != nil {
		//response err
		errorMsg := err.Error()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		errorJSON, _ := json.Marshal(map[string]interface{}{"errorMsg": errorMsg})
		w.Write([]byte(errorJSON))
		return
	}
	bodyJSON, _ := json.Marshal(sortedBids)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(bodyJSON))
}
