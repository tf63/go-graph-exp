package resolver

import (
	"errors"

	"github.com/tf63/go-graph-exp/internal/entity"
)

func TodoErrorHandler(operation string, err error) error {
	if err == entity.STATUS_NOT_FOUND {
		res := errors.New("[" + operation + "] expense not found")
		return res
	} else if err == entity.STATUS_SERVICE_UNAVAILABLE {
		res := errors.New("[" + operation + "] internal server error")
		return res
	} else if err != nil {
		res := errors.New("[" + operation + "] enexpected error")
		return res
	} else {
		res := errors.New("not implemented")
		return res
	}
}
