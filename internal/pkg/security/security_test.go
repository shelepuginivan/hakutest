package security

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mock sqlmock.Sqlmock

func setup() func() {
	var (
		conn *sql.DB
		err  error
	)

	conn, mock, err = sqlmock.New()
	if err != nil {
		panic(err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      conn,
		DriverName:                "mysql",
		SkipInitializeWithVersion: true,
	})

	db, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return func() {
		conn.Close()
	}
}

func TestCreateUser(t *testing.T) {
	defer setup()()

	mock.ExpectBegin()

	mock.
		ExpectExec("INSERT INTO `hakutest_users`").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := CreateUser("John Doe", "password", []string{"student"})
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUser_Credentials(t *testing.T) {
	tests := []struct {
		user *User
		want *Credentials
	}{
		{
			user: &User{
				Username: "John Doe",
				Password: "123456",
				Roles:    "student",
			},
			want: &Credentials{
				Username: "John Doe",
				Roles:    []string{"student"},
			},
		},
		{
			user: &User{
				Username: "Jane Smith",
				Password: "abcdef",
				Roles:    "teacher,admin",
			},
			want: &Credentials{
				Username: "Jane Smith",
				Roles:    []string{"teacher", "admin"},
			},
		},
		{
			user: &User{
				Username: "Alice Johnson",
				Password: "password123",
				Roles:    "student",
			},
			want: &Credentials{
				Username: "Alice Johnson",
				Roles:    []string{"student"},
			},
		},
		{
			user: &User{
				Username: "Bob Brown",
				Password: "qwerty",
				Roles:    "",
			},
			want: &Credentials{
				Username: "Bob Brown",
				Roles:    nil,
			},
		},
		{
			user: &User{
				Username: "Charlie Black",
				Password: "letmein",
				Roles:    "admin,editor",
			},
			want: &Credentials{
				Username: "Charlie Black",
				Roles:    []string{"admin", "editor"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.user.Username, func(t *testing.T) {
			got := tt.user.Credentials()

			assert.Equal(t, tt.want, got)
		})
	}
}
