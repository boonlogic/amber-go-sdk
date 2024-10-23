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

// GetStatusResponse get status response
//
// swagger:model getStatusResponse
type GetStatusResponse struct {
	ModelStatus
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetStatusResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ModelStatus
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ModelStatus = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetStatusResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ModelStatus)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get status response
func (m *GetStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelStatus
	if err := m.ModelStatus.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get status response based on the context it is used
func (m *GetStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ModelStatus
	if err := m.ModelStatus.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetStatusResponse) UnmarshalBinary(b []byte) error {
	var res GetStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
