package services

import (
	"errors"

	"golang.org/x/net/context"
)

// Tax service offers 2 methods to add and substract tax value from
// the given value
//go:generate counterfeiter . Tax
type Tax interface {
	// Add will add tax to the given value and return it
	Add(ctx context.Context, value float64) (float64, error)
	// Sub will remove the tax value from the given value and return it
	Sub(ctx context.Context, value float64) (float64, error)
}

var ErrNegativeInput = errors.New("negative input")

type tax struct {
	percentage float64
}

// NewTax returns a new Tax service implementation with the pre-configured
// tax value set
func NewTax(taxPercentage float64) Tax {
	return tax{percentage: taxPercentage}
}

func (t tax) Add(_ context.Context, value float64) (float64, error) {
	if value < 0 {
		return 0, ErrNegativeInput
	}

	return value + value*t.percentage, nil
}

func (t tax) Sub(_ context.Context, value float64) (float64, error) {
	if value < 0 {
		return 0, ErrNegativeInput
	}

	return value / (1 + t.percentage), nil
}
