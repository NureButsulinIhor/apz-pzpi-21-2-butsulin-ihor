package admin

import (
	"apz-backend/services/slot"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetSlotRequestData struct {
	SlotID uint `json:"slotID"`
}

func GetSlot(logger *slog.Logger, storage slot.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.admin.GetSlot"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody GetSlotRequestData
		err := render.DecodeJSON(r.Body, &requestBody)
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  "invalid request body",
				Body:   nil,
			})
			l.Debug("err in decoding json")
			return
		}

		slotModel, err := slot.Get(requestBody.SlotID,
			slot.Configuration{
				Logger:  l,
				Storage: storage,
				Context: r.Context(),
			})
		if err != nil {
			w.WriteHeader(400)
			render.JSON(w, r, types.Response[any]{
				Status: false,
				Error:  err.Error(),
				Body:   nil,
			})
			return
		}

		render.JSON(w, r, types.Response[models.Slot]{
			Status: true,
			Body:   *slotModel,
		})
	}
}
