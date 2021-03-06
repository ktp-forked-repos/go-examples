package requesterror

import "fmt"

type CustomUserError struct {
	errorMessage string
	errorCode    int64
}

func (s CustomUserError) Error() string {
	return fmt.Sprintf("Message %q, code %d", s.errorMessage, s.errorCode)
}

func InvalidRequest() CustomUserError {
	customErrorMessage := "Invalid Request"
	return CustomUserError{errorMessage: customErrorMessage, errorCode: 404}
}

func MalformedRequest() CustomUserError {
	customErrorMessage := "Malformed request rejected"
	return CustomUserError{errorMessage: customErrorMessage, errorCode: 400}
}
