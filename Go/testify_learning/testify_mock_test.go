package testify_learning

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type StudentInterface interface {
	GetStudentById(id int) string
}

type ServiceImpl struct {
	StudentInterface
}

func (service ServiceImpl) GetStudentById(id int) string {
	getRes := service.StudentInterface.GetStudentById(id)
	return getRes
}

type StudentDaoMock struct {
	mock.Mock
}

func (studentDao StudentDaoMock) GetStudentById(id int) string {
	arguments := studentDao.Called(id)
	return arguments.String(0)
}

func TestStudentInterface(t *testing.T) {
	assert := require.New(t)

	//mock logic, bind the input and the expected output to the method
	studentDaoMock := StudentDaoMock{}
	//studentDaoMock.On("GetStudentById", 0).Return("123")
	studentDaoMock.On("GetStudentById", 0).Return("htx")
	studentDaoMock.On("GetStudentById", -1).Return("")
	studentDaoMock.On("GetStudentById", 1).Return("alj")

	//based on the mock object, create the service object
	serviceImpl := ServiceImpl{studentDaoMock}

	//test case logic，table-drived test
	testCases := map[int]string{0: "htx", 1: "alj", -1: ""}

	for input, expectedOutput := range testCases {
		getRes := serviceImpl.GetStudentById(input)
		assert.Equal(expectedOutput, getRes)
	}
}
