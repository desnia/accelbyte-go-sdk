// Code generated by go-swagger; DO NOT EDIT.

package item

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

// PublicGetItemReader is a Reader for the PublicGetItem structure.
type PublicGetItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicGetItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicGetItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPublicGetItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/items/{itemId}/locale returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicGetItemOK creates a PublicGetItemOK with default headers values
func NewPublicGetItemOK() *PublicGetItemOK {
	return &PublicGetItemOK{}
}

/*PublicGetItemOK handles this case with default header values.

  successful operation
*/
type PublicGetItemOK struct {
	Payload *platformclientmodels.PopulatedItemInfo
}

func (o *PublicGetItemOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/items/{itemId}/locale][%d] publicGetItemOK  %+v", 200, o.Payload)
}

func (o *PublicGetItemOK) GetPayload() *platformclientmodels.PopulatedItemInfo {
	return o.Payload
}

func (o *PublicGetItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.PopulatedItemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicGetItemNotFound creates a PublicGetItemNotFound with default headers values
func NewPublicGetItemNotFound() *PublicGetItemNotFound {
	return &PublicGetItemNotFound{}
}

/*PublicGetItemNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>30341</td><td>Item [{itemId}] does not exist in namespace [{namespace}]</td></tr></table>
*/
type PublicGetItemNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *PublicGetItemNotFound) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/items/{itemId}/locale][%d] publicGetItemNotFound  %+v", 404, o.Payload)
}

func (o *PublicGetItemNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *PublicGetItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}