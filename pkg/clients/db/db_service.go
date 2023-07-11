package db

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"sync"

	"go/skeleton/config"

	"github.com/DATA-DOG/go-sqlmock"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBGormDelegate interface {
	Init()
	InitNoUse()
	Get(ctx context.Context) *gorm.DB
	GetMock() sqlmock.Sqlmock
	BeginTx() *gorm.DB
	Rollback()
	Commit()
}

type dbDelegate struct {
	dbGorm *gorm.DB
	once   sync.Once
	debug  bool
	tx     *gorm.DB
}

func NewDBdelegate(debug bool) DBGormDelegate {
	return &dbDelegate{
		debug: debug,
	}
}

// Init database client specific db
func (dbdget *dbDelegate) Init() {
	runMigration()
	dbdget.run(true)
}

func runMigration() {
	db := config.Config.DB
	m, err := migrate.New("file://cmd/apiserver/app/migrations", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.Username, db.Password, db.Host, db.Port, db.Name))
	if err != nil {
		panic("migration error: " + err.Error())
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic("migration error: " + err.Error())
	} else {
		nowVersion, _, _ := m.Version()
		log.Println("migration running at version:", nowVersion)
	}
}

// InitNoUse client not specific db
func (dbdget *dbDelegate) InitNoUse() {
	dbdget.run(false)
}

func (dbdget *dbDelegate) GetMock() sqlmock.Sqlmock {
	return nil
}

func (dbdget *dbDelegate) run(withDB bool) {
	dbdget.once.Do(func() {
		var logLevel logger.LogLevel
		if config.Config.DB.Debug {
			logLevel = logger.Info
		} else {
			logLevel = logger.Silent
		}

		var err error
		dbdget.dbGorm, err = gorm.Open(
			postgres.Open(dbSource(withDB)),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
				Logger:                                   logger.Default.LogMode(logLevel),
			},
		)
		if err != nil {
			panic("init database failed: " + err.Error())
		}

		if dbdget.debug {
			dbdget.dbGorm = dbdget.dbGorm.Debug()
		}
	})
}

func (dbdget *dbDelegate) Get(ctx context.Context) *gorm.DB {
	tx := ctx.Value("tx")
	if tx != nil {
		return tx.(*gorm.DB)
	}

	return dbdget.dbGorm
}

// new transactions

func (dbdget *dbDelegate) BeginTx() *gorm.DB {
	return dbdget.dbGorm.Begin()
}

func (dbdget *dbDelegate) Rollback() {
	dbdget.dbGorm.Rollback()
}

func (dbdget *dbDelegate) Commit() {
	dbdget.dbGorm.Commit()
}

func dbSource(withDB bool) string {
	config := config.Config.DB
	if !withDB {
		config.Name = ""
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s",
		config.Host,
		config.Username,
		config.Password,
		config.Name,
		config.Port,
		config.TimeZone,
	)
}
