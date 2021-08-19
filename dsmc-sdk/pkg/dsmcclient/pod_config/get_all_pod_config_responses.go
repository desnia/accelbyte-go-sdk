// Code generated by go-swagger; DO NOT EDIT.

package pod_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dsmc-sdk/pkg/dsmcclientmodels"
)

// GetAllPodConfigReader is a Reader for the GetAllPodConfig structure.
type GetAllPodConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllPodConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllPodConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAllPodConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllPodConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAllPodConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /dsmcontroller/admin/namespaces/{namespace}/configs/pods returns an error %d: %s", response.Code(), string(data))
	}
}

// NewGetAllPodConfigOK creates a GetAllPodConfigOK with default headers values
func NewGetAllPodConfigOK() *GetAllPodConfigOK {
	return &GetAllPodConfigOK{}
}

/*GetAllPodConfigOK handles this case with default header values.

  ok
*/
type GetAllPodConfigOK struct {
	Payload *dsmcclientmodels.ModelsListPodConfigResponse
}

func (o *GetAllPodConfigOK) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/configs/pods][%d] getAllPodConfigOK  %+v", 200, o.Payload)
}

func (o *GetAllPodConfigOK) GetPayload() *dsmcclientmodels.ModelsListPodConfigResponse {
	return o.Payload
}

func (o *GetAllPodConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ModelsListPodConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPodConfigBadRequest creates a GetAllPodConfigBadRequest with default headers values
func NewGetAllPodConfigBadRequest() *GetAllPodConfigBadRequest {
	return &GetAllPodConfigBadRequest{}
}

/*GetAllPodConfigBadRequest handles this case with default header values.

  malformed request
*/
type GetAllPodConfigBadRequest struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetAllPodConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/configs/pods][%d] getAllPodConfigBadRequest  %+v", 400, o.Payload)
}

func (o *GetAllPodConfigBadRequest) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetAllPodConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPodConfigUnauthorized creates a GetAllPodConfigUnauthorized with default headers values
func NewGetAllPodConfigUnauthorized() *GetAllPodConfigUnauthorized {
	return &GetAllPodConfigUnauthorized{}
}

/*GetAllPodConfigUnauthorized handles this case with default header values.

  Unauthorized
*/
type GetAllPodConfigUnauthorized struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetAllPodConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/configs/pods][%d] getAllPodConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllPodConfigUnauthorized) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetAllPodConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPodConfigInternalServerError creates a GetAllPodConfigInternalServerError with default headers values
func NewGetAllPodConfigInternalServerError() *GetAllPodConfigInternalServerError {
	return &GetAllPodConfigInternalServerError{}
}

/*GetAllPodConfigInternalServerError handles this case with default header values.

  Internal Server Error
*/
type GetAllPodConfigInternalServerError struct {
	Payload *dsmcclientmodels.ResponseError
}

func (o *GetAllPodConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dsmcontroller/admin/namespaces/{namespace}/configs/pods][%d] getAllPodConfigInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllPodConfigInternalServerError) GetPayload() *dsmcclientmodels.ResponseError {
	return o.Payload
}

func (o *GetAllPodConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dsmcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}