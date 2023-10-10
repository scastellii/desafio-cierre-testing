package products

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Crear un mock de Service
type MockService struct {
	mock.Mock
}

func (m *MockService) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}

func TestHandler_GetProducts(t *testing.T) {
	// Given
	responseProduct := []Product{
		{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55},
	}
	methodName := "GetAllBySeller"
	argument := "FEX112AC"
	mockService := new(MockService)
	// Configurar el comportamiento del mock para GetAllBySeller
	mockService.On(methodName, argument).Return(responseProduct, nil)
	// Crear un Handler utilizando el mock
	handler := NewHandler(mockService)
	// Crear una solicitud falsa
	req, _ := http.NewRequest("GET", "/api/v1/products?seller_id=FEX112AC", nil)
	res := httptest.NewRecorder()
	// Crear un motor Gin y registrar la ruta
	r := gin.Default()
	r.GET("/api/v1/products", handler.GetProducts)

	//When
	r.ServeHTTP(res, req)

	//Then
	assert.Equal(t, http.StatusOK, res.Code)
	// Verificar el cuerpo de la respuesta
	expectedResponse := `[{"ID":"mock","SellerID":"FEX112AC","Description":"generic product","Price":123.55}]`
	assert.JSONEq(t, expectedResponse, res.Body.String())
	// Verificar que se llamó a GetAllBySeller en el mock
	mockService.AssertExpectations(t)
}

func TestHandler_GetProducts_Error(t *testing.T) {
	//Given
	// Crear un mock de Service con error
	var responseProduct []Product
	methodName := "GetAllBySeller"
	argument := "FEX112AC"
	mockService := new(MockService)
	errorMocked := errors.New("Error en el servicio")
	// Configurar el comportamiento del mock para GetAllBySeller
	path := "/api/v1/products?seller_id=FEX112AC"
	mockService.On(methodName, argument).Return(responseProduct, errorMocked)
	handler := NewHandler(mockService)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.GET("api/v1/products", handler.GetProducts)

	//When
	r.ServeHTTP(res, req)

	//Then
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	expectedResponse := `{"error":"Error en el servicio"}`
	assert.JSONEq(t, expectedResponse, res.Body.String())
	mockService.AssertExpectations(t)
}

func TestHandler_GetProducts_Id_Query_Param(t *testing.T) {
	//Given
	// Crear un mock de Service con error
	var responseProduct []Product
	methodName := "GetAllBySeller"
	argument := "FEX112AC"
	mockService := new(MockService)
	errorMocked := errors.New("Error en el servicio")
	mockService.On(methodName, argument).Return(responseProduct, errorMocked)
	// Crear un Handler mocked
	handler := NewHandler(mockService)
	// Crear una solicitud falsa
	path := "/api/v1/products"
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.GET(path, handler.GetProducts)

	//When
	r.ServeHTTP(res, req)

	//Then
	assert.Equal(t, http.StatusBadRequest, res.Code)
	expectedResponse := `{"error":"seller_id query param is required"}`
	assert.JSONEq(t, expectedResponse, res.Body.String())
}

func TestHandler_GetProductsFunctional(t *testing.T) {
	// Given
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)
	// Crear una solicitud falsa
	req, _ := http.NewRequest("GET", "/api/v1/products?seller_id=FEX112AC", nil)
	res := httptest.NewRecorder()
	// Crear un motor Gin y registrar la ruta
	r := gin.Default()
	r.GET("/api/v1/products", handler.GetProducts)

	//When
	r.ServeHTTP(res, req)

	//Then
	assert.Equal(t, http.StatusOK, res.Code)
	// Verificar el cuerpo de la respuesta
	expectedResponse := `[{"ID":"mock","SellerID":"FEX112AC","Description":"generic product","Price":123.55}]`
	assert.JSONEq(t, expectedResponse, res.Body.String())
}

func TestHandler_GetProducts_Error_funcional(t *testing.T) {
	//Given
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)
	// Configurar el comportamiento del mock para GetAllBySeller
	path := "/api/v1/products?seller_id=111"
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.GET("api/v1/products", handler.GetProducts)

	//When
	r.ServeHTTP(res, req)

	//Then
	assert.Equal(t, http.StatusInternalServerError, res.Code)
	expectedResponse := `{"error":"Error en el repositorio"}`
	assert.JSONEq(t, expectedResponse, res.Body.String())
}

func TestHandler_GetProducts_Id_Query_Param_Funcional(t *testing.T) {
	//Given
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)
	// Crear una solicitud falsa
	path := "/api/v1/products"
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.GET(path, handler.GetProducts)

	//When
	r.ServeHTTP(res, req)

	//Then
	assert.Equal(t, http.StatusBadRequest, res.Code)
	expectedResponse := `{"error":"seller_id query param is required"}`
	assert.JSONEq(t, expectedResponse, res.Body.String())
}
