// Copyright (c) 2021 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

// Code generated; DO NOT EDIT.

package operations

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SymbolHandlerReader is a Reader for the SymbolHandler structure.
type SymbolHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SymbolHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSymbolHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		data, err := ioutil.ReadAll(response.Body())
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("Requested GET /session/admin/internal/pprof/symbol returns an error %d: %s", response.Code(), string(data))
	}
}

// NewSymbolHandlerOK creates a SymbolHandlerOK with default headers values
func NewSymbolHandlerOK() *SymbolHandlerOK {
	return &SymbolHandlerOK{}
}

/*SymbolHandlerOK handles this case with default header values.

  OK
*/
type SymbolHandlerOK struct {
}

func (o *SymbolHandlerOK) Error() string {
	return fmt.Sprintf("[GET /session/admin/internal/pprof/symbol][%d] symbolHandlerOK ", 200)
}

func (o *SymbolHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// handle file responses
	contentDisposition := response.GetHeader("Content-Disposition")
	if strings.Contains(strings.ToLower(contentDisposition), "filename=") {
		consumer = runtime.ByteStreamConsumer()
	}

	return nil
}
