// Code generated by go-swagger; DO NOT EDIT.

package nr_admin_content

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// AdminDeleteContentReader is a Reader for the AdminDeleteContent structure.
type AdminDeleteContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminDeleteContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminDeleteContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAdminDeleteContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAdminDeleteContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAdminDeleteContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested DELETE /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/{contentId} returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminDeleteContentOK creates a AdminDeleteContentOK with default headers values
func NewAdminDeleteContentOK() *AdminDeleteContentOK {
	return &AdminDeleteContentOK{}
}

/*AdminDeleteContentOK handles this case with default header values.

  OK
*/
type AdminDeleteContentOK struct {
	Payload *ugcclientmodels.ModelsCreateContentResponse
}

func (o *AdminDeleteContentOK) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/{contentId}][%d] adminDeleteContentOK  %+v", 200, o.Payload)
}

func (o *AdminDeleteContentOK) GetPayload() *ugcclientmodels.ModelsCreateContentResponse {
	return o.Payload
}

func (o *AdminDeleteContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ModelsCreateContentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteContentUnauthorized creates a AdminDeleteContentUnauthorized with default headers values
func NewAdminDeleteContentUnauthorized() *AdminDeleteContentUnauthorized {
	return &AdminDeleteContentUnauthorized{}
}

/*AdminDeleteContentUnauthorized handles this case with default header values.

  Unauthorized
*/
type AdminDeleteContentUnauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminDeleteContentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/{contentId}][%d] adminDeleteContentUnauthorized  %+v", 401, o.Payload)
}

func (o *AdminDeleteContentUnauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminDeleteContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteContentNotFound creates a AdminDeleteContentNotFound with default headers values
func NewAdminDeleteContentNotFound() *AdminDeleteContentNotFound {
	return &AdminDeleteContentNotFound{}
}

/*AdminDeleteContentNotFound handles this case with default header values.

  Not Found
*/
type AdminDeleteContentNotFound struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminDeleteContentNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/{contentId}][%d] adminDeleteContentNotFound  %+v", 404, o.Payload)
}

func (o *AdminDeleteContentNotFound) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminDeleteContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminDeleteContentInternalServerError creates a AdminDeleteContentInternalServerError with default headers values
func NewAdminDeleteContentInternalServerError() *AdminDeleteContentInternalServerError {
	return &AdminDeleteContentInternalServerError{}
}

/*AdminDeleteContentInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AdminDeleteContentInternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AdminDeleteContentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ugc/v1/admin/namespaces/{namespace}/users/{userId}/channels/{channelId}/contents/{contentId}][%d] adminDeleteContentInternalServerError  %+v", 500, o.Payload)
}

func (o *AdminDeleteContentInternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AdminDeleteContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}