package covenant

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Gorm struct {
	database *gorm.DB
}

func NewGorm(migrationEntity interface{}, dsn string) (*Gorm, error) {

	database, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	database.AutoMigrate(migrationEntity)

	return &Gorm{
		database: database,
	}, nil

}

func (g *Gorm) Create(destiny interface{}) error {
	result := g.database.Create(destiny)

	return result.Error
}

func (g *Gorm) Read(destiny interface{}, conditionals ...interface{}) error {
	result := g.database.Find(destiny, conditionals...)

	return result.Error
}

func (g *Gorm) Update(destiny interface{}) error {
	result := g.database.Save(destiny)

	return result.Error
}

func (g *Gorm) Delete(destiny interface{}, conditionals ...interface{}) error {
	result := g.database.Delete(destiny, conditionals...)

	return result.Error
}
