package db

import (
	"context"
	"fmt"
	"sync"

	"go-skeleton/config"

	"github.com/DATA-DOG/go-sqlmock"

	"gorm.io/driver/mysql"
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

// Init mysql client specific db
func (dbdget *dbDelegate) Init() {
	dbdget.run(true)
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
			mysql.Open(mysqlSource(withDB)),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
				Logger:                                   logger.Default.LogMode(logLevel),
			},
		)
		if err != nil {
			panic("init mysql failed: " + err.Error())
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

func mysqlSource(withDB bool) string {
	config := config.Config.DB
	if !withDB {
		config.Name = ""
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local&multiStatements=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Encoding,
	)
}
