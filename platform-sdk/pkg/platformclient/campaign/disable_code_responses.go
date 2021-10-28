// Code generated by go-swagger; DO NOT EDIT.

package campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/platform-sdk/pkg/platformclientmodels"
)

// DisableCodeReader is a Reader for the DisableCode structure.
type DisableCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDisableCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/codes/{code}/disable returns an error %d: %s", response.Code(), string(data))
	}
}

// NewDisableCodeOK creates a DisableCodeOK with default headers values
func NewDisableCodeOK() *DisableCodeOK {
	return &DisableCodeOK{}
}

/*DisableCodeOK handles this case with default header values.

  successful operation
*/
type DisableCodeOK struct {
	Payload *platformclientmodels.CodeInfo
}

func (o *DisableCodeOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/codes/{code}/disable][%d] disableCodeOK  %+v", 200, o.Payload)
}

func (o *DisableCodeOK) GetPayload() *platformclientmodels.CodeInfo {
	return o.Payload
}

func (o *DisableCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.CodeInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableCodeNotFound creates a DisableCodeNotFound with default headers values
func NewDisableCodeNotFound() *DisableCodeNotFound {
	return &DisableCodeNotFound{}
}

/*DisableCodeNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>37142</td><td>Code [{code}] does not exist in namespace [{namespace}]</td></tr><tr><td>37176</td><td>Code [{code}] has been redeemed in namespace [{namespace}]</td></tr></table>
*/
type DisableCodeNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *DisableCodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/codes/{code}/disable][%d] disableCodeNotFound  %+v", 404, o.Payload)
}

func (o *DisableCodeNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *DisableCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}