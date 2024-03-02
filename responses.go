package lambda

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	log "github.com/33cn/chain33/common/log/log15"
)

// NewResponse creates a new responses object with the pointer type fields initialised
func NewResponse() events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		Headers:           map[string]string{},
		MultiValueHeaders: map[string][]string{},
		Cookies:           []string{},
	}
}

// InternalErrorResponse returns a response object for an internal error.  The `message` and `err`
// parameters are written to the log with Error level. The error return value is *always* nil to satisfy the requirements
// of the lambda handler. The body is set to "Internal Server Error"
// Any additional values in logValues are written to the log after the initial values
func InternalErrorResponseWithLog(message string, err error, logger log.Logger, logValues ...interface{}) (events.APIGatewayV2HTTPResponse, error) {
	return ErrorResponse(message, "", err, http.StatusInternalServerError, logger, logValues...)
}

// JSONResponse returns a response object with the content type set to "application/json" and the
// string value as the body. The error return value is *always* nil to satisfy the requirements
// of the lambda handler.
func JSONResponse(body string, status int) (events.APIGatewayV2HTTPResponse, error) {
	return ResponseWithType(body, "application/json", status)
}

// ResponseWithType returns a response object with the body and content-type set with the
// required status
func ResponseWithType(body, contentType string, status int) (events.APIGatewayV2HTTPResponse, error) {
	response := NewResponse()
	response.Headers["Content-Type"] = contentType
	response.StatusCode = status
	response.Body = body
	return response, nil
}

// ErrorResponse returns a response object with a user defined error. The messane and error
// are written to the log at Error level. The error return value is *always* nil to satisfy the requirements
// of the lambda handler. The value of the `body` parameter is used as the body of the response. If this
// an empty string, the string value of the error will be used instead.
// Any additional values in logValues are written to the log after the initial values
func ErrorResponse(message, body string, err error, status int, logger log.Logger, logValues ...interface{}) (events.APIGatewayV2HTTPResponse, error) {
	toLog := append([]interface{}{"error", err}, logValues...)
	logger.Error(message, toLog...)
	response := NewResponse()
	response.StatusCode = status
	if body == "" {
		response.Body = err.Error()
	} else {
		response.Body = body
	}

	return response, nil
}

// NotFoundResponse returns a response object set to NotFound. The `message“ and `what` parameters
// are written to the log at Warning level. The error return value is *always* nil to satisfy the requirements
// of the lambda handler. Any additional values in logValues are written to the log after the initial values
func NotFoundResponse(message, what string, logger log.Logger, logValues ...interface{}) (events.APIGatewayV2HTTPResponse, error) {
	toLog := append([]interface{}{"what", what}, logValues...)
	logger.Warn(message, toLog...)
	response := NewResponse()
	response.StatusCode = http.StatusNotFound
	response.Body = fmt.Sprintf("%s not found", what)
	return response, nil
}

// ForbiddenResponse returns a response object set to Forbidden. The `what` parameter can be set
// to the object access is denied to or the user or whatever is convenient. The `message` and `what` parameters
// are written to the log at Warning level. The error return value is *always* nil to satisfy the requirements
// of the lambda handler. Any additional values in logValues are written to the log after the initial values
func ForbiddenResponse(message, what string, logger log.Logger, logValues ...interface{}) (events.APIGatewayV2HTTPResponse, error) {
	toLog := append([]interface{}{"what", what}, logValues...)
	logger.Warn(message, toLog...)
	response := NewResponse()
	response.StatusCode = http.StatusForbidden
	response.Body = "Access Denied"
	return response, nil
}

// ConflictResponse
