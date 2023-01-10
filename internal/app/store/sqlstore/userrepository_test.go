package sqlstore_test

import (
	"testing"

	"github.com/AisanaOlpan/GoProject/internal/app/model"
	"github.com/AisanaOlpan/GoProject/internal/app/store"
	"github.com/AisanaOlpan/GoProject/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)
	email := "ajsanaolpan123@mail.ru"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u1 := model.TestUser(t)
	u1.Email = email
	s.User().Create(u1)
	u2, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	//fmt.Print(u2)

}
