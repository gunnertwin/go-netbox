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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/gunnertwin/go-netbox/models"
)

// NewDcimConsolePortTemplatesPartialUpdateParams creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the default values initialized.
func NewDcimConsolePortTemplatesPartialUpdateParams() *DcimConsolePortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsolePortTemplatesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesPartialUpdateParamsWithTimeout creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsolePortTemplatesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesPartialUpdateParamsWithContext creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsolePortTemplatesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsolePortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim console port templates partial update operation typically these are written to a http.Request
*/
type DcimConsolePortTemplatesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableConsolePortTemplate
	/*ID
	  A unique integer value identifying this console port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithData(data *models.WritableConsolePortTemplate) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetData(data *models.WritableConsolePortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithID(id int64) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
