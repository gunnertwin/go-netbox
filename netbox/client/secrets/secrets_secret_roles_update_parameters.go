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

package secrets

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

// NewSecretsSecretRolesUpdateParams creates a new SecretsSecretRolesUpdateParams object
// with the default values initialized.
func NewSecretsSecretRolesUpdateParams() *SecretsSecretRolesUpdateParams {
	var ()
	return &SecretsSecretRolesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSecretsSecretRolesUpdateParamsWithTimeout creates a new SecretsSecretRolesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSecretsSecretRolesUpdateParamsWithTimeout(timeout time.Duration) *SecretsSecretRolesUpdateParams {
	var ()
	return &SecretsSecretRolesUpdateParams{

		timeout: timeout,
	}
}

// NewSecretsSecretRolesUpdateParamsWithContext creates a new SecretsSecretRolesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSecretsSecretRolesUpdateParamsWithContext(ctx context.Context) *SecretsSecretRolesUpdateParams {
	var ()
	return &SecretsSecretRolesUpdateParams{

		Context: ctx,
	}
}

// NewSecretsSecretRolesUpdateParamsWithHTTPClient creates a new SecretsSecretRolesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSecretsSecretRolesUpdateParamsWithHTTPClient(client *http.Client) *SecretsSecretRolesUpdateParams {
	var ()
	return &SecretsSecretRolesUpdateParams{
		HTTPClient: client,
	}
}

/*SecretsSecretRolesUpdateParams contains all the parameters to send to the API endpoint
for the secrets secret roles update operation typically these are written to a http.Request
*/
type SecretsSecretRolesUpdateParams struct {

	/*Data*/
	Data *models.SecretRole
	/*ID
	  A unique integer value identifying this secret role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) WithTimeout(timeout time.Duration) *SecretsSecretRolesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) WithContext(ctx context.Context) *SecretsSecretRolesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) WithHTTPClient(client *http.Client) *SecretsSecretRolesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) WithData(data *models.SecretRole) *SecretsSecretRolesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) SetData(data *models.SecretRole) {
	o.Data = data
}

// WithID adds the id to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) WithID(id int64) *SecretsSecretRolesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the secrets secret roles update params
func (o *SecretsSecretRolesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SecretsSecretRolesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
