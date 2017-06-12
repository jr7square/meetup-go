package meetup

import (
  "fmt"
)

type ErrorDetail struct {
  Code string `json:"code"`
  Message string `json:"message"`
  Problem string `json:"problem"`
}

type ApiError struct {
  Errors []ErrorDetail `json:"errors"`
}

func (e ApiError) Error() string {
	if len(e.Errors) > 0 {
		err := e.Errors[0]
		return fmt.Sprintf("twitter: %d %v", err.Code, err.Message)
	}
	return ""
}

func (e ApiError) Empty() bool {
	if len(e.Errors) == 0 {
		return true
	}
	return false
}

func relevantError(httpError error, apiError ApiError) error {
	if httpError != nil {
		return httpError
	}
	if apiError.Empty() {
		return nil
	}
	return apiError
}
