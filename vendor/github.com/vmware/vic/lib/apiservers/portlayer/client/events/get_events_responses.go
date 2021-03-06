package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetEventsReader is a Reader for the GetEvents structure.
type GetEventsReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *GetEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventsOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetEventsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsOK creates a GetEventsOK with default headers values
func NewGetEventsOK(writer io.Writer) *GetEventsOK {
	return &GetEventsOK{
		Payload: writer,
	}
}

/*GetEventsOK handles this case with default header values.

OK
*/
type GetEventsOK struct {
	Payload io.Writer
}

func (o *GetEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] getEventsOK  %+v", 200, o.Payload)
}

func (o *GetEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventsInternalServerError creates a GetEventsInternalServerError with default headers values
func NewGetEventsInternalServerError() *GetEventsInternalServerError {
	return &GetEventsInternalServerError{}
}

/*GetEventsInternalServerError handles this case with default header values.

Failed to get events
*/
type GetEventsInternalServerError struct {
}

func (o *GetEventsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /events][%d] getEventsInternalServerError ", 500)
}

func (o *GetEventsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
