package warehouse

import (
	"context"
	"log/slog"
)

type Storage interface {
	StorageAdder
	Adder
	UserGetter
	ManagerAdder
	Deleter
	Getter
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
