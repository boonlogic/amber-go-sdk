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

// PostSensorReader is a Reader for the PostSensor structure.
type PostSensorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSensorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSensorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSensorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostSensorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostSensorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSensorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostSensorOK creates a PostSensorOK with default headers values
func NewPostSensorOK() *PostSensorOK {
	return &PostSensorOK{}
}

/*
	PostSensorOK describes a response with status code 200, with default header values.

Operation was successful
*/
type PostSensorOK struct {
	Payload *models.PostSensorResponse
}

func (o *PostSensorOK) Error() string {
	return fmt.Sprintf("[POST /sensor][%d] postSensorOK  %+v", 200, o.Payload)
}
func (o *PostSensorOK) GetPayload() *models.PostSensorResponse {
	return o.Payload
}

func (o *PostSensorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostSensorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSensorBadRequest creates a PostSensorBadRequest with default headers values
func NewPostSensorBadRequest() *PostSensorBadRequest {
	return &PostSensorBadRequest{}
}

/*
	PostSensorBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostSensorBadRequest struct {
	Payload *models.Error
}

func (o *PostSensorBadRequest) Error() string {
	return fmt.Sprintf("[POST /sensor][%d] postSensorBadRequest  %+v", 400, o.Payload)
}
func (o *PostSensorBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSensorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSensorUnauthorized creates a PostSensorUnauthorized with default headers values
func NewPostSensorUnauthorized() *PostSensorUnauthorized {
	return &PostSensorUnauthorized{}
}

/*
	PostSensorUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostSensorUnauthorized struct {
	Payload *models.Error
}

func (o *PostSensorUnauthorized) Error() string {
	return fmt.Sprintf("[POST /sensor][%d] postSensorUnauthorized  %+v", 401, o.Payload)
}
func (o *PostSensorUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSensorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSensorNotFound creates a PostSensorNotFound with default headers values
func NewPostSensorNotFound() *PostSensorNotFound {
	return &PostSensorNotFound{}
}

/*
	PostSensorNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type PostSensorNotFound struct {
	Payload *models.Error
}

func (o *PostSensorNotFound) Error() string {
	return fmt.Sprintf("[POST /sensor][%d] postSensorNotFound  %+v", 404, o.Payload)
}
func (o *PostSensorNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSensorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSensorInternalServerError creates a PostSensorInternalServerError with default headers values
func NewPostSensorInternalServerError() *PostSensorInternalServerError {
	return &PostSensorInternalServerError{}
}

/*
	PostSensorInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostSensorInternalServerError struct {
	Payload *models.Error
}

func (o *PostSensorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /sensor][%d] postSensorInternalServerError  %+v", 500, o.Payload)
}
func (o *PostSensorInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostSensorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
