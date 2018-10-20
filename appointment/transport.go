package appointment

import (
	"church-adoration/models"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type HttpAppointmentHandler struct {
	AppointmentUsecase AppointmentUsecase
}

func (h *HttpAppointmentHandler) GetAppointments(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userID := q.Get("user_id")
	page, _ := strconv.Atoi(q.Get("page"))
	if userID == ""{
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: "There is no user_id param",
			Data: nil,
		})
		return
	}
	c, err := h.AppointmentUsecase.GetAppointments(userID, page)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Retrieved all appointments.",
		Data:    &c,
	})
}

func (h *HttpAppointmentHandler) GetAppointmentsAdmin(w http.ResponseWriter, r *http.Request) {
	c, err := h.AppointmentUsecase.GetAppointmentsAdmin()
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Retrieved all appointments.",
		Data:    &c,
	})
}

func (h *HttpAppointmentHandler) GetSlotsAvailability(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	date := q.Get("date")

	if date == "" {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: "There is no date param",
			Data: nil,
		})
		return
	}

	c, err := h.AppointmentUsecase.GetSlotsAvailability(date)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Retrieved all slot available.",
		Data:    &c,
	})
}


func (h *HttpAppointmentHandler) AddAppointment(w http.ResponseWriter, r *http.Request) {
	c := models.Appointment{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}
	c.CreatedAt = time.Now()

	err = c.Validate()
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	appointment, err := h.AppointmentUsecase.AddAppointment(&c)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Added appointment.",
		Data:    &appointment,
	})
}

func (h *HttpAppointmentHandler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: "This is not a id",
			Data: nil,
		})
		return
	}

	c := models.Appointment{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}
	c.UpdatedAt = time.Now()

	err = c.Validate()
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	appointment, err := h.AppointmentUsecase.UpdateAppointment(id, &c)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Updated appointment.",
		Data:    &appointment,
	})
}

func (h *HttpAppointmentHandler) CancelAppointment(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: "This is not a id",
			Data: nil,
		})
		return
	}

	c := models.CancelAppointment{}

	json.NewDecoder(r.Body).Decode(&c)

	err := c.Validate()
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	err = h.AppointmentUsecase.CancelAppointment(id, c.CancellationReason)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Appointment cancelled.",
		Data:    nil,
	})
}

func NewAppointmentHttpHandler(r *chi.Mux, su AppointmentUsecase) {
	handler := HttpAppointmentHandler{
		AppointmentUsecase: su,
	}

	r.Get("/nu/appointments", handler.GetAppointments)
	r.Get("/a/appointments", handler.GetAppointmentsAdmin)
	r.Put("/nu/appointments/{id}", handler.UpdateAppointment)
	r.Post("/nu/appointments", handler.AddAppointment)
	r.Post("/nu/appointments/cancel/{id}", handler.CancelAppointment)
	r.Get("/nu/appointments/slots", handler.GetSlotsAvailability)

}
