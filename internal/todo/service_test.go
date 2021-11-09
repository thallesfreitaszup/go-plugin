package todo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)
type MockRepository struct {
	mock.Mock
}

func (m MockRepository) Create(todo Todo) (int64, error){
	return 0, nil
}


func (m MockRepository) Update(todo Todo) {

}

func (m MockRepository) Delete(todo Todo) {

}

func (m MockRepository) Get(todo Todo) {
}

func TestTodoCreateService(t *testing.T){
	mockRepository := new(MockRepository)
	service := Service{
		Repository: mockRepository,
	}
	mockRepository.On("Create").Return(0, nil )
	todo:= Todo{
		Name: "mock-task",
		Status: "DOING",
	}
	intValue, err :=service.Create(todo)
	assert.Equal(t, intValue, int64(0))
	assert.Equal(t, err, nil)
}
