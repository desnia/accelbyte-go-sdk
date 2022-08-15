// Code generated by go-swagger; DO NOT EDIT.

package dsmcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsCreateDeploymentRequest models create deployment request
//
// swagger:model models.CreateDeploymentRequest
type ModelsCreateDeploymentRequest struct {

	// allow version override
	// Required: true
	AllowVersionOverride *bool `json:"allow_version_override"`

	// buffer count
	// Required: true
	BufferCount *int32 `json:"buffer_count"`

	// buffer percent
	// Required: true
	BufferPercent *int32 `json:"buffer_percent"`

	// configuration
	// Required: true
	Configuration *string `json:"configuration"`

	// enable region overrides
	// Required: true
	EnableRegionOverrides *bool `json:"enable_region_overrides"`

	// game version
	// Required: true
	GameVersion *string `json:"game_version"`

	// max count
	// Required: true
	MaxCount *int32 `json:"max_count"`

	// min count
	// Required: true
	MinCount *int32 `json:"min_count"`

	// overrides
	// Required: true
	Overrides map[string]ModelsDeploymentConfigOverride `json:"overrides"`

	// region overrides
	// Required: true
	RegionOverrides map[string]ModelsPodCountConfigOverride `json:"region_overrides"`

	// regions
	// Required: true
	Regions []string `json:"regions"`

	// session timeout
	SessionTimeout int32 `json:"session_timeout,omitempty"`

	// use buffer percent
	// Required: true
	UseBufferPercent *bool `json:"use_buffer_percent"`
}

// Validate validates this models create deployment request
func (m *ModelsCreateDeploymentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowVersionOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBufferCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBufferPercent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnableRegionOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGameVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseBufferPercent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsCreateDeploymentRequest) validateAllowVersionOverride(formats strfmt.Registry) error {

	if err := validate.Required("allow_version_override", "body", m.AllowVersionOverride); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateBufferCount(formats strfmt.Registry) error {

	if err := validate.Required("buffer_count", "body", m.BufferCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateBufferPercent(formats strfmt.Registry) error {

	if err := validate.Required("buffer_percent", "body", m.BufferPercent); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateEnableRegionOverrides(formats strfmt.Registry) error {

	if err := validate.Required("enable_region_overrides", "body", m.EnableRegionOverrides); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateGameVersion(formats strfmt.Registry) error {

	if err := validate.Required("game_version", "body", m.GameVersion); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateMaxCount(formats strfmt.Registry) error {

	if err := validate.Required("max_count", "body", m.MaxCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateMinCount(formats strfmt.Registry) error {

	if err := validate.Required("min_count", "body", m.MinCount); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateOverrides(formats strfmt.Registry) error {

	for k := range m.Overrides {

		if err := validate.Required("overrides"+"."+k, "body", m.Overrides[k]); err != nil {
			return err
		}
		if val, ok := m.Overrides[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateRegionOverrides(formats strfmt.Registry) error {

	for k := range m.RegionOverrides {

		if err := validate.Required("region_overrides"+"."+k, "body", m.RegionOverrides[k]); err != nil {
			return err
		}
		if val, ok := m.RegionOverrides[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateRegions(formats strfmt.Registry) error {

	if err := validate.Required("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

func (m *ModelsCreateDeploymentRequest) validateUseBufferPercent(formats strfmt.Registry) error {

	if err := validate.Required("use_buffer_percent", "body", m.UseBufferPercent); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsCreateDeploymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsCreateDeploymentRequest) UnmarshalBinary(b []byte) error {
	var res ModelsCreateDeploymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
