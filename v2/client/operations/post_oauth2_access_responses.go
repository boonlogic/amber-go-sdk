// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/boonlogic/amber-go-sdk/v2/models"
)

// PostOauth2AccessReader is a Reader for the PostOauth2Access structure.
type PostOauth2AccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauth2AccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOauth2AccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOauth2AccessBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOauth2AccessUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOauth2AccessInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOauth2AccessOK creates a PostOauth2AccessOK with default headers values
func NewPostOauth2AccessOK() *PostOauth2AccessOK {
	return &PostOauth2AccessOK{}
}

/*
	PostOauth2AccessOK describes a response with status code 200, with default header values.

OK
*/
type PostOauth2AccessOK struct {
	Payload *models.PostOauth2AccessResponse
}

func (o *PostOauth2AccessOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/access][%d] postOauth2AccessOK  %+v", 200, o.Payload)
}
func (o *PostOauth2AccessOK) GetPayload() *models.PostOauth2AccessResponse {
	return o.Payload
}

func (o *PostOauth2AccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostOauth2AccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauth2AccessBadRequest creates a PostOauth2AccessBadRequest with default headers values
func NewPostOauth2AccessBadRequest() *PostOauth2AccessBadRequest {
	return &PostOauth2AccessBadRequest{}
}

/*
	PostOauth2AccessBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostOauth2AccessBadRequest struct {
	Payload *models.Error
}

func (o *PostOauth2AccessBadRequest) Error() string {
	return fmt.Sprintf("[POST /oauth2/access][%d] postOauth2AccessBadRequest  %+v", 400, o.Payload)
}
func (o *PostOauth2AccessBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostOauth2AccessBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauth2AccessUnauthorized creates a PostOauth2AccessUnauthorized with default headers values
func NewPostOauth2AccessUnauthorized() *PostOauth2AccessUnauthorized {
	return &PostOauth2AccessUnauthorized{}
}

/*
	PostOauth2AccessUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostOauth2AccessUnauthorized struct {
	Payload *models.Error
}

func (o *PostOauth2AccessUnauthorized) Error() string {
	return fmt.Sprintf("[POST /oauth2/access][%d] postOauth2AccessUnauthorized  %+v", 401, o.Payload)
}
func (o *PostOauth2AccessUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostOauth2AccessUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauth2AccessInternalServerError creates a PostOauth2AccessInternalServerError with default headers values
func NewPostOauth2AccessInternalServerError() *PostOauth2AccessInternalServerError {
	return &PostOauth2AccessInternalServerError{}
}

/*
	PostOauth2AccessInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostOauth2AccessInternalServerError struct {
	Payload *models.Error
}

func (o *PostOauth2AccessInternalServerError) Error() string {
	return fmt.Sprintf("[POST /oauth2/access][%d] postOauth2AccessInternalServerError  %+v", 500, o.Payload)
}
func (o *PostOauth2AccessInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostOauth2AccessInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}