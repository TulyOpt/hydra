// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// IntrospectOAuth2TokenReader is a Reader for the IntrospectOAuth2Token structure.
type IntrospectOAuth2TokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntrospectOAuth2TokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntrospectOAuth2TokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIntrospectOAuth2TokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIntrospectOAuth2TokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntrospectOAuth2TokenOK creates a IntrospectOAuth2TokenOK with default headers values
func NewIntrospectOAuth2TokenOK() *IntrospectOAuth2TokenOK {
	return &IntrospectOAuth2TokenOK{}
}

/*IntrospectOAuth2TokenOK handles this case with default header values.

oAuth2TokenIntrospection
*/
type IntrospectOAuth2TokenOK struct {
	Payload *models.OAuth2TokenIntrospection
}

func (o *IntrospectOAuth2TokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/introspect][%d] introspectOAuth2TokenOK  %+v", 200, o.Payload)
}

func (o *IntrospectOAuth2TokenOK) GetPayload() *models.OAuth2TokenIntrospection {
	return o.Payload
}

func (o *IntrospectOAuth2TokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2TokenIntrospection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntrospectOAuth2TokenUnauthorized creates a IntrospectOAuth2TokenUnauthorized with default headers values
func NewIntrospectOAuth2TokenUnauthorized() *IntrospectOAuth2TokenUnauthorized {
	return &IntrospectOAuth2TokenUnauthorized{}
}

/*IntrospectOAuth2TokenUnauthorized handles this case with default header values.

genericError
*/
type IntrospectOAuth2TokenUnauthorized struct {
	Payload *models.GenericError
}

func (o *IntrospectOAuth2TokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /oauth2/introspect][%d] introspectOAuth2TokenUnauthorized  %+v", 401, o.Payload)
}

func (o *IntrospectOAuth2TokenUnauthorized) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *IntrospectOAuth2TokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntrospectOAuth2TokenInternalServerError creates a IntrospectOAuth2TokenInternalServerError with default headers values
func NewIntrospectOAuth2TokenInternalServerError() *IntrospectOAuth2TokenInternalServerError {
	return &IntrospectOAuth2TokenInternalServerError{}
}

/*IntrospectOAuth2TokenInternalServerError handles this case with default header values.

genericError
*/
type IntrospectOAuth2TokenInternalServerError struct {
	Payload *models.GenericError
}

func (o *IntrospectOAuth2TokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/introspect][%d] introspectOAuth2TokenInternalServerError  %+v", 500, o.Payload)
}

func (o *IntrospectOAuth2TokenInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *IntrospectOAuth2TokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
