package products

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// Crear un mock de Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}

func TestService_GetAllBySeller(t *testing.T) {
	//Given
	responseProduct := []Product{
		{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55},
	}
	methodName := "GetAllBySeller"
	argument := "FEX112AC"
	mockRepo := new(MockRepository)
	// Configurar el comportamiento del mock para GetAllBySeller
	mockRepo.On(methodName, argument).Return(responseProduct, nil)
	// Configurar el comportamiento del mock para GetAllBySeller
	svc := NewService(mockRepo)

	//When
	products, err := svc.GetAllBySeller(argument)

	//Then
	// Verificar que se llamó a GetAllBySeller en el mock
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(products))
	assert.Equal(t, "mock", products[0].ID)
	assert.Equal(t, argument, products[0].SellerID)
}

func TestService_GetAllBySeller_Error(t *testing.T) {
	//Given
	// Crear un mock de Repository que devuelve un error
	mockRepo := new(MockRepository)
	var responseProduct []Product
	methodName := "GetAllBySeller"
	argument := "111"
	errorMocked := errors.New("Error en el repositorio")
	mockRepo.On(methodName, argument).Return(responseProduct, errorMocked)
	svc := NewService(mockRepo)

	//When
	_, err := svc.GetAllBySeller("111")

	//Then
	// Verificar que se llamó a GetAllBySeller en el mock
	mockRepo.AssertExpectations(t)
	assert.Equal(t, "Error en el repositorio", err.Error())

}
