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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/gunnertwin/go-netbox/models"
)

// VirtualizationClusterGroupsPartialUpdateReader is a Reader for the VirtualizationClusterGroupsPartialUpdate structure.
type VirtualizationClusterGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationClusterGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClusterGroupsPartialUpdateOK creates a VirtualizationClusterGroupsPartialUpdateOK with default headers values
func NewVirtualizationClusterGroupsPartialUpdateOK() *VirtualizationClusterGroupsPartialUpdateOK {
	return &VirtualizationClusterGroupsPartialUpdateOK{}
}

/*VirtualizationClusterGroupsPartialUpdateOK handles this case with default header values.

VirtualizationClusterGroupsPartialUpdateOK virtualization cluster groups partial update o k
*/
type VirtualizationClusterGroupsPartialUpdateOK struct {
	Payload *models.ClusterGroup
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsPartialUpdateDefault creates a VirtualizationClusterGroupsPartialUpdateDefault with default headers values
func NewVirtualizationClusterGroupsPartialUpdateDefault(code int) *VirtualizationClusterGroupsPartialUpdateDefault {
	return &VirtualizationClusterGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*VirtualizationClusterGroupsPartialUpdateDefault handles this case with default header values.

VirtualizationClusterGroupsPartialUpdateDefault virtualization cluster groups partial update default
*/
type VirtualizationClusterGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization cluster groups partial update default response
func (o *VirtualizationClusterGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/cluster-groups/{id}/][%d] virtualization_cluster-groups_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
