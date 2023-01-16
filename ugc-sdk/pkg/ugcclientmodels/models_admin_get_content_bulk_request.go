// Code generated by go-swagger; DO NOT EDIT.

package ugcclientmodels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ModelsAdminGetContentBulkRequest models admin get content bulk request
//
// swagger:model models.AdminGetContentBulkRequest
type ModelsAdminGetContentBulkRequest struct {

	// content ids
	// Required: true
	ContentIds []string `json:"contentIds"`
}

// Validate validates this models admin get content bulk request
func (m *ModelsAdminGetContentBulkRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ModelsAdminGetContentBulkRequest) validateContentIds(formats strfmt.Registry) error {

	if err := validate.Required("contentIds", "body", m.ContentIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ModelsAdminGetContentBulkRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelsAdminGetContentBulkRequest) UnmarshalBinary(b []byte) error {
	var res ModelsAdminGetContentBulkRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
