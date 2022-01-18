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

// PostStreamReader is a Reader for the PostStream structure.
type PostStreamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostStreamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostStreamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostStreamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostStreamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostStreamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostStreamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostStreamServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostStreamOK creates a PostStreamOK with default headers values
func NewPostStreamOK() *PostStreamOK {
	return &PostStreamOK{}
}

/* PostStreamOK describes a response with status code 200, with default header values.

Operation was successful
*/
type PostStreamOK struct {
	Payload *models.PostStreamResponse
}

func (o *PostStreamOK) Error() string {
	return fmt.Sprintf("[POST /stream][%d] postStreamOK  %+v", 200, o.Payload)
}
func (o *PostStreamOK) GetPayload() *models.PostStreamResponse {
	return o.Payload
}

func (o *PostStreamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostStreamResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStreamBadRequest creates a PostStreamBadRequest with default headers values
func NewPostStreamBadRequest() *PostStreamBadRequest {
	return &PostStreamBadRequest{}
}

/* PostStreamBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostStreamBadRequest struct {
	Payload *models.Error
}

func (o *PostStreamBadRequest) Error() string {
	return fmt.Sprintf("[POST /stream][%d] postStreamBadRequest  %+v", 400, o.Payload)
}
func (o *PostStreamBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStreamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStreamUnauthorized creates a PostStreamUnauthorized with default headers values
func NewPostStreamUnauthorized() *PostStreamUnauthorized {
	return &PostStreamUnauthorized{}
}

/* PostStreamUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostStreamUnauthorized struct {
	Payload *models.Error
}

func (o *PostStreamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /stream][%d] postStreamUnauthorized  %+v", 401, o.Payload)
}
func (o *PostStreamUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStreamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStreamNotFound creates a PostStreamNotFound with default headers values
func NewPostStreamNotFound() *PostStreamNotFound {
	return &PostStreamNotFound{}
}

/* PostStreamNotFound describes a response with status code 404, with default header values.

The specified resource was not found
*/
type PostStreamNotFound struct {
	Payload *models.Error
}

func (o *PostStreamNotFound) Error() string {
	return fmt.Sprintf("[POST /stream][%d] postStreamNotFound  %+v", 404, o.Payload)
}
func (o *PostStreamNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStreamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStreamInternalServerError creates a PostStreamInternalServerError with default headers values
func NewPostStreamInternalServerError() *PostStreamInternalServerError {
	return &PostStreamInternalServerError{}
}

/* PostStreamInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostStreamInternalServerError struct {
	Payload *models.Error
}

func (o *PostStreamInternalServerError) Error() string {
	return fmt.Sprintf("[POST /stream][%d] postStreamInternalServerError  %+v", 500, o.Payload)
}
func (o *PostStreamInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStreamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostStreamServiceUnavailable creates a PostStreamServiceUnavailable with default headers values
func NewPostStreamServiceUnavailable() *PostStreamServiceUnavailable {
	return &PostStreamServiceUnavailable{}
}

/* PostStreamServiceUnavailable describes a response with status code 503, with default header values.

Server busy
*/
type PostStreamServiceUnavailable struct {
	Payload *models.Error
}

func (o *PostStreamServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /stream][%d] postStreamServiceUnavailable  %+v", 503, o.Payload)
}
func (o *PostStreamServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostStreamServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}