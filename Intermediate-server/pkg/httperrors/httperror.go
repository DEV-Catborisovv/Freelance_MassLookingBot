package httperrors

import "net/http"

type HTTPError struct {
	Code    int
	Message string
}

var (
	BadRequest = HTTPError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}
	Unauthorized = HTTPError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	PaymentRequired = HTTPError{
		Code:    http.StatusPaymentRequired,
		Message: "Payment Required",
	}
	Forbidden = HTTPError{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}
	NotFound = HTTPError{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
	MethodNotAllowed = HTTPError{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method Not Allowed",
	}
	InternalServerError = HTTPError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
)
