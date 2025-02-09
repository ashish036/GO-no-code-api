package utils

import (
	"context"
	"errors"
	"net/http"
	"strings"

	repoUtils "github.com/ashish036/GO-no-code-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DOC_NOT_FOUND_ERROR = "mongo: no documents in result"
	DUPLICATE_KEY_ERROR = "E11000 duplicate key error collection"
)

func ParseMongoError(err error) repoUtils.Status {
	if err == nil {
		return repoUtils.Status{}
	}

	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return repoUtils.Status{
			Code:    http.StatusNotFound,
			Message: "No documents found",
			Error:   err,
		}

	case errors.Is(err, mongo.ErrClientDisconnected):
		return repoUtils.Status{
			Code:    http.StatusServiceUnavailable,
			Message: "MongoDB client is disconnected",
			Error:   err,
		}

	case errors.Is(err, mongo.ErrNilDocument):
		return repoUtils.Status{
			Code:    http.StatusBadRequest,
			Message: "Attempted to insert a nil document",
			Error:   err,
		}

	case errors.Is(err, mongo.ErrUnacknowledgedWrite):
		return repoUtils.Status{
			Code:    http.StatusInternalServerError,
			Message: "Write operation was not acknowledged",
			Error:   err,
		}

	case strings.Contains(err.Error(), "duplicate key error"):
		return repoUtils.Status{
			Code:    http.StatusConflict,
			Message: "Duplicate key violation",
			Error:   err,
		}

	case strings.Contains(err.Error(), "network error"):
		return repoUtils.Status{
			Code:    http.StatusServiceUnavailable,
			Message: "Network error while communicating with MongoDB",
			Error:   err,
		}

	case errors.Is(err, context.DeadlineExceeded):
		return repoUtils.Status{
			Code:    http.StatusGatewayTimeout,
			Message: "MongoDB operation timed out",
			Error:   err,
		}

	case errors.Is(err, context.Canceled):
		return repoUtils.Status{
			Code:    http.StatusRequestTimeout,
			Message: "MongoDB request was canceled",
			Error:   err,
		}

	default:
		return repoUtils.Status{
			Code:    http.StatusInternalServerError,
			Message: "An unknown MongoDB error occurred",
			Error:   err,
		}
	}
}
