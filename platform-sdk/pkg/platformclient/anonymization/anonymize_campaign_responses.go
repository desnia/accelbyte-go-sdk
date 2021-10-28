// Code generated by go-swagger; DO NOT EDIT.

package anonymization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AnonymizeCampaignReader is a Reader for the AnonymizeCampaign structure.
type AnonymizeCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AnonymizeCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAnonymizeCampaignNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /admin/namespaces/{namespace}/users/{userId}/anonymization/campaign returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAnonymizeCampaignNoContent creates a AnonymizeCampaignNoContent with default headers values
func NewAnonymizeCampaignNoContent() *AnonymizeCampaignNoContent {
	return &AnonymizeCampaignNoContent{}
}

/*AnonymizeCampaignNoContent handles this case with default header values.

  Anonymize successfully
*/
type AnonymizeCampaignNoContent struct {
}

func (o *AnonymizeCampaignNoContent) Error() string {
	return fmt.Sprintf("[DELETE /admin/namespaces/{namespace}/users/{userId}/anonymization/campaign][%d] anonymizeCampaignNoContent ", 204)
}

func (o *AnonymizeCampaignNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}