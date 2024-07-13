package database

import (
	"github.com/dpcamargo/fullcycle-api/internal/entity"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err = db.AutoMigrate(&entity.User{}); err != nil {
		t.Error(err)
	}
	user, _ := entity.NewUser("John Doe", "email@email.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	require.NoError(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	require.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	if err = db.AutoMigrate(&entity.User{}); err != nil {
		t.Error(err)
	}
	user, _ := entity.NewUser("John Doe", "email@email.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	require.NoError(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	require.NoError(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
