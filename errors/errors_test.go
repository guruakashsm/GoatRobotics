package errors_test

import (
	"testing"

	"github.com/guruakashsm/GoatRobotics/errors"
)

func TestErrorMethods(t *testing.T) {
	// Arrange
	mockErrorDetails := errors.Error{
		Code:    "404",
		Message: "Not Found",
	}

	errorDetailsOutput := mockErrorDetails.Error()         
	errorStringOutput := errorDetailsOutput.Error() 

	expectedErrorString := "CODE : 404 Message : Not Found"
	if errorStringOutput != expectedErrorString {
		t.Errorf("Expected '%s', but got '%s'", expectedErrorString, errorStringOutput)
	}

	expectedErrorDetails := errors.ErrorDetails{
		ErrorDetails: mockErrorDetails,
	} 
	
	if errorDetailsOutput != expectedErrorDetails {
		t.Errorf("Expected ErrorDetails %+v, but got %+v", expectedErrorDetails, errorDetailsOutput)
	}
}
