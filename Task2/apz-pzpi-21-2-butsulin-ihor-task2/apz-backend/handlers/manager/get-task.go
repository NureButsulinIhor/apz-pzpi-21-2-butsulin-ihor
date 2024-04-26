package manager

import (
	"apz-backend/services/task"
	"apz-backend/types"
	"apz-backend/types/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
)

type GetTaskRequestData struct {
	TaskID uint `json:"taskID"`
}

func GetTask(logger *slog.Logger, storage task.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logger.With(
			slog.String("op", "handlers.manager.GetTask"),
			slog.String("requestId", middleware.GetReqID(r.Context())),
		)

		var requestBody GetTaskRequestData
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

		taskModel, err := task.Get(requestBody.TaskID,
			task.Configuration{
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

		render.JSON(w, r, types.Response[models.Task]{
			Status: true,
			Body:   *taskModel,
		})
	}
}
