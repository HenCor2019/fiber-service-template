package services_test

import (
	"testing"

	health_svc "github.com/HenCor2019/fiber-service-template/internal/health/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDependencyService struct {
	mock.Mock
}

func (m *MockDependencyService) SomeMethod() string {
	args := m.Called()
	return args.String(0)
}

func TestCheckHealth(t *testing.T) {
	service := health_svc.New()

	t.Run("Should return OK", func(t *testing.T) {
		t.Parallel()
		result := service.CheckHealth()
		assert.Equal(t, "ok", result)
	})
}
