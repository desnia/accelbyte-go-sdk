// Code generated by go-swagger; DO NOT EDIT.

package policy_versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/legal-sdk/pkg/legalclientmodels"
)

// RetrieveSinglePolicyVersionReader is a Reader for the RetrieveSinglePolicyVersion structure.
type RetrieveSinglePolicyVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveSinglePolicyVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrieveSinglePolicyVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRetrieveSinglePolicyVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/policies/{policyId}/versions returns an error %d: %s", response.Code(), string(data))
	}
}

// NewRetrieveSinglePolicyVersionOK creates a RetrieveSinglePolicyVersionOK with default headers values
func NewRetrieveSinglePolicyVersionOK() *RetrieveSinglePolicyVersionOK {
	return &RetrieveSinglePolicyVersionOK{}
}

/*RetrieveSinglePolicyVersionOK handles this case with default header values.

  successful operation
*/
type RetrieveSinglePolicyVersionOK struct {
	Payload []*legalclientmodels.RetrievePolicyVersionResponse
}

func (o *RetrieveSinglePolicyVersionOK) Error() string {
	return fmt.Sprintf("[GET /admin/policies/{policyId}/versions][%d] retrieveSinglePolicyVersionOK  %+v", 200, o.Payload)
}

func (o *RetrieveSinglePolicyVersionOK) GetPayload() []*legalclientmodels.RetrievePolicyVersionResponse {
	return o.Payload
}

func (o *RetrieveSinglePolicyVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveSinglePolicyVersionNotFound creates a RetrieveSinglePolicyVersionNotFound with default headers values
func NewRetrieveSinglePolicyVersionNotFound() *RetrieveSinglePolicyVersionNotFound {
	return &RetrieveSinglePolicyVersionNotFound{}
}

/*RetrieveSinglePolicyVersionNotFound handles this case with default header values.

  <table><tr><td>NumericErrorCode</td><td>ErrorCode</td></tr><tr><td>40035</td><td>errors.net.accelbyte.platform.legal.policy_version_not_found</td></tr></table>
*/
type RetrieveSinglePolicyVersionNotFound struct {
	Payload *legalclientmodels.ErrorEntity
}

func (o *RetrieveSinglePolicyVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/policies/{policyId}/versions][%d] retrieveSinglePolicyVersionNotFound  %+v", 404, o.Payload)
}

func (o *RetrieveSinglePolicyVersionNotFound) GetPayload() *legalclientmodels.ErrorEntity {
	return o.Payload
}

func (o *RetrieveSinglePolicyVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legalclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}