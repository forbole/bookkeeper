package tables_test

import (
	"testing"
	"github.com/stretchr/testify/suite"
  )
  
  // We'll be able to store suite-wide
  // variables and add methods to this
  // test suite struct
  type CosmosTableTestSuite struct {
	  suite.Suite
  }

  
  // We need this function to kick off the test suite, otherwise
  // "go test" won't know about our tests
  func TestCosmosTableTestSuite(t *testing.T) {
	  suite.Run(t, new(CosmosTableTestSuite))
  }