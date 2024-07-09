package must

import (
	"fmt"

	"github.com/tonghia/go-challenge-transaction-app/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMySQL(cfg *config.MySQL) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.FormatDSN()), &gorm.Config{
		TranslateError: true,
		// Logger:         zapgorm.New(ll.Logger).LogMode(cfg.GormLogLevel()),
	})
	if err != nil {
		return nil, fmt.Errorf("error open mysql: %w", err)
	}

	if err := db.Raw("SELECT 1").Error; err != nil {
		return nil, fmt.Errorf("error querying SELECT 1: %w", err)
	}

	return db, nil
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
