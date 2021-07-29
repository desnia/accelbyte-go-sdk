// Code generated by go-swagger; DO NOT EDIT.

package payment_config

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

// GetPaymentMerchantConfigReader is a Reader for the GetPaymentMerchantConfig structure.
type GetPaymentMerchantConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMerchantConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentMerchantConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPaymentMerchantConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /admin/payment/config/merchant/{id} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetPaymentMerchantConfigOK creates a GetPaymentMerchantConfigOK with default headers values
func NewGetPaymentMerchantConfigOK() *GetPaymentMerchantConfigOK {
	return &GetPaymentMerchantConfigOK{}
}

/*GetPaymentMerchantConfigOK handles this case with default header values.

  successful operation
*/
type GetPaymentMerchantConfigOK struct {
	Payload *platformclientmodels.PaymentMerchantConfigInfo
}

func (o *GetPaymentMerchantConfigOK) Error() string {
	return fmt.Sprintf("[GET /admin/payment/config/merchant/{id}][%d] getPaymentMerchantConfigOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMerchantConfigOK) GetPayload() *platformclientmodels.PaymentMerchantConfigInfo {
	return o.Payload
}

func (o *GetPaymentMerchantConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.PaymentMerchantConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMerchantConfigNotFound creates a GetPaymentMerchantConfigNotFound with default headers values
func NewGetPaymentMerchantConfigNotFound() *GetPaymentMerchantConfigNotFound {
	return &GetPaymentMerchantConfigNotFound{}
}

/*GetPaymentMerchantConfigNotFound handles this case with default header values.

  <table><tr><td>ErrorCode</td><td>ErrorMessage</td></tr><tr><td>33242</td><td>Payment merchant config [{id}] does not exist</td></tr></table>
*/
type GetPaymentMerchantConfigNotFound struct {
	Payload *platformclientmodels.ErrorEntity
}

func (o *GetPaymentMerchantConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /admin/payment/config/merchant/{id}][%d] getPaymentMerchantConfigNotFound  %+v", 404, o.Payload)
}

func (o *GetPaymentMerchantConfigNotFound) GetPayload() *platformclientmodels.ErrorEntity {
	return o.Payload
}

func (o *GetPaymentMerchantConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}