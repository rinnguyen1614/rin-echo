package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v4"
	"github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func newUserRepositoryTest() {

}

func TestUserRepository_First(t *testing.T) {
	testCases := []struct {
		name         string
		data         domain.User
		buildQueries func(mock sqlmock.Sqlmock, data domain.User)
		expectedErr  error
	}{
		{
			name: "Success",
			data: randomUser(t),
			buildQueries: func(mock sqlmock.Sqlmock, data domain.User) {
				mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE "id" = ? (.+)`).
					WithArgs(data.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "full_name"}).
						AddRow(data.ID, data.Username, data.FullName))
			},
		},
		{
			name: "Not found",
			data: randomUser(t),
			buildQueries: func(mock sqlmock.Sqlmock, data domain.User) {
				mock.ExpectQuery(`SELECT (.+) FROM "users" WHERE "id" = ? (.+)`).
					WithArgs(data.ID).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			expectedErr: gorm.ErrRecordNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := newDbMock()
			require.NoError(t, err, "Failed to initialize mock DB")

			testCase.buildQueries(mock, testCase.data)

			repo := NewUserRepository(db)
			got := domain.User{}
			err = repo.FirstID(&got, testCase.data.ID, nil)

			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, testCase.data, got)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUserRepository_ChangePhone(t *testing.T) {
	testCases := []struct {
		name                               string
		id                                 uint
		phone, phoneVerificationCodeHashed string
		buildQueries                       func(mock sqlmock.Sqlmock, id uint, phone, phoneVerificationCodeHashed string)
		expectedErr                        error
	}{
		{
			name:                        "Success",
			id:                          1,
			phone:                       faker.Phonenumber(),
			phoneVerificationCodeHashed: faker.Password(),
			buildQueries: func(mock sqlmock.Sqlmock, id uint, phone, phoneVerificationCodeHashed string) {
				mock.ExpectBegin()
				// expect user phone update
				mock.ExpectExec(`UPDATE "users" SET "phone"=([$]\w+),"phone_verification_code_hashed"=([$]\w+),"phone_verified"=([$]\w+) WHERE "id" = ([$]\w+)$`).
					WithArgs(phone, phoneVerificationCodeHashed, false, id).
					WillReturnResult(sqlmock.NewResult(0, 1)) // no insert id, 1 affected row

				mock.ExpectCommit()
			},
		},
		{
			name:                        gorm.ErrPrimaryKeyRequired.Error(),
			id:                          0,
			phone:                       faker.Phonenumber(),
			phoneVerificationCodeHashed: faker.Password(),
			buildQueries: func(mock sqlmock.Sqlmock, id uint, phone, phoneVerificationCodeHashed string) {
				mock.ExpectBegin()
				// expect user phone update
				mock.ExpectExec(`UPDATE "users" SET "phone"=([$]\w+),"phone_verification_code_hashed"=([$]\w+),"phone_verified"=([$]\w+) WHERE "id" = ([$]\w+)$`).
					WithArgs(phone, phoneVerificationCodeHashed, false, 0).
					WillReturnError(gorm.ErrPrimaryKeyRequired) // no insert id, 1 affected row

				mock.ExpectRollback()
			},
			expectedErr: gorm.ErrPrimaryKeyRequired,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := newDbMock()
			require.NoError(t, err, "Failed to initialize mock DB")

			testCase.buildQueries(mock, testCase.id, testCase.phone, testCase.phoneVerificationCodeHashed)

			repo := NewUserRepository(db)
			err = repo.ChangePhone(testCase.id, testCase.phone, testCase.phoneVerificationCodeHashed)

			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func randomUser(t *testing.T) domain.User {
	testData := domain.User{
		Username: faker.Username(),
		FullName: faker.Name(),
	}
	testData.ID = 1
	return testData
}
