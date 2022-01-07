// Code generated by go-swagger; DO NOT EDIT.

package all_terminated_servers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AccelByte/accelbyte-go-sdk/dslogmanager-sdk/pkg/dslogmanagerclientmodels"
)

// BatchDownloadServerLogsReader is a Reader for the BatchDownloadServerLogs structure.
type BatchDownloadServerLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchDownloadServerLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBatchDownloadServerLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewBatchDownloadServerLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewBatchDownloadServerLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested POST /dslogmanager/servers/logs/download returns an error %d: %s", response.Code(), string(data))
	}
}

// NewBatchDownloadServerLogsOK creates a BatchDownloadServerLogsOK with default headers values
func NewBatchDownloadServerLogsOK() *BatchDownloadServerLogsOK {
	return &BatchDownloadServerLogsOK{}
}

/*BatchDownloadServerLogsOK handles this case with default header values.

  server logs downloaded.
*/
type BatchDownloadServerLogsOK struct {
}

func (o *BatchDownloadServerLogsOK) Error() string {
	return fmt.Sprintf("[POST /dslogmanager/servers/logs/download][%d] batchDownloadServerLogsOK ", 200)
}

func (o *BatchDownloadServerLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBatchDownloadServerLogsBadRequest creates a BatchDownloadServerLogsBadRequest with default headers values
func NewBatchDownloadServerLogsBadRequest() *BatchDownloadServerLogsBadRequest {
	return &BatchDownloadServerLogsBadRequest{}
}

/*BatchDownloadServerLogsBadRequest handles this case with default header values.

  Bad Request
*/
type BatchDownloadServerLogsBadRequest struct {
	Payload *dslogmanagerclientmodels.ResponseError
}

func (o *BatchDownloadServerLogsBadRequest) Error() string {
	return fmt.Sprintf("[POST /dslogmanager/servers/logs/download][%d] batchDownloadServerLogsBadRequest  %+v", 400, o.Payload)
}

func (o *BatchDownloadServerLogsBadRequest) GetPayload() *dslogmanagerclientmodels.ResponseError {
	return o.Payload
}

func (o *BatchDownloadServerLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dslogmanagerclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchDownloadServerLogsInternalServerError creates a BatchDownloadServerLogsInternalServerError with default headers values
func NewBatchDownloadServerLogsInternalServerError() *BatchDownloadServerLogsInternalServerError {
	return &BatchDownloadServerLogsInternalServerError{}
}

/*BatchDownloadServerLogsInternalServerError handles this case with default header values.

  Internal Server Error
*/
type BatchDownloadServerLogsInternalServerError struct {
	Payload *dslogmanagerclientmodels.ResponseError
}

func (o *BatchDownloadServerLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dslogmanager/servers/logs/download][%d] batchDownloadServerLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *BatchDownloadServerLogsInternalServerError) GetPayload() *dslogmanagerclientmodels.ResponseError {
	return o.Payload
}

func (o *BatchDownloadServerLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(dslogmanagerclientmodels.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}