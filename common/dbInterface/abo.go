package dbInterface

import (
	"log"

	"github.com/TAIM-Studio/Debita-Backend/common/datamodel"
	dbconnector "github.com/TAIM-Studio/Debita-Backend/common/dbConnector"
	"github.com/google/uuid"
)

func CreateAbo(abo *datamodel.Abo) (*datamodel.Abo, error) {
	id, _ := uuid.NewRandom()
	abo.Id = id
	result := dbconnector.DebitaDatabase.Create(&abo)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return abo, nil
}

func GetAllAbos() (*[]datamodel.Abo, error) {
	var abos []datamodel.Abo
	result := dbconnector.DebitaDatabase.Find(&abos)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &abos, nil
}
