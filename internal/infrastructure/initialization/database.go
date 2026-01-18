package initialization

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	gorm *gorm.DB
	sql  *sql.DB
}

func InitDatabase(cfg config.PostgreSQLConfig) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.DBName,
		cfg.User,
		cfg.Password,
		cfg.SSLMode,
	)

	gormLogger := logger.New(
		log.New(os.Stdout, "[DB] ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
			ParameterizedQueries:      true,
		},
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 gormLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	db := &Database{
		gorm: gormDB,
		sql:  sqlDB,
	}

	if err := db.migrate(); err != nil {
		db.Close()
		return nil, err
	}

	return &Database{
		gormDB,
		sqlDB,
	}, nil
}

func (d *Database) ORM() *gorm.DB {
	return d.gorm
}

func (d *Database) Close() {
  _ = d.sql.Close()
}

func (d *Database) migrate() error {
	models := []any{
		&model.Department{},
		&model.User{},
		&model.Token{},
	}

	return d.gorm.AutoMigrate(models...)
}
