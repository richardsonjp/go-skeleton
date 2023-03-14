package mock

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-skeleton/pkg/clients/db"
)

type dbDelegateMock struct {
	dbMock sqlmock.Sqlmock
	dbGorm *gorm.DB
	once   sync.Once
	tx     *gorm.DB
}

func NewDBdelegate() db.DBGormDelegate {
	return &dbDelegateMock{}
}

// Init build client include database
func (dbdget *dbDelegateMock) Init() {
	dbdget.run()
}

// InitNoUse build client not include database
func (dbdget *dbDelegateMock) InitNoUse() {
	dbdget.run()
}

func (dbdget *dbDelegateMock) GetMock() sqlmock.Sqlmock {
	return dbdget.dbMock
}

func (dbdget *dbDelegateMock) run() {
	dbdget.once.Do(func() {
		var err error
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherFunc(customMacher)))
		if err != nil {
			panic("init mock failed: " + err.Error())
		}

		dbdget.dbMock = mock
		TestInitDB(mock)
		dbdget.dbGorm, err = gorm.Open(
			mysql.New(mysql.Config{
				Conn: db,
			}),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)
		if err != nil {
			panic("init mysql failed: " + err.Error())
		}
	})
}

func (dbdget *dbDelegateMock) Get(ctx context.Context) *gorm.DB {
	return dbdget.dbGorm
}

func (dbdget *dbDelegateMock) GetWithContext(ctx context.Context) *gorm.DB {
	tx := ctx.Value("tx")
	if tx != nil {
		return tx.(*gorm.DB)
	}

	return dbdget.dbGorm
}

//func (dbdget *dbDelegateMock) Paginate(value interface{}, pagination *model.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
//	return nil
//}

func (dbdget *dbDelegateMock) BeginTx() *gorm.DB {
	return dbdget.dbGorm.Begin()
}

func (dbdget *dbDelegateMock) Rollback() {
	dbdget.dbGorm.Rollback()
}

func (dbdget *dbDelegateMock) Commit() {
	dbdget.dbGorm.Commit()
}

var re = regexp.MustCompile("\\s+")

func customMacher(expectedSQL, actualSQL string) error {
	expect := stripQuery(expectedSQL)
	actual := stripQuery(actualSQL)
	if actual != expect {
		return fmt.Errorf("\nactual sql: \"%s\"\nexpect sql: \"%s\"", actual, expect)
	}
	return nil
}

func stripQuery(q string) (s string) {
	return strings.TrimSpace(re.ReplaceAllString(q, " "))
}

func TestInitDB(dbmock sqlmock.Sqlmock) {
	dbmock.
		ExpectQuery("SELECT VERSION()").
		WillReturnRows(
			sqlmock.NewRows([]string{"VERSION()"}).
				AddRow("8.0.23"),
		)
}
