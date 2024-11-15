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

// GetVersionReader is a Reader for the GetVersion structure.
type GetVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVersionOK creates a GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

/*
	GetVersionOK describes a response with status code 200, with default header values.

OK
*/
type GetVersionOK struct {
	Payload *models.GetVersionResponse
}

func (o *GetVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] getVersionOK  %+v", 200, o.Payload)
}
func (o *GetVersionOK) GetPayload() *models.GetVersionResponse {
	return o.Payload
}

func (o *GetVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionUnauthorized creates a GetVersionUnauthorized with default headers values
func NewGetVersionUnauthorized() *GetVersionUnauthorized {
	return &GetVersionUnauthorized{}
}

/*
	GetVersionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVersionUnauthorized struct {
	Payload *models.Error
}

func (o *GetVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /version][%d] getVersionUnauthorized  %+v", 401, o.Payload)
}
func (o *GetVersionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionInternalServerError creates a GetVersionInternalServerError with default headers values
func NewGetVersionInternalServerError() *GetVersionInternalServerError {
	return &GetVersionInternalServerError{}
}

/*
	GetVersionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetVersionInternalServerError struct {
	Payload *models.Error
}

func (o *GetVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /version][%d] getVersionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetVersionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}