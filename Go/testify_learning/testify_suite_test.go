package testify_learning

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	flag bool
}

func (suite *TestSuite) TestMethod1() {
	assert.Equal(suite.T(), 2, 2, "2!=2")
}

func (suite *TestSuite) TestMethod2() {
	assert.Equal(suite.T(), suite.flag, true, "suite.flag == true")
}

func (suite *TestSuite) TestMethod3() {
	assert.Equal(suite.T(), "", "")
}

// which will run before each test in this suite
func (suite *TestSuite) SetupTest() {
	suite.flag = true
}

// which will run after each test in this suite
func (suite *TestSuite) TearDownTest() {
	suite.flag = false
}

// which will run before all the tests of this suite
func (suite TestSuite) SetupSuite() {

}

// which will run after all the tests of this suite
func (suite TestSuite) TearDownSuite() {

}

// use the test suite,
func TestSuiteMethods(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
