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

// GetModelRootCauseReader is a Reader for the GetModelRootCause structure.
type GetModelRootCauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelRootCauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelRootCauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetModelRootCauseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetModelRootCauseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModelRootCauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetModelRootCauseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetModelRootCauseOK creates a GetModelRootCauseOK with default headers values
func NewGetModelRootCauseOK() *GetModelRootCauseOK {
	return &GetModelRootCauseOK{}
}

/*
	GetModelRootCauseOK describes a response with status code 200, with default header values.

OK
*/
type GetModelRootCauseOK struct {
	Payload *models.GetRootCauseResponse
}

func (o *GetModelRootCauseOK) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/rootCause][%d] getModelRootCauseOK  %+v", 200, o.Payload)
}
func (o *GetModelRootCauseOK) GetPayload() *models.GetRootCauseResponse {
	return o.Payload
}

func (o *GetModelRootCauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRootCauseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRootCauseBadRequest creates a GetModelRootCauseBadRequest with default headers values
func NewGetModelRootCauseBadRequest() *GetModelRootCauseBadRequest {
	return &GetModelRootCauseBadRequest{}
}

/*
	GetModelRootCauseBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetModelRootCauseBadRequest struct {
	Payload *models.Error
}

func (o *GetModelRootCauseBadRequest) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/rootCause][%d] getModelRootCauseBadRequest  %+v", 400, o.Payload)
}
func (o *GetModelRootCauseBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelRootCauseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRootCauseUnauthorized creates a GetModelRootCauseUnauthorized with default headers values
func NewGetModelRootCauseUnauthorized() *GetModelRootCauseUnauthorized {
	return &GetModelRootCauseUnauthorized{}
}

/*
	GetModelRootCauseUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetModelRootCauseUnauthorized struct {
	Payload *models.Error
}

func (o *GetModelRootCauseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/rootCause][%d] getModelRootCauseUnauthorized  %+v", 401, o.Payload)
}
func (o *GetModelRootCauseUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelRootCauseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRootCauseNotFound creates a GetModelRootCauseNotFound with default headers values
func NewGetModelRootCauseNotFound() *GetModelRootCauseNotFound {
	return &GetModelRootCauseNotFound{}
}

/*
	GetModelRootCauseNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetModelRootCauseNotFound struct {
	Payload *models.Error
}

func (o *GetModelRootCauseNotFound) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/rootCause][%d] getModelRootCauseNotFound  %+v", 404, o.Payload)
}
func (o *GetModelRootCauseNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelRootCauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRootCauseInternalServerError creates a GetModelRootCauseInternalServerError with default headers values
func NewGetModelRootCauseInternalServerError() *GetModelRootCauseInternalServerError {
	return &GetModelRootCauseInternalServerError{}
}

/*
	GetModelRootCauseInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetModelRootCauseInternalServerError struct {
	Payload *models.Error
}

func (o *GetModelRootCauseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/rootCause][%d] getModelRootCauseInternalServerError  %+v", 500, o.Payload)
}
func (o *GetModelRootCauseInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelRootCauseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}