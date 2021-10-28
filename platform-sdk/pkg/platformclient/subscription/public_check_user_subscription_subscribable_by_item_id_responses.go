// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// PublicCheckUserSubscriptionSubscribableByItemIDReader is a Reader for the PublicCheckUserSubscriptionSubscribableByItemID structure.
type PublicCheckUserSubscriptionSubscribableByItemIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicCheckUserSubscriptionSubscribableByItemIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicCheckUserSubscriptionSubscribableByItemIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /public/namespaces/{namespace}/users/{userId}/subscriptions/subscribable/byItemId returns an error %d: %s", response.Code(), string(data))
	}
}

// NewPublicCheckUserSubscriptionSubscribableByItemIDOK creates a PublicCheckUserSubscriptionSubscribableByItemIDOK with default headers values
func NewPublicCheckUserSubscriptionSubscribableByItemIDOK() *PublicCheckUserSubscriptionSubscribableByItemIDOK {
	return &PublicCheckUserSubscriptionSubscribableByItemIDOK{}
}

/*PublicCheckUserSubscriptionSubscribableByItemIDOK handles this case with default header values.

  successful operation
*/
type PublicCheckUserSubscriptionSubscribableByItemIDOK struct {
	Payload *platformclientmodels.Subscribable
}

func (o *PublicCheckUserSubscriptionSubscribableByItemIDOK) Error() string {
	return fmt.Sprintf("[GET /public/namespaces/{namespace}/users/{userId}/subscriptions/subscribable/byItemId][%d] publicCheckUserSubscriptionSubscribableByItemIdOK  %+v", 200, o.Payload)
}

func (o *PublicCheckUserSubscriptionSubscribableByItemIDOK) GetPayload() *platformclientmodels.Subscribable {
	return o.Payload
}

func (o *PublicCheckUserSubscriptionSubscribableByItemIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.Subscribable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}