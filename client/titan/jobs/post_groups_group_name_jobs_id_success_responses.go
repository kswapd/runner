package jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/iron-io/titan/runner/client/models"
)

// PostGroupsGroupNameJobsIDSuccessReader is a Reader for the PostGroupsGroupNameJobsIDSuccess structure.
type PostGroupsGroupNameJobsIDSuccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostGroupsGroupNameJobsIDSuccessReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostGroupsGroupNameJobsIDSuccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostGroupsGroupNameJobsIDSuccessNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostGroupsGroupNameJobsIDSuccessConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostGroupsGroupNameJobsIDSuccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostGroupsGroupNameJobsIDSuccessOK creates a PostGroupsGroupNameJobsIDSuccessOK with default headers values
func NewPostGroupsGroupNameJobsIDSuccessOK() *PostGroupsGroupNameJobsIDSuccessOK {
	return &PostGroupsGroupNameJobsIDSuccessOK{}
}

/*PostGroupsGroupNameJobsIDSuccessOK handles this case with default header values.

Job updated.
*/
type PostGroupsGroupNameJobsIDSuccessOK struct {
	Payload *models.JobWrapper
}

func (o *PostGroupsGroupNameJobsIDSuccessOK) Error() string {
	return fmt.Sprintf("[POST /groups/{group_name}/jobs/{id}/success][%d] postGroupsGroupNameJobsIdSuccessOK  %+v", 200, o.Payload)
}

func (o *PostGroupsGroupNameJobsIDSuccessOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsGroupNameJobsIDSuccessNotFound creates a PostGroupsGroupNameJobsIDSuccessNotFound with default headers values
func NewPostGroupsGroupNameJobsIDSuccessNotFound() *PostGroupsGroupNameJobsIDSuccessNotFound {
	return &PostGroupsGroupNameJobsIDSuccessNotFound{}
}

/*PostGroupsGroupNameJobsIDSuccessNotFound handles this case with default header values.

Job does not exist.
*/
type PostGroupsGroupNameJobsIDSuccessNotFound struct {
	Payload *models.Error
}

func (o *PostGroupsGroupNameJobsIDSuccessNotFound) Error() string {
	return fmt.Sprintf("[POST /groups/{group_name}/jobs/{id}/success][%d] postGroupsGroupNameJobsIdSuccessNotFound  %+v", 404, o.Payload)
}

func (o *PostGroupsGroupNameJobsIDSuccessNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsGroupNameJobsIDSuccessConflict creates a PostGroupsGroupNameJobsIDSuccessConflict with default headers values
func NewPostGroupsGroupNameJobsIDSuccessConflict() *PostGroupsGroupNameJobsIDSuccessConflict {
	return &PostGroupsGroupNameJobsIDSuccessConflict{}
}

/*PostGroupsGroupNameJobsIDSuccessConflict handles this case with default header values.

Job was not in running state.
*/
type PostGroupsGroupNameJobsIDSuccessConflict struct {
	Payload *models.IDStatus
}

func (o *PostGroupsGroupNameJobsIDSuccessConflict) Error() string {
	return fmt.Sprintf("[POST /groups/{group_name}/jobs/{id}/success][%d] postGroupsGroupNameJobsIdSuccessConflict  %+v", 409, o.Payload)
}

func (o *PostGroupsGroupNameJobsIDSuccessConflict) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGroupsGroupNameJobsIDSuccessDefault creates a PostGroupsGroupNameJobsIDSuccessDefault with default headers values
func NewPostGroupsGroupNameJobsIDSuccessDefault(code int) *PostGroupsGroupNameJobsIDSuccessDefault {
	return &PostGroupsGroupNameJobsIDSuccessDefault{
		_statusCode: code,
	}
}

/*PostGroupsGroupNameJobsIDSuccessDefault handles this case with default header values.

Unexpected error
*/
type PostGroupsGroupNameJobsIDSuccessDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post groups group name jobs ID success default response
func (o *PostGroupsGroupNameJobsIDSuccessDefault) Code() int {
	return o._statusCode
}

func (o *PostGroupsGroupNameJobsIDSuccessDefault) Error() string {
	return fmt.Sprintf("[POST /groups/{group_name}/jobs/{id}/success][%d] PostGroupsGroupNameJobsIDSuccess default  %+v", o._statusCode, o.Payload)
}

func (o *PostGroupsGroupNameJobsIDSuccessDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
