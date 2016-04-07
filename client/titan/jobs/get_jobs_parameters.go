package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetJobsParams creates a new GetJobsParams object
// with the default values initialized.
func NewGetJobsParams() *GetJobsParams {
	var (
		nDefault int32 = int32(1)
	)
	return &GetJobsParams{
		N: &nDefault,
	}
}

/*GetJobsParams contains all the parameters to send to the API endpoint
for the get jobs operation typically these are written to a http.Request
*/
type GetJobsParams struct {

	/*N
	  Number of jobs to return.

	*/
	N *int32
}

// WithN adds the n to the get jobs params
func (o *GetJobsParams) WithN(n *int32) *GetJobsParams {
	o.N = n
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetJobsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.N != nil {

		// query param n
		var qrN int32
		if o.N != nil {
			qrN = *o.N
		}
		qN := swag.FormatInt32(qrN)
		if qN != "" {
			if err := r.SetQueryParam("n", qN); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
