package test_data

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/cheggaaa/pb/v3"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"social/internal/domain/user"
)

const countBuildings = 1000000
const batchSize = 5000
const minAge = 18
const maxAge = 30

type generator struct {
	db *gorm.DB
}

func NewGenerator(db *gorm.DB) *generator {
	db.Logger = newLogger()
	return &generator{
		db: db,
	}
}

func newLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
}

func (m generator) GenerateTestData() (err error) {
	log.Println("generate buildings")
	if err := m.generateBuildings(); err != nil {
		return err
	}

	return nil
}
func New() *boolgen {
	return &boolgen{src: rand.NewSource(time.Now().UnixNano())}
}
func (m generator) generateBuildings() error {
	users := make([]user.User, 0, batchSize)
	boolGen := newBoolGenerator()

	bar := pb.StartNew(countBuildings)
	for i := 0; i < countBuildings; i += 5000 {
		for j := 0; j < batchSize; j++ {
			age := uint(rand.Intn(maxAge-minAge) + minAge)
			users = append(users, user.User{
				Name: gofakeit.Name(),
				Age: age,
				Sex: boolGen.Bool(),
				City: gofakeit.City(),
				Interest: gofakeit.Phrase(),
			})

			err := m.db.Create(&users).Error
			if err != nil {
				return err
			}

			bar.Add(len(users))

			users = make([]user.User, 0, batchSize)
		}
	}
	bar.Finish()

	return nil
}