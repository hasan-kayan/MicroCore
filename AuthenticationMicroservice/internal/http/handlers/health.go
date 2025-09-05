package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/domain"
	"github.com/hasan-kayan/MicroCore/AuthenticationMicroservice/internal/core/ports"
)

type HealthHandler struct {
	Svc ports.HealthService
}

type storeResp struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (h *HealthHandler) PostHealth(w http.ResponseWriter, r *http.Request) {
	var p domain.HealthPayload
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	info := ports.RequestContext{
		RemoteIP:  remoteIP(r),
		RequestID: r.Header.Get("X-Request-ID"),
		UserAgent: r.UserAgent(),
	}

	id, err := h.Svc.Record(context.Background(), p, info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(storeResp{ID: id, Status: "stored"})
}

func remoteIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	ip := r.RemoteAddr
	return ip
}
