// Package security provides private security methods.
package security

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Supported SQL dialects.
const (
	DialectPostgreSQL = "postgres"
	DialectMySQL      = "mysql"
	DialectSQLite     = "sqlite"
)

var db *gorm.DB

// InitDB initializes database used to store users.
//
// List of supported dialects:
//   - `postgres` for PostgreSQL,
//   - `mysql` for MySQL,
//   - `sqlite` for SQLite (default).
//
// If database initialization fails, InitDB panics.
func InitDB(dsn, dialect string) {
	var (
		dialector gorm.Dialector
		err       error
	)

	switch dialect {
	case DialectPostgreSQL:
		dialector = postgres.New(postgres.Config{
			DSN: dsn,
		})
	case DialectMySQL:
		dialector = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		dialector = sqlite.New(sqlite.Config{
			DSN: dsn,
		})
	}

	db, err = gorm.Open(dialector)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
}

// InitDBWithDialector initializes database used to store users. Unlike
// [InitDB], it accepts GORM dialector, allowing dependency injection.
func InitDBWithDialector(dialector gorm.Dialector) {
	var err error
	db, err = gorm.Open(dialector)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
}

// Credentials represents authorization data.
// It contains Role and Username.
type Credentials struct {
	Roles    []string
	Username string
}

// User is a database model of user.
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Roles    string // Roles are comma-separated.
}

// TableName overrides the table name used by User to `hakutest_users`.
func (User) TableName() string {
	return "hakutest_users"
}

// Credentials returns [*Credentials] associated with the user.
func (m *User) Credentials() *Credentials {
	var roles []string

	if strings.TrimSpace(m.Roles) != "" {
		roles = strings.Split(m.Roles, ",")
	}

	return &Credentials{
		Username: m.Username,
		Roles:    roles,
	}
}

// CreateUser creates a new [User] entry in the database.
func CreateUser(username, password string, roles []string) error {
	u := User{
		Username: username,
		Password: password,
		Roles:    strings.Join(roles, ","),
	}

	return db.Create(&u).Error
}

// Login retrieves first [User] with matching username and password. It returns
// [*Credentials] of the found user.
//
// If user does not exist, error is returned.
func Login(username, password string) (*Credentials, error) {
	var u User

	res := db.First(&u, "username = ? AND password = ?", username, password)
	if res.Error != nil {
		return nil, res.Error
	}

	return u.Credentials(), nil
}

// DeleteUser deletes users by id from the database.
func DeleteUser(ids ...int) error {
	return db.Delete(&User{}, ids).Error
}
