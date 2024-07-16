package httpstatuses

import "net/http"

type HTTPStatus struct {
	Code    int
	Message string
}

var (
	Processing = HTTPStatus{
		Code:    http.StatusProcessing,
		Message: "Processing",
	}
	BadRequest = HTTPStatus{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}
	Unauthorized = HTTPStatus{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	PaymentRequired = HTTPStatus{
		Code:    http.StatusPaymentRequired,
		Message: "Payment Required",
	}
	Forbidden = HTTPStatus{
		Code:    http.StatusForbidden,
		Message: "Forbidden",
	}
	NotFound = HTTPStatus{
		Code:    http.StatusNotFound,
		Message: "Not Found",
	}
	MethodNotAllowed = HTTPStatus{
		Code:    http.StatusMethodNotAllowed,
		Message: "Method Not Allowed",
	}
	InternalServerError = HTTPStatus{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
)
