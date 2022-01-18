// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostStreamRequest post stream request
//
// swagger:model postStreamRequest
type PostStreamRequest struct {

	// data
	// Required: true
	Data *string `json:"data"`

	// save the sensor after calculation
	// Required: true
	SaveImage *bool `json:"saveImage"`
}

// Validate validates this post stream request
func (m *PostStreamRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSaveImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostStreamRequest) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *PostStreamRequest) validateSaveImage(formats strfmt.Registry) error {

	if err := validate.Required("saveImage", "body", m.SaveImage); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post stream request based on context it is used
func (m *PostStreamRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostStreamRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostStreamRequest) UnmarshalBinary(b []byte) error {
	var res PostStreamRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}