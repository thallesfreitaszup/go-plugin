package task

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)
type MockRepository struct {
	mock.Mock
}

func (m MockRepository) Create(Task Task) (int64, error){
	return 0, nil
}


func (m MockRepository) Update(Task Task) {

}

func (m MockRepository) Delete(Task Task) {

}

func (m MockRepository) Get(Task Task) {
}

func TestTaskCreateService(t *testing.T){
	mockRepository := new(MockRepository)
	service := Service{
		Repository: mockRepository,
	}
	mockRepository.On("Create").Return(0, nil )
	Task:= Task{
		Name: "mock-task",
		Status: "DOING",
	}
	intValue, err :=service.Create(Task)
	assert.Equal(t, intValue, int64(0))
	assert.Equal(t, err, nil)
}
