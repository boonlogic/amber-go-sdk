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

// PostModelDataReader is a Reader for the PostModelData structure.
type PostModelDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostModelDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostModelDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostModelDataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostModelDataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostModelDataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostModelDataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostModelDataOK creates a PostModelDataOK with default headers values
func NewPostModelDataOK() *PostModelDataOK {
	return &PostModelDataOK{}
}

/*
	PostModelDataOK describes a response with status code 200, with default header values.

OK
*/
type PostModelDataOK struct {
	Payload *models.PostDataResponse
}

func (o *PostModelDataOK) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/data][%d] postModelDataOK  %+v", 200, o.Payload)
}
func (o *PostModelDataOK) GetPayload() *models.PostDataResponse {
	return o.Payload
}

func (o *PostModelDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostDataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelDataBadRequest creates a PostModelDataBadRequest with default headers values
func NewPostModelDataBadRequest() *PostModelDataBadRequest {
	return &PostModelDataBadRequest{}
}

/*
	PostModelDataBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostModelDataBadRequest struct {
	Payload *models.Error
}

func (o *PostModelDataBadRequest) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/data][%d] postModelDataBadRequest  %+v", 400, o.Payload)
}
func (o *PostModelDataBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelDataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelDataUnauthorized creates a PostModelDataUnauthorized with default headers values
func NewPostModelDataUnauthorized() *PostModelDataUnauthorized {
	return &PostModelDataUnauthorized{}
}

/*
	PostModelDataUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PostModelDataUnauthorized struct {
	Payload *models.Error
}

func (o *PostModelDataUnauthorized) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/data][%d] postModelDataUnauthorized  %+v", 401, o.Payload)
}
func (o *PostModelDataUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelDataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelDataNotFound creates a PostModelDataNotFound with default headers values
func NewPostModelDataNotFound() *PostModelDataNotFound {
	return &PostModelDataNotFound{}
}

/*
	PostModelDataNotFound describes a response with status code 404, with default header values.

Resource not found
*/
type PostModelDataNotFound struct {
	Payload *models.Error
}

func (o *PostModelDataNotFound) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/data][%d] postModelDataNotFound  %+v", 404, o.Payload)
}
func (o *PostModelDataNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelDataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostModelDataInternalServerError creates a PostModelDataInternalServerError with default headers values
func NewPostModelDataInternalServerError() *PostModelDataInternalServerError {
	return &PostModelDataInternalServerError{}
}

/*
	PostModelDataInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PostModelDataInternalServerError struct {
	Payload *models.Error
}

func (o *PostModelDataInternalServerError) Error() string {
	return fmt.Sprintf("[POST /models/{modelId}/data][%d] postModelDataInternalServerError  %+v", 500, o.Payload)
}
func (o *PostModelDataInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostModelDataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
