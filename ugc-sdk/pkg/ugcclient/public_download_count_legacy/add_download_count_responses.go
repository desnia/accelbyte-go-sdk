// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package public_download_count_legacy

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/ugc-sdk/pkg/ugcclientmodels"
)

// AddDownloadCountReader is a Reader for the AddDownloadCount structure.
type AddDownloadCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDownloadCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddDownloadCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddDownloadCountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddDownloadCountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewAddDownloadCountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewAddDownloadCountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /ugc/v1/public/namespaces/{namespace}/contents/{contentId}/downloadcount returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAddDownloadCountOK creates a AddDownloadCountOK with default headers values
func NewAddDownloadCountOK() *AddDownloadCountOK {
	return &AddDownloadCountOK{}
}

/*AddDownloadCountOK handles this case with default header values.

  OK
*/
type AddDownloadCountOK struct {
	Payload *ugcclientmodels.ModelsAddDownloadCountResponse
}

func (o *AddDownloadCountOK) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/public/namespaces/{namespace}/contents/{contentId}/downloadcount][%d] addDownloadCountOK  %+v", 200, o.ToJSONString())
}

func (o *AddDownloadCountOK) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *AddDownloadCountOK) GetPayload() *ugcclientmodels.ModelsAddDownloadCountResponse {
	return o.Payload
}

func (o *AddDownloadCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(ugcclientmodels.ModelsAddDownloadCountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDownloadCountBadRequest creates a AddDownloadCountBadRequest with default headers values
func NewAddDownloadCountBadRequest() *AddDownloadCountBadRequest {
	return &AddDownloadCountBadRequest{}
}

/*AddDownloadCountBadRequest handles this case with default header values.

  Bad Request
*/
type AddDownloadCountBadRequest struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AddDownloadCountBadRequest) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/public/namespaces/{namespace}/contents/{contentId}/downloadcount][%d] addDownloadCountBadRequest  %+v", 400, o.ToJSONString())
}

func (o *AddDownloadCountBadRequest) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *AddDownloadCountBadRequest) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AddDownloadCountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDownloadCountUnauthorized creates a AddDownloadCountUnauthorized with default headers values
func NewAddDownloadCountUnauthorized() *AddDownloadCountUnauthorized {
	return &AddDownloadCountUnauthorized{}
}

/*AddDownloadCountUnauthorized handles this case with default header values.

  Unauthorized
*/
type AddDownloadCountUnauthorized struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AddDownloadCountUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/public/namespaces/{namespace}/contents/{contentId}/downloadcount][%d] addDownloadCountUnauthorized  %+v", 401, o.ToJSONString())
}

func (o *AddDownloadCountUnauthorized) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *AddDownloadCountUnauthorized) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AddDownloadCountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDownloadCountNotFound creates a AddDownloadCountNotFound with default headers values
func NewAddDownloadCountNotFound() *AddDownloadCountNotFound {
	return &AddDownloadCountNotFound{}
}

/*AddDownloadCountNotFound handles this case with default header values.

  Not Found
*/
type AddDownloadCountNotFound struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AddDownloadCountNotFound) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/public/namespaces/{namespace}/contents/{contentId}/downloadcount][%d] addDownloadCountNotFound  %+v", 404, o.ToJSONString())
}

func (o *AddDownloadCountNotFound) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *AddDownloadCountNotFound) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AddDownloadCountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDownloadCountInternalServerError creates a AddDownloadCountInternalServerError with default headers values
func NewAddDownloadCountInternalServerError() *AddDownloadCountInternalServerError {
	return &AddDownloadCountInternalServerError{}
}

/*AddDownloadCountInternalServerError handles this case with default header values.

  Internal Server Error
*/
type AddDownloadCountInternalServerError struct {
	Payload *ugcclientmodels.ResponseError
}

func (o *AddDownloadCountInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ugc/v1/public/namespaces/{namespace}/contents/{contentId}/downloadcount][%d] addDownloadCountInternalServerError  %+v", 500, o.ToJSONString())
}

func (o *AddDownloadCountInternalServerError) ToJSONString() string {
	if o.Payload == nil {
		return "{}"
	}

	b, err := json.Marshal(o.Payload)
	if err != nil {
		fmt.Println(err)

		return fmt.Sprintf("Failed to marshal the payload: %+v", o.Payload)
	}

	return fmt.Sprintf("%+v", string(b))
}

func (o *AddDownloadCountInternalServerError) GetPayload() *ugcclientmodels.ResponseError {
	return o.Payload
}

func (o *AddDownloadCountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	o.Payload = new(ugcclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}