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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// IpamServicesUpdateReader is a Reader for the IpamServicesUpdate structure.
type IpamServicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamServicesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesUpdateOK creates a IpamServicesUpdateOK with default headers values
func NewIpamServicesUpdateOK() *IpamServicesUpdateOK {
	return &IpamServicesUpdateOK{}
}

/*IpamServicesUpdateOK handles this case with default header values.

IpamServicesUpdateOK ipam services update o k
*/
type IpamServicesUpdateOK struct {
	Payload *models.Service
}

func (o *IpamServicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipamServicesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesUpdateDefault creates a IpamServicesUpdateDefault with default headers values
func NewIpamServicesUpdateDefault(code int) *IpamServicesUpdateDefault {
	return &IpamServicesUpdateDefault{
		_statusCode: code,
	}
}

/*IpamServicesUpdateDefault handles this case with default header values.

IpamServicesUpdateDefault ipam services update default
*/
type IpamServicesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam services update default response
func (o *IpamServicesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamServicesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/services/{id}/][%d] ipam_services_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
