package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/dpcamargo/fullcycle-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err = db.AutoMigrate(&entity.Product{}); err != nil {
		t.Error(err)
	}
	product, _ := entity.NewProduct("Product 1", 10.0)
	require.NoError(t, err)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	require.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindALlProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err = db.AutoMigrate(&entity.Product{}); err != nil {
		t.Error(err)
	}

	for i := range 30 {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i+1), rand.Float64()*100)
		require.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err = db.AutoMigrate(&entity.Product{}); err != nil {
		t.Error(err)
	}
	product, err := entity.NewProduct("Product 1", 10.0)
	require.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err = db.AutoMigrate(&entity.Product{}); err != nil {
		t.Error(err)
	}
	product, err := entity.NewProduct("Product 1", 10.0)
	require.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	require.NoError(t, err)
	product.Name = "Product 2"
	err = productDB.Update(product)
	require.NoError(t, err)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct("Product 1", 10.0)
	require.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	require.NoError(t, err)

	product, err = productDB.FindByID(product.ID.String())
	require.Error(t, err)
	assert.Equal(t, product, &entity.Product{})
}
