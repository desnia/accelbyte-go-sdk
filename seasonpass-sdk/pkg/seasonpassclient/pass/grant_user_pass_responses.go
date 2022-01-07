// Code generated by go-swagger; DO NOT EDIT.

package pass

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/seasonpass-sdk/pkg/seasonpassclientmodels"
)

// GrantUserPassReader is a Reader for the GrantUserPass structure.
type GrantUserPassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GrantUserPassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGrantUserPassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGrantUserPassBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /admin/namespaces/{namespace}/users/{userId}/seasons/current/passes returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGrantUserPassOK creates a GrantUserPassOK with default headers values
func NewGrantUserPassOK() *GrantUserPassOK {
	return &GrantUserPassOK{}
}

/*GrantUserPassOK handles this case with default header values.

  successful operation
*/
type GrantUserPassOK struct {
	Payload *seasonpassclientmodels.UserSeasonSummary
}

func (o *GrantUserPassOK) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/seasons/current/passes][%d] grantUserPassOK  %+v", 200, o.Payload)
}

func (o *GrantUserPassOK) GetPayload() *seasonpassclientmodels.UserSeasonSummary {
	return o.Payload
}

func (o *GrantUserPassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(seasonpassclientmodels.UserSeasonSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGrantUserPassBadRequest creates a GrantUserPassBadRequest with default headers values
func NewGrantUserPassBadRequest() *GrantUserPassBadRequest {
	return &GrantUserPassBadRequest{}
}

/*GrantUserPassBadRequest handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>20026</td><td>publisher namespace not allowed</td></tr></table>
*/
type GrantUserPassBadRequest struct {
	Payload *seasonpassclientmodels.ErrorEntity
}

func (o *GrantUserPassBadRequest) Error() string {
	return fmt.Sprintf("[POST /admin/namespaces/{namespace}/users/{userId}/seasons/current/passes][%d] grantUserPassBadRequest  %+v", 400, o.Payload)
}

func (o *GrantUserPassBadRequest) GetPayload() *seasonpassclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GrantUserPassBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(seasonpassclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}