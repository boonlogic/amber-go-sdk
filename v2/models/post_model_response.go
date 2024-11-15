// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostModelResponse post model response
//
// swagger:model postModelResponse
type PostModelResponse struct {
	Model
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PostModelResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Model
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Model = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PostModelResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Model)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post model response
func (m *PostModelResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Model
	if err := m.Model.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this post model response based on the context it is used
func (m *PostModelResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Model
	if err := m.Model.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostModelResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostModelResponse) UnmarshalBinary(b []byte) error {
	var res PostModelResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}