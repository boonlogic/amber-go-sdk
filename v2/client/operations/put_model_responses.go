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

// PutModelReader is a Reader for the PutModel structure.
type PutModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutModelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutModelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutModelOK creates a PutModelOK with default headers values
func NewPutModelOK() *PutModelOK {
	return &PutModelOK{}
}

/*
	PutModelOK describes a response with status code 200, with default header values.

OK
*/
type PutModelOK struct {
	Payload *models.PostModelResponse
}

func (o *PutModelOK) Error() string {
	return fmt.Sprintf("[PUT /models/{modelId}][%d] putModelOK  %+v", 200, o.Payload)
}
func (o *PutModelOK) GetPayload() *models.PostModelResponse {
	return o.Payload
}

func (o *PutModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostModelResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutModelBadRequest creates a PutModelBadRequest with default headers values
func NewPutModelBadRequest() *PutModelBadRequest {
	return &PutModelBadRequest{}
}

/*
	PutModelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutModelBadRequest struct {
	Payload *models.Error
}

func (o *PutModelBadRequest) Error() string {
	return fmt.Sprintf("[PUT /models/{modelId}][%d] putModelBadRequest  %+v", 400, o.Payload)
}
func (o *PutModelBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutModelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutModelUnauthorized creates a PutModelUnauthorized with default headers values
func NewPutModelUnauthorized() *PutModelUnauthorized {
	return &PutModelUnauthorized{}
}

/*
	PutModelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutModelUnauthorized struct {
	Payload *models.Error
}

func (o *PutModelUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /models/{modelId}][%d] putModelUnauthorized  %+v", 401, o.Payload)
}
func (o *PutModelUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutModelNotFound creates a PutModelNotFound with default headers values
func NewPutModelNotFound() *PutModelNotFound {
	return &PutModelNotFound{}
}

/*
	PutModelNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type PutModelNotFound struct {
	Payload *models.Error
}

func (o *PutModelNotFound) Error() string {
	return fmt.Sprintf("[PUT /models/{modelId}][%d] putModelNotFound  %+v", 404, o.Payload)
}
func (o *PutModelNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutModelInternalServerError creates a PutModelInternalServerError with default headers values
func NewPutModelInternalServerError() *PutModelInternalServerError {
	return &PutModelInternalServerError{}
}

/*
	PutModelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutModelInternalServerError struct {
	Payload *models.Error
}

func (o *PutModelInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /models/{modelId}][%d] putModelInternalServerError  %+v", 500, o.Payload)
}
func (o *PutModelInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutModelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}