package counter

import (
	"math/rand"

	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

// New ...
func NewRandom(app *model.App) Counter {
	return &randomCounter{}
}

type randomCounter struct {
}

func (rc *randomCounter) MustInit() error {
	return nil
}

func (rc *randomCounter) GetOrderNumber() (int, error) {
	return rand.Intn(1000000), nil
}
