// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"amber-go-sdk/ambergen/models"
)

// GetConfigReader is a Reader for the GetConfig structure.
type GetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigOK creates a GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {
	return &GetConfigOK{}
}

/* GetConfigOK describes a response with status code 200, with default header values.

Operation was successful
*/
type GetConfigOK struct {
	Payload *models.GetConfigResponse
}

func (o *GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigOK  %+v", 200, o.Payload)
}
func (o *GetConfigOK) GetPayload() *models.GetConfigResponse {
	return o.Payload
}

func (o *GetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigBadRequest creates a GetConfigBadRequest with default headers values
func NewGetConfigBadRequest() *GetConfigBadRequest {
	return &GetConfigBadRequest{}
}

/* GetConfigBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetConfigBadRequest struct {
	Payload *models.Error
}

func (o *GetConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigBadRequest  %+v", 400, o.Payload)
}
func (o *GetConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigUnauthorized creates a GetConfigUnauthorized with default headers values
func NewGetConfigUnauthorized() *GetConfigUnauthorized {
	return &GetConfigUnauthorized{}
}

/* GetConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetConfigUnauthorized struct {
	Payload *models.Error
}

func (o *GetConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *GetConfigUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigNotFound creates a GetConfigNotFound with default headers values
func NewGetConfigNotFound() *GetConfigNotFound {
	return &GetConfigNotFound{}
}

/* GetConfigNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type GetConfigNotFound struct {
	Payload *models.Error
}

func (o *GetConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigNotFound  %+v", 404, o.Payload)
}
func (o *GetConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigInternalServerError creates a GetConfigInternalServerError with default headers values
func NewGetConfigInternalServerError() *GetConfigInternalServerError {
	return &GetConfigInternalServerError{}
}

/* GetConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetConfigInternalServerError struct {
	Payload *models.Error
}

func (o *GetConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *GetConfigInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
