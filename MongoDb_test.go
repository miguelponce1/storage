package storage

import (
	"testing"
)

func TestUsuario_NewAlumno(t *testing.T) {

	mongoDb, err := NewMongoDbStorage()
	defer mongoDb.cancel()

	defer func() {
		if err = mongoDb.client.Disconnect(mongoDb.ctx); err != nil {
			msg := "Error al desconectar de MongoDB: " + err.Error()
			t.Error(msg)
		}
	}()

	if mongoDb == nil || err != nil {
		t.Errorf("No se realizaron las operaciones en mongodb")
	}
	mongoDb.SetAlumno()

}
