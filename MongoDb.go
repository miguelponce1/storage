package storage

type MongoStorage struct {
}

func (ms *MongoStorage) SetAlumno() error {
	return nil
}

func (ms *MongoStorage) SetProfesor() error {
	return nil
}

func (ms *MongoStorage) SetTurno() error {
	return nil
}
func (ms *MongoStorage) SetHorario() error {
	return nil
}
func (ms *MongoStorage) GetAlumnto() {
	return
}
func (ms *MongoStorage) GetProfesor() {
	return
}

func NewMongoDb() *MongoStorage {
	ms := MongoStorage{}
	return &ms
}
