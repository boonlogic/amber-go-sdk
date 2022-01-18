// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/boonlogic/amber-go-sdk/models"
)

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/* GetStatusOK describes a response with status code 200, with default header values.

Operation was successful
*/
type GetStatusOK struct {
	Payload *models.GetStatusResponse
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusOK  %+v", 200, o.Payload)
}
func (o *GetStatusOK) GetPayload() *models.GetStatusResponse {
	return o.Payload
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusBadRequest creates a GetStatusBadRequest with default headers values
func NewGetStatusBadRequest() *GetStatusBadRequest {
	return &GetStatusBadRequest{}
}

/* GetStatusBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetStatusBadRequest struct {
	Payload *models.Error
}

func (o *GetStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusBadRequest  %+v", 400, o.Payload)
}
func (o *GetStatusBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusUnauthorized creates a GetStatusUnauthorized with default headers values
func NewGetStatusUnauthorized() *GetStatusUnauthorized {
	return &GetStatusUnauthorized{}
}

/* GetStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetStatusUnauthorized struct {
	Payload *models.Error
}

func (o *GetStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusUnauthorized  %+v", 401, o.Payload)
}
func (o *GetStatusUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusNotFound creates a GetStatusNotFound with default headers values
func NewGetStatusNotFound() *GetStatusNotFound {
	return &GetStatusNotFound{}
}

/* GetStatusNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetStatusNotFound struct {
	Payload *models.Error
}

func (o *GetStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusNotFound  %+v", 404, o.Payload)
}
func (o *GetStatusNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusInternalServerError creates a GetStatusInternalServerError with default headers values
func NewGetStatusInternalServerError() *GetStatusInternalServerError {
	return &GetStatusInternalServerError{}
}

/* GetStatusInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetStatusInternalServerError struct {
	Payload *models.Error
}

func (o *GetStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /status][%d] getStatusInternalServerError  %+v", 500, o.Payload)
}
func (o *GetStatusInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}