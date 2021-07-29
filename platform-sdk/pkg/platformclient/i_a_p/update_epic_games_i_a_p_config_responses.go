// Code generated by go-swagger; DO NOT EDIT.

package i_a_p

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

// UpdateEpicGamesIAPConfigReader is a Reader for the UpdateEpicGamesIAPConfig structure.
type UpdateEpicGamesIAPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEpicGamesIAPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEpicGamesIAPConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested PUT /admin/namespaces/{namespace}/iap/config/epicgames returns an error %d: %s", response.Code(), string(data))
	}
}

// NewUpdateEpicGamesIAPConfigOK creates a UpdateEpicGamesIAPConfigOK with default headers values
func NewUpdateEpicGamesIAPConfigOK() *UpdateEpicGamesIAPConfigOK {
	return &UpdateEpicGamesIAPConfigOK{}
}

/*UpdateEpicGamesIAPConfigOK handles this case with default header values.

  successful operation
*/
type UpdateEpicGamesIAPConfigOK struct {
	Payload *platformclientmodels.EpicGamesIAPConfigInfo
}

func (o *UpdateEpicGamesIAPConfigOK) Error() string {
	return fmt.Sprintf("[PUT /admin/namespaces/{namespace}/iap/config/epicgames][%d] updateEpicGamesIAPConfigOK  %+v", 200, o.Payload)
}

func (o *UpdateEpicGamesIAPConfigOK) GetPayload() *platformclientmodels.EpicGamesIAPConfigInfo {
	return o.Payload
}

func (o *UpdateEpicGamesIAPConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(platformclientmodels.EpicGamesIAPConfigInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}