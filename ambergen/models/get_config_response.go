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
	"github.com/go-openapi/validate"
)

// GetConfigResponse get config response
//
// swagger:model getConfigResponse
type GetConfigResponse struct {

	// the number of samples to use when calculating AH
	// Required: true
	AnomalyHistoryWindow *uint32 `json:"anomalyHistoryWindow"`

	// number of features per sample
	// Required: true
	FeatureCount *uint16 `json:"featureCount"`

	// features
	// Required: true
	Features []*FeatureConfig `json:"features"`

	// learning graduation requirement for stopping learning upon reaching this cluster count
	// Required: true
	LearningMaxClusters *uint16 `json:"learningMaxClusters"`

	// learning graduation requirement for stopping learning after acquiring this many samples
	// Required: true
	LearningMaxSamples *uint64 `json:"learningMaxSamples"`

	// enables graduation requirements for learning
	// Required: true
	LearningRateDenominator *uint64 `json:"learningRateDenominator"`

	// enables graduation requirements for learning
	// Required: true
	LearningRateNumerator *uint64 `json:"learningRateNumerator"`

	// the percent variation (for instance, 0.025 gives 2.5% variation) used for clustering
	// Required: true
	PercentVariation *float32 `json:"percentVariation"`

	// the number of samples to be applied before autotuning begins
	// Required: true
	SamplesToBuffer *uint32 `json:"samplesToBuffer"`

	// streaming window size
	// Required: true
	// Maximum: 500
	StreamingWindowSize *uint16 `json:"streamingWindowSize"`
}

// Validate validates this get config response
func (m *GetConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnomalyHistoryWindow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatureCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLearningMaxClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLearningMaxSamples(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLearningRateDenominator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLearningRateNumerator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentVariation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSamplesToBuffer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreamingWindowSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConfigResponse) validateAnomalyHistoryWindow(formats strfmt.Registry) error {

	if err := validate.Required("anomalyHistoryWindow", "body", m.AnomalyHistoryWindow); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateFeatureCount(formats strfmt.Registry) error {

	if err := validate.Required("featureCount", "body", m.FeatureCount); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
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

func (m *GetConfigResponse) validateLearningMaxClusters(formats strfmt.Registry) error {

	if err := validate.Required("learningMaxClusters", "body", m.LearningMaxClusters); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateLearningMaxSamples(formats strfmt.Registry) error {

	if err := validate.Required("learningMaxSamples", "body", m.LearningMaxSamples); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateLearningRateDenominator(formats strfmt.Registry) error {

	if err := validate.Required("learningRateDenominator", "body", m.LearningRateDenominator); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateLearningRateNumerator(formats strfmt.Registry) error {

	if err := validate.Required("learningRateNumerator", "body", m.LearningRateNumerator); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validatePercentVariation(formats strfmt.Registry) error {

	if err := validate.Required("percentVariation", "body", m.PercentVariation); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateSamplesToBuffer(formats strfmt.Registry) error {

	if err := validate.Required("samplesToBuffer", "body", m.SamplesToBuffer); err != nil {
		return err
	}

	return nil
}

func (m *GetConfigResponse) validateStreamingWindowSize(formats strfmt.Registry) error {

	if err := validate.Required("streamingWindowSize", "body", m.StreamingWindowSize); err != nil {
		return err
	}

	if err := validate.MaximumUint("streamingWindowSize", "body", uint64(*m.StreamingWindowSize), 500, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get config response based on the context it is used
func (m *GetConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetConfigResponse) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *GetConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetConfigResponse) UnmarshalBinary(b []byte) error {
	var res GetConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
