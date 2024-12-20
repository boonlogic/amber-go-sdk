// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/boonlogic/amber-go-sdk/v1/models"
)

// PutConfigReader is a Reader for the PutConfig structure.
type PutConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConfigServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConfigOK creates a PutConfigOK with default headers values
func NewPutConfigOK() *PutConfigOK {
	return &PutConfigOK{}
}

/*
	PutConfigOK describes a response with status code 200, with default header values.

Operation was successful
*/
type PutConfigOK struct {
	Payload *models.PutConfigResponse
}

func (o *PutConfigOK) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigOK  %+v", 200, o.Payload)
}
func (o *PutConfigOK) GetPayload() *models.PutConfigResponse {
	return o.Payload
}

func (o *PutConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigBadRequest creates a PutConfigBadRequest with default headers values
func NewPutConfigBadRequest() *PutConfigBadRequest {
	return &PutConfigBadRequest{}
}

/*
	PutConfigBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PutConfigBadRequest struct {
	Payload *models.Error
}

func (o *PutConfigBadRequest) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigBadRequest  %+v", 400, o.Payload)
}
func (o *PutConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigUnauthorized creates a PutConfigUnauthorized with default headers values
func NewPutConfigUnauthorized() *PutConfigUnauthorized {
	return &PutConfigUnauthorized{}
}

/*
	PutConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PutConfigUnauthorized struct {
	Payload *models.Error
}

func (o *PutConfigUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *PutConfigUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigNotFound creates a PutConfigNotFound with default headers values
func NewPutConfigNotFound() *PutConfigNotFound {
	return &PutConfigNotFound{}
}

/*
	PutConfigNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type PutConfigNotFound struct {
	Payload *models.Error
}

func (o *PutConfigNotFound) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigNotFound  %+v", 404, o.Payload)
}
func (o *PutConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigInternalServerError creates a PutConfigInternalServerError with default headers values
func NewPutConfigInternalServerError() *PutConfigInternalServerError {
	return &PutConfigInternalServerError{}
}

/*
	PutConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PutConfigInternalServerError struct {
	Payload *models.Error
}

func (o *PutConfigInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *PutConfigInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConfigServiceUnavailable creates a PutConfigServiceUnavailable with default headers values
func NewPutConfigServiceUnavailable() *PutConfigServiceUnavailable {
	return &PutConfigServiceUnavailable{}
}

/*
	PutConfigServiceUnavailable describes a response with status code 503, with default header values.

Server busy
*/
type PutConfigServiceUnavailable struct {
	Payload *models.Error
}

func (o *PutConfigServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /config][%d] putConfigServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PutConfigServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutConfigServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
