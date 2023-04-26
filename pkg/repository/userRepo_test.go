package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUersSignUp(t *testing.T) {
	tests := []struct {
		name           string
		input          helperStruct.UserReq
		expectedOutput response.UserData
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			name: "successful creations",
			input: helperStruct.UserReq{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "akshay",
				Email:  "akshay@gmail.com",
				Mobile: "9072001341",
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				// //The query that need to execuetd
				// query := "INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4)"
				// //The arguments for the query exicution
				// args := []driver.Value{"akshay", "akshay@gmail.com", "9072001341"}
				// // NewRows creates a new result set for use in database testing.
				// mock.ExpectQuery(query).
				// 	WithArgs(args...).
				// 	WillReturnRows(rows)
				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}).
					AddRow(1, "akshay", "akshay@gmail.comsdf", "9072001341")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("akshay", "akshay@gmail.com", "9072001341", "123456789").
					WillReturnRows(rows)
			},
			expectedErr: nil,
		},
		{
			name: "duplicate user",
			input: helperStruct.UserReq{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			expectedOutput: response.UserData{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("akshay", "akshay@gmail.com", "9072001341", "123456789").
					WillReturnError(errors.New("email or phone number alredy used"))
			},
			expectedErr: errors.New("email or phone number alredy used"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//New() method from sqlmock package create sqlmock database connection and a mock to manage expectations.
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			//initialize the db instance with the mock db connection
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.UserSignUp(context.TODO(), tt.input)
			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
