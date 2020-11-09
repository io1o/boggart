// Code generated by go-swagger; DO NOT EDIT.

package monitoring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kihamo/boggart/providers/hilink/models"
)

// GetMonitoringTrafficStatisticsReader is a Reader for the GetMonitoringTrafficStatistics structure.
type GetMonitoringTrafficStatisticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMonitoringTrafficStatisticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMonitoringTrafficStatisticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMonitoringTrafficStatisticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMonitoringTrafficStatisticsOK creates a GetMonitoringTrafficStatisticsOK with default headers values
func NewGetMonitoringTrafficStatisticsOK() *GetMonitoringTrafficStatisticsOK {
	return &GetMonitoringTrafficStatisticsOK{}
}

/*GetMonitoringTrafficStatisticsOK handles this case with default header values.

Successful operation
*/
type GetMonitoringTrafficStatisticsOK struct {
	Payload *models.MonitoringTrafficStatistics
}

func (o *GetMonitoringTrafficStatisticsOK) Error() string {
	return fmt.Sprintf("[GET /api/monitoring/traffic-statistics][%d] getMonitoringTrafficStatisticsOK  %+v", 200, o.Payload)
}

func (o *GetMonitoringTrafficStatisticsOK) GetPayload() *models.MonitoringTrafficStatistics {
	return o.Payload
}

func (o *GetMonitoringTrafficStatisticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringTrafficStatistics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMonitoringTrafficStatisticsDefault creates a GetMonitoringTrafficStatisticsDefault with default headers values
func NewGetMonitoringTrafficStatisticsDefault(code int) *GetMonitoringTrafficStatisticsDefault {
	return &GetMonitoringTrafficStatisticsDefault{
		_statusCode: code,
	}
}

/*GetMonitoringTrafficStatisticsDefault handles this case with default header values.

Unexpected error
*/
type GetMonitoringTrafficStatisticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get monitoring traffic statistics default response
func (o *GetMonitoringTrafficStatisticsDefault) Code() int {
	return o._statusCode
}

func (o *GetMonitoringTrafficStatisticsDefault) Error() string {
	return fmt.Sprintf("[GET /api/monitoring/traffic-statistics][%d] getMonitoringTrafficStatistics default  %+v", o._statusCode, o.Payload)
}

func (o *GetMonitoringTrafficStatisticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetMonitoringTrafficStatisticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
