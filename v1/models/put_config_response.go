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

// PutConfigResponse put config response
//
// swagger:model putConfigResponse
type PutConfigResponse struct {

	// features
	Features []*FusionConfig `json:"features"`

	// streaming
	Streaming *LearningParameters `json:"streaming,omitempty"`
}

// Validate validates this put config response
func (m *PutConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreaming(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutConfigResponse) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {
		if swag.IsZero(m.Features[i]) { // not required
			continue
		}

		if m.Features[i] != nil {
			if err := m.Features[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutConfigResponse) validateStreaming(formats strfmt.Registry) error {
	if swag.IsZero(m.Streaming) { // not required
		return nil
	}

	if m.Streaming != nil {
		if err := m.Streaming.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streaming")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streaming")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put config response based on the context it is used
func (m *PutConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreaming(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutConfigResponse) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Features); i++ {

		if m.Features[i] != nil {
			if err := m.Features[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("features" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("features" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PutConfigResponse) contextValidateStreaming(ctx context.Context, formats strfmt.Registry) error {

	if m.Streaming != nil {
		if err := m.Streaming.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("streaming")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("streaming")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutConfigResponse) UnmarshalBinary(b []byte) error {
	var res PutConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}