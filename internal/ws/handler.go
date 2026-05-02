package ws

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
	"tracking.xlkv.com/internal/domain"
	"tracking.xlkv.com/internal/response"
)

type Cache interface {
	SetLocation(ctx context.Context, vehicleID int64, location domain.Location) error
	PublishLocation(ctx context.Context, vehicleID int64, location domain.Location) error
	SubscribeLocation(ctx context.Context, vehicleID int64) *redis.PubSub // ← qo'sh
}

type VehicleWSHandler struct {
	cache Cache
}

func NewVehicleWSHandler(cache Cache) *VehicleWSHandler {
	return &VehicleWSHandler{cache: cache}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (wH *VehicleWSHandler) HandleVehicleWS(w http.ResponseWriter, r *http.Request) {

	strVehicleID := chi.URLParam(r, "vehicleID")

	if strVehicleID == "" {
		response.Error(w, http.StatusBadRequest, "invalid vehicle id")
		return
	}

	vehicleID, err := strconv.Atoi(strVehicleID)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid vehicle id")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		slog.Error("websocket upgrade error", "err", err)
		return
	}

	pubSub := wH.cache.SubscribeLocation(r.Context(), int64(vehicleID))

	defer pubSub.Close()
	defer conn.Close()

	for msg := range pubSub.Channel() {
		conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	}
}
