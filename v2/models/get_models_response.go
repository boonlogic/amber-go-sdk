// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetModelsResponse get models response
//
// swagger:model getModelsResponse
type GetModelsResponse struct {

	// model list
	ModelList []*Model `json:"modelList"`
}

// Validate validates this get models response
func (m *GetModelsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModelList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetModelsResponse) validateModelList(formats strfmt.Registry) error {
	if swag.IsZero(m.ModelList) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelList); i++ {
		if swag.IsZero(m.ModelList[i]) { // not required
			continue
		}

		if m.ModelList[i] != nil {
			if err := m.ModelList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modelList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modelList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get models response based on the context it is used
func (m *GetModelsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModelList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetModelsResponse) contextValidateModelList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModelList); i++ {

		if m.ModelList[i] != nil {
			if err := m.ModelList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modelList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modelList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetModelsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetModelsResponse) UnmarshalBinary(b []byte) error {
	var res GetModelsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
