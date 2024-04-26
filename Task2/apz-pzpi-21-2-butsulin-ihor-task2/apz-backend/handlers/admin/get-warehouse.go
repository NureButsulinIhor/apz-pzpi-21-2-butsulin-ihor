package admin

import (
	"apz-backend/services/warehouse"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetWarehouseRequestData struct {
	WarehouseID uint `json:"warehouseID"`
}

func GetWarehouse(logger *slog.Logger, storage warehouse.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.admin.GetWarehouse"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody GetWarehouseRequestData
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

		warehouseModel, err := warehouse.Get(requestBody.WarehouseID,
			warehouse.Configuration{
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

		render.JSON(w, r, types.Response[models.Warehouse]{
			Status: true,
			Body:   *warehouseModel,
		})
	}
}
