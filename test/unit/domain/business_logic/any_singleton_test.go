package business_logic_test

import (
	"space_traveler/src/domain/business_logic"
	"testing"
)

func TestSingleton(testing *testing.T) {
	for i := 0; i < 10; i++ {
		business_logic.GetDao()
	}
}
