package repository

import (
	"church-adoration/appointment"
	"church-adoration/models"
	"github.com/pkg/errors"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type mongoAppointmentsRepo struct {
	Session                *mgo.Session
	AppointmentsCollection *mgo.Collection
}

func NewMongoAppointmentRepository(session *mgo.Session) appointment.AppointmentRepository {

	return &mongoAppointmentsRepo{
		Session:                session,
		AppointmentsCollection: session.DB("church-adoration").C("appointments"),
	}
}

func (m *mongoAppointmentsRepo) GetAppointments(userID string, page int) (*models.Appointments, error) {
	c := models.Appointments{}

	var err error
	if userID == "*" {
		err = m.AppointmentsCollection.Find(bson.M{}).All(&c)
	} else {
		err = m.AppointmentsCollection.Find(bson.M{
			"user_id": bson.ObjectIdHex(userID),
		}).Skip((page - 1) * 5).Limit(5).All(&c)
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (m *mongoAppointmentsRepo) GetAppointmentsAdmin() (*models.Appointments, error) {
	c := models.Appointments{}
	query := []bson.M{{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "user_id",
			"foreignField": "_id",
			"as":           "user",
		}},
	}

		pipe := m.AppointmentsCollection.Pipe(query)
		err := pipe.All(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

/*
	based on the day we click in the calendar get slots available and those already takes

*/
func (m *mongoAppointmentsRepo) GetSlotsAvailability(unixTimestamp string) (*models.Appointments, error) {
	c := models.Appointments{}

	start := time.Now().AddDate(2018, 12, 11)
	end := time.Now().AddDate(2018, 12, 12)

	err := m.AppointmentsCollection.Find(bson.M{
		"date": bson.M{"$gte": start, "$lt": end},
	}).All(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}


func (m *mongoAppointmentsRepo) AddAppointment(a *models.Appointment) (*models.Appointment, error) {
	// check if the appointment is already added
	//TODO
	/*
	var start = new Date(2017, 10, 08);
	var end = new Date(2018, 10, 08);
	var from_hour_new = "08"
	var to_hour_new = "09"
	db.getCollection('appointments').find({
	//     "date": {$gte: start, $lte: end},
   $or: [{
   $and: [
           {"from_hour": {$gte: from_hour_new}},
           {"to_hour": {$lte: to_hour_new}}
   ]},{
   $and: [
           {"from_hour": {$lte: from_hour_new}},
           {"to_hour": {$gte: to_hour_new}}
   ]}
   ]
})
	*/

	// check if this slot is already taken
	app := models.Appointments{}

	m.AppointmentsCollection.Find(bson.M{
		"$or": []bson.M{{
			"$and": []bson.M{
				{"from_hour": bson.M{ "$gte": a.FromHour}},
				{"to_hour": bson.M{ "$lte": a.ToHour}},
			},
		}, {
			"$and": []bson.M{
				{"from_hour": bson.M{ "$lte": a.FromHour}},
				{"to_hour": bson.M{ "$gte": a.ToHour}},
			},
		}},
	}).All(&app)

	if len(app) >= 1 {
		return nil, errors.New("Slot already taken")
	}

	err := m.AppointmentsCollection.Insert(&a)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (m *mongoAppointmentsRepo) UpdateAppointment(id string, a *models.Appointment) (*models.Appointment, error) {
	err := m.AppointmentsCollection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &a)

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (m *mongoAppointmentsRepo) CancelAppointment(id string, reason string) error {
	err := m.AppointmentsCollection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{
		"$set": bson.M{
			"cancelled": true,
			"cancellation_reason": reason,
	}})

	if err != nil {
		return err
	}

	return nil
}
