package manager

import (
	"apz-backend/services/transfer"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetTransfersRequestBody struct {
	CarID uint `json:"carID"`
}

type GetTransfersResponse struct {
	Transfers []models.Transfer `json:"transfers"`
}

func GetTransfers(logger *slog.Logger, storage transfer.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.GetTransfers"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody GetTransfersRequestBody
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

		transfers, err := transfer.GetAll(requestBody.CarID, transfer.Configuration{
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

		render.JSON(w, r, types.Response[GetTransfersResponse]{
			Status: true,
			Body:   GetTransfersResponse{Transfers: transfers},
		})
	}
}
