// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated by go-swagger; DO NOT EDIT.

package gametelemetry_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetReader is a Reader for the AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGet structure.
type AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /game-telemetry/v1/admin/telemetrynamespace returns an error %d: %s", response.Code(), string(data))
	}
}

// NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK creates a AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK with default headers values
func NewAdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK() *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK {
	return &AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK{}
}

/*AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK handles this case with default header values.

  Successful Response
*/
type AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK struct {
	Payload interface{}
}

func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK) Error() string {
	return fmt.Sprintf("[GET /game-telemetry/v1/admin/telemetrynamespace][%d] adminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK  %+v", 200, o.Payload)
}

func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AdminGetNamespaceGameTelemetryV1AdminTelemetrynamespaceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}