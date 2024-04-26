package car

import (
	"context"
	"log/slog"
)

type Storage interface {
	StorageAdder
	Adder
	UserGetter
	Deleter
	Getter
	TransfersByCarGetter
	ManagerGetter
}

type Configuration struct {
	Storage Storage
	Logger  *slog.Logger
	Context context.Context
}
