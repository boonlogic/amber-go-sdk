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

// GetModelDataSetRunResponse get model data set run response
//
// swagger:model getModelDataSetRunResponse
type GetModelDataSetRunResponse struct {
	DataSetRunResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetModelDataSetRunResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DataSetRunResponse
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DataSetRunResponse = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetModelDataSetRunResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.DataSetRunResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get model data set run response
func (m *GetModelDataSetRunResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DataSetRunResponse
	if err := m.DataSetRunResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this get model data set run response based on the context it is used
func (m *GetModelDataSetRunResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DataSetRunResponse
	if err := m.DataSetRunResponse.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetModelDataSetRunResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetModelDataSetRunResponse) UnmarshalBinary(b []byte) error {
	var res GetModelDataSetRunResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
