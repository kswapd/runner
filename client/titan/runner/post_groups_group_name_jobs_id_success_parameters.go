package runner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewPostGroupsGroupNameJobsIDSuccessParams creates a new PostGroupsGroupNameJobsIDSuccessParams object
// with the default values initialized.
func NewPostGroupsGroupNameJobsIDSuccessParams() *PostGroupsGroupNameJobsIDSuccessParams {
	var ()
	return &PostGroupsGroupNameJobsIDSuccessParams{}
}

/*PostGroupsGroupNameJobsIDSuccessParams contains all the parameters to send to the API endpoint
for the post groups group name jobs ID success operation typically these are written to a http.Request
*/
type PostGroupsGroupNameJobsIDSuccessParams struct {

	/*GroupName
	  Name of group for this set of jobs.

	*/
	GroupName string
	/*ID
	  Job id

	*/
	ID string
}

// WithGroupName adds the groupName to the post groups group name jobs ID success params
func (o *PostGroupsGroupNameJobsIDSuccessParams) WithGroupName(groupName string) *PostGroupsGroupNameJobsIDSuccessParams {
	o.GroupName = groupName
	return o
}

// WithID adds the id to the post groups group name jobs ID success params
func (o *PostGroupsGroupNameJobsIDSuccessParams) WithID(id string) *PostGroupsGroupNameJobsIDSuccessParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostGroupsGroupNameJobsIDSuccessParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param group_name
	if err := r.SetPathParam("group_name", o.GroupName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
