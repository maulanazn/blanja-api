package builder

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	db := config.GetConnection()

	db.AutoMigrate(&entity.Customer{})
	myBornDate := time.Now().AddDate(2003, 9, 19)

	result := &entity.Customer{UserImage: "https://image.com", UserName: "maulana", Email: "maulanazn@mail.com", Phone: 8298392, Gender: "Man", DateofBirth: myBornDate, Password: "maulanazn123"}

	db.Create(&result)

	assert.Equal(t, "maulana", result.UserName, "Not valid")
}
