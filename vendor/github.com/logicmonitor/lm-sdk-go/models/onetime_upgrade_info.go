// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OnetimeUpgradeInfo onetime upgrade info
// swagger:model OnetimeUpgradeInfo
type OnetimeUpgradeInfo struct {

	// created by
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end epoch
	// Read Only: true
	EndEpoch int64 `json:"endEpoch,omitempty"`

	// level
	// Read Only: true
	Level string `json:"level,omitempty"`

	// major version
	// Required: true
	MajorVersion *int32 `json:"majorVersion"`

	// minor version
	// Required: true
	MinorVersion *int32 `json:"minorVersion"`

	// start epoch
	// Required: true
	StartEpoch *int64 `json:"startEpoch"`

	// timezone
	Timezone string `json:"timezone,omitempty"`

	// type
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this onetime upgrade info
func (m *OnetimeUpgradeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMajorVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinorVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartEpoch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnetimeUpgradeInfo) validateMajorVersion(formats strfmt.Registry) error {

	if err := validate.Required("majorVersion", "body", m.MajorVersion); err != nil {
		return err
	}

	return nil
}

func (m *OnetimeUpgradeInfo) validateMinorVersion(formats strfmt.Registry) error {

	if err := validate.Required("minorVersion", "body", m.MinorVersion); err != nil {
		return err
	}

	return nil
}

func (m *OnetimeUpgradeInfo) validateStartEpoch(formats strfmt.Registry) error {

	if err := validate.Required("startEpoch", "body", m.StartEpoch); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnetimeUpgradeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnetimeUpgradeInfo) UnmarshalBinary(b []byte) error {
	var res OnetimeUpgradeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}