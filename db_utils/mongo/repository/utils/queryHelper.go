package utils

import (
	"fmt"
	"net/http"

	repoUtils "github.com/ashish036/GO-no-code-api/utils"
)

func ParseQueryParams[T any](paramName QueryParam, queryParams map[QueryParam]interface{}, isRequired bool) (T, repoUtils.Status) {
	var zeroValue T
	param, ok := queryParams[paramName]
	if ok {
		typedParam, ok := param.(T)
		if !ok {
			return zeroValue, repoUtils.Status{
				Code:    http.StatusBadRequest,
				Message: fmt.Sprintf("%s : Invalid parameter type", paramName),
				Error:   fmt.Errorf("%s : Invalid parameter type", paramName),
			}
		}

		return typedParam, repoUtils.Status{}
	}

	if isRequired {
		return zeroValue, repoUtils.Status{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("%s : required parameter is missing", paramName),
			Error:   fmt.Errorf("%s : required parameter is missing", paramName),
		}
	}

	return zeroValue, repoUtils.Status{}
}
