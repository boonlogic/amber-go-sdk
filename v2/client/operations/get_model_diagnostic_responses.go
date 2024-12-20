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

// GetModelDiagnosticReader is a Reader for the GetModelDiagnostic structure.
type GetModelDiagnosticReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetModelDiagnosticReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelDiagnosticOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetModelDiagnosticUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModelDiagnosticNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetModelDiagnosticInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetModelDiagnosticOK creates a GetModelDiagnosticOK with default headers values
func NewGetModelDiagnosticOK(writer io.Writer) *GetModelDiagnosticOK {
	return &GetModelDiagnosticOK{

		Payload: writer,
	}
}

/*
	GetModelDiagnosticOK describes a response with status code 200, with default header values.

OK
*/
type GetModelDiagnosticOK struct {
	Payload io.Writer
}

func (o *GetModelDiagnosticOK) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/diagnostic][%d] getModelDiagnosticOK  %+v", 200, o.Payload)
}
func (o *GetModelDiagnosticOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *GetModelDiagnosticOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelDiagnosticUnauthorized creates a GetModelDiagnosticUnauthorized with default headers values
func NewGetModelDiagnosticUnauthorized() *GetModelDiagnosticUnauthorized {
	return &GetModelDiagnosticUnauthorized{}
}

/*
	GetModelDiagnosticUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetModelDiagnosticUnauthorized struct {
	Payload *models.Error
}

func (o *GetModelDiagnosticUnauthorized) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/diagnostic][%d] getModelDiagnosticUnauthorized  %+v", 401, o.Payload)
}
func (o *GetModelDiagnosticUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelDiagnosticUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelDiagnosticNotFound creates a GetModelDiagnosticNotFound with default headers values
func NewGetModelDiagnosticNotFound() *GetModelDiagnosticNotFound {
	return &GetModelDiagnosticNotFound{}
}

/*
	GetModelDiagnosticNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetModelDiagnosticNotFound struct {
	Payload *models.Error
}

func (o *GetModelDiagnosticNotFound) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/diagnostic][%d] getModelDiagnosticNotFound  %+v", 404, o.Payload)
}
func (o *GetModelDiagnosticNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelDiagnosticNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelDiagnosticInternalServerError creates a GetModelDiagnosticInternalServerError with default headers values
func NewGetModelDiagnosticInternalServerError() *GetModelDiagnosticInternalServerError {
	return &GetModelDiagnosticInternalServerError{}
}

/*
	GetModelDiagnosticInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetModelDiagnosticInternalServerError struct {
	Payload *models.Error
}

func (o *GetModelDiagnosticInternalServerError) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/diagnostic][%d] getModelDiagnosticInternalServerError  %+v", 500, o.Payload)
}
func (o *GetModelDiagnosticInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelDiagnosticInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
