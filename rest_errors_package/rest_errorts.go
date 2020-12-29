package rest_errors_package
import "net/http"

type RestErr struct {
	Message string
	Code int
	Error string
	Causes []interface{}
}

func NewBadRequestError(message string)(*RestErr) {
	return &RestErr{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}


func NewNotFoundError(message string)(*RestErr) {
	return &RestErr{
		Message: message,
		Code: http.StatusNotFound,
		Error: "not_found",
	}
}


func NewInternalServerError(message string, err error) (*RestErr){
	result := &RestErr{
		Message: message,
		Code: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
	if result!= nil {
		result.Causes = append(result.Causes, err.Error())
	}
	return result
}

func NewUniqueContraintViolationcompany_name_listing_active_uniqueError(message string) (*RestErr) {
	return &RestErr{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "bad_request",
	}
}

