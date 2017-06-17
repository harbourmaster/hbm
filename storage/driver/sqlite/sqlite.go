package sqlite

import (
	"path"

	"github.com/kassisol/hbm/storage"
	"github.com/kassisol/hbm/storage/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Justifying comment
)

func init() {
	storage.RegisterDriver("sqlite", New)
}

// Config structure
type Config struct {
	DB *gorm.DB
}

// New function
func New(config string) (driver.Storager, error) {
	debug := false

	file := path.Join(config, "data.db")

	db, err := gorm.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	db.LogMode(debug)

	db.AutoMigrate(&AppConfig{}, &User{}, &Group{}, &Resource{}, &Collection{}, &Policy{})

	return &Config{DB: db}, nil
}

// End function
func (c *Config) End() {
	c.DB.Close()
}
