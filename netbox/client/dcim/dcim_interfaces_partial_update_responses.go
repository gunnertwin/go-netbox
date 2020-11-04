// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// DcimInterfacesPartialUpdateReader is a Reader for the DcimInterfacesPartialUpdate structure.
type DcimInterfacesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesPartialUpdateOK creates a DcimInterfacesPartialUpdateOK with default headers values
func NewDcimInterfacesPartialUpdateOK() *DcimInterfacesPartialUpdateOK {
	return &DcimInterfacesPartialUpdateOK{}
}

/*DcimInterfacesPartialUpdateOK handles this case with default header values.

DcimInterfacesPartialUpdateOK dcim interfaces partial update o k
*/
type DcimInterfacesPartialUpdateOK struct {
	Payload *models.Interface
}

func (o *DcimInterfacesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcimInterfacesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesPartialUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesPartialUpdateDefault creates a DcimInterfacesPartialUpdateDefault with default headers values
func NewDcimInterfacesPartialUpdateDefault(code int) *DcimInterfacesPartialUpdateDefault {
	return &DcimInterfacesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*DcimInterfacesPartialUpdateDefault handles this case with default header values.

DcimInterfacesPartialUpdateDefault dcim interfaces partial update default
*/
type DcimInterfacesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces partial update default response
func (o *DcimInterfacesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInterfacesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/interfaces/{id}/][%d] dcim_interfaces_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
