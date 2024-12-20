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

// GetModelConfigReader is a Reader for the GetModelConfig structure.
type GetModelConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetModelConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModelConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetModelConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetModelConfigOK creates a GetModelConfigOK with default headers values
func NewGetModelConfigOK() *GetModelConfigOK {
	return &GetModelConfigOK{}
}

/*
	GetModelConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetModelConfigOK struct {
	Payload *models.PostConfigResponse
}

func (o *GetModelConfigOK) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/config][%d] getModelConfigOK  %+v", 200, o.Payload)
}
func (o *GetModelConfigOK) GetPayload() *models.PostConfigResponse {
	return o.Payload
}

func (o *GetModelConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelConfigUnauthorized creates a GetModelConfigUnauthorized with default headers values
func NewGetModelConfigUnauthorized() *GetModelConfigUnauthorized {
	return &GetModelConfigUnauthorized{}
}

/*
	GetModelConfigUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetModelConfigUnauthorized struct {
	Payload *models.Error
}

func (o *GetModelConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/config][%d] getModelConfigUnauthorized  %+v", 401, o.Payload)
}
func (o *GetModelConfigUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelConfigNotFound creates a GetModelConfigNotFound with default headers values
func NewGetModelConfigNotFound() *GetModelConfigNotFound {
	return &GetModelConfigNotFound{}
}

/*
	GetModelConfigNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type GetModelConfigNotFound struct {
	Payload *models.Error
}

func (o *GetModelConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/config][%d] getModelConfigNotFound  %+v", 404, o.Payload)
}
func (o *GetModelConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelConfigInternalServerError creates a GetModelConfigInternalServerError with default headers values
func NewGetModelConfigInternalServerError() *GetModelConfigInternalServerError {
	return &GetModelConfigInternalServerError{}
}

/*
	GetModelConfigInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetModelConfigInternalServerError struct {
	Payload *models.Error
}

func (o *GetModelConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /models/{modelId}/config][%d] getModelConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *GetModelConfigInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetModelConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
