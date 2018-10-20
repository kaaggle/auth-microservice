package usecase

import (
	"church-adoration/appointment"
	"church-adoration/core"
	"church-adoration/models"
)

type appointmentUsecase struct {
	appointmentRepo appointment.AppointmentRepository
}

func NewAppointmentUsecase(appointmentRepo appointment.AppointmentRepository) appointment.AppointmentRepository {
	return &appointmentUsecase{
		appointmentRepo: appointmentRepo,
	}
}

func (a *appointmentUsecase) GetAppointments(userID string, page int) (*models.Appointments, error) {
	c, err := a.appointmentRepo.GetAppointments(userID, page)

	if err != nil {
		return nil, err
	}

	return c, nil
}


func (a *appointmentUsecase) GetAppointmentsAdmin() (*models.Appointments, error) {
	c, err := a.appointmentRepo.GetAppointmentsAdmin()

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (a *appointmentUsecase) GetSlotsAvailability(unixTimestamp string) (*models.Appointments, error) {
	c, err := a.appointmentRepo.GetSlotsAvailability(unixTimestamp)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (a *appointmentUsecase) AddAppointment(ap *models.Appointment) (*models.Appointment, error) {
	appointment, err := a.appointmentRepo.AddAppointment(ap)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (a *appointmentUsecase) UpdateAppointment(id string, ap *models.Appointment) (*models.Appointment, error) {
	appointment, err := a.appointmentRepo.UpdateAppointment(id, ap)

	if err != nil {
		return nil, err
	}

	return appointment, nil
}

func (a *appointmentUsecase) CancelAppointment(id string, reason string) error {
	err := a.appointmentRepo.CancelAppointment(id, reason)

	if err != nil {
		return err
	}

	// send email to the admin
	err = core.SendEmail(
		"erkidhoxholli@gmail.com",
		"User cancelled the appointment",
		"Erkid Hoxholli with email: erkidhoxholli@gmail.com cancelled the appointment",
		)
	if err != nil {
		return err
	}

	return nil
}
