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

// ConfigResponse config response
//
// swagger:model configResponse
type ConfigResponse struct {

	// autotuning
	Autotuning *AutotuneConfig `json:"autotuning,omitempty"`

	// features
	Features []*FeatureConfigResponse `json:"features"`

	// percent variation
	PercentVariation *PercentVariation `json:"percentVariation,omitempty"`

	// streaming window
	StreamingWindow StreamingWindow `json:"streamingWindow,omitempty"`

	// training
	Training *TrainingConfig `json:"training,omitempty"`
}

// Validate validates this config response
func (m *ConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutotuning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentVariation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreamingWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraining(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigResponse) validateAutotuning(formats strfmt.Registry) error {
	if swag.IsZero(m.Autotuning) { // not required
		return nil
	}

	if m.Autotuning != nil {
		if err := m.Autotuning.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autotuning")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autotuning")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigResponse) validateFeatures(formats strfmt.Registry) error {
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

func (m *ConfigResponse) validatePercentVariation(formats strfmt.Registry) error {
	if swag.IsZero(m.PercentVariation) { // not required
		return nil
	}

	if m.PercentVariation != nil {
		if err := m.PercentVariation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("percentVariation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("percentVariation")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigResponse) validateStreamingWindow(formats strfmt.Registry) error {
	if swag.IsZero(m.StreamingWindow) { // not required
		return nil
	}

	if err := m.StreamingWindow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("streamingWindow")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("streamingWindow")
		}
		return err
	}

	return nil
}

func (m *ConfigResponse) validateTraining(formats strfmt.Registry) error {
	if swag.IsZero(m.Training) { // not required
		return nil
	}

	if m.Training != nil {
		if err := m.Training.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("training")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("training")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config response based on the context it is used
func (m *ConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutotuning(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePercentVariation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreamingWindow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTraining(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigResponse) contextValidateAutotuning(ctx context.Context, formats strfmt.Registry) error {

	if m.Autotuning != nil {
		if err := m.Autotuning.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autotuning")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autotuning")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigResponse) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConfigResponse) contextValidatePercentVariation(ctx context.Context, formats strfmt.Registry) error {

	if m.PercentVariation != nil {
		if err := m.PercentVariation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("percentVariation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("percentVariation")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigResponse) contextValidateStreamingWindow(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StreamingWindow.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("streamingWindow")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("streamingWindow")
		}
		return err
	}

	return nil
}

func (m *ConfigResponse) contextValidateTraining(ctx context.Context, formats strfmt.Registry) error {

	if m.Training != nil {
		if err := m.Training.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("training")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("training")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigResponse) UnmarshalBinary(b []byte) error {
	var res ConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}