package appointment

import "church-adoration/models"

type AppointmentUsecase interface {
	AddAppointment(*models.Appointment) (*models.Appointment, error)
	UpdateAppointment(string, *models.Appointment) (*models.Appointment, error)
	GetAppointments(string, int) (*models.Appointments, error)
	GetAppointmentsAdmin() (*models.Appointments, error)
	GetSlotsAvailability(string) (*models.Appointments, error)
	CancelAppointment(id string, reason string) error
}
