package server

import (
	"encoding/json"
	"net/http"

	"github.com/sudotoame/devops-practice-1/logic"
)

type Handlers struct {
	Logic *logic.Logic
}

func NewHandlers(logic *logic.Logic) *Handlers {
	return &Handlers{
		Logic: logic,
	}
}

func (h *Handlers) HandleIncrease(w http.ResponseWriter, r *http.Request) {
	h.Logic.Incriment()

	res := h.Logic.NumberOutput()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&res)
}

func (h *Handlers) HandleOutput(w http.ResponseWriter, r *http.Request) {
	res := h.Logic.NumberOutput()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&res)
}

func (h *Handlers) HanldeHealthcheck(w http.ResponseWriter, r *http.Request) {
	res := h.Logic.Healthcheck()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(&res)
}
