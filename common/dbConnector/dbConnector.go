package dbconnector

import (
	"log"
	"os"

	"github.com/TAIM-Studio/Debita-Backend/common/datamodel"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DebitaDatabase *gorm.DB

func InitializeConnection(readyChannel chan bool) {
	dsn := "root:root@tcp(" + os.Getenv("DATABASE_HOST") + ")/debita?parseTime=true"
	gDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected!")
	db, _ := gDB.DB()

	if errPing := db.Ping(); errPing != nil {
		log.Fatal("Database not reachable")
	}
	DebitaDatabase = gDB
	setupDatabase()
	readyChannel <- true

}

func setupDatabase() {
	DebitaDatabase.AutoMigrate(&datamodel.Abo{})
	DebitaDatabase.AutoMigrate(&datamodel.Debt{})
	DebitaDatabase.AutoMigrate(&datamodel.Debtor{})
	DebitaDatabase.AutoMigrate(&datamodel.Payement{})
}
