package tests

import (
	"testing"
	json "github.com/vkatari10/trading-bot/src/runtime/go-src/json"
	technicals "github.com/vkatari10/trading-bot/src/runtime/go-src/technicals"
)

func TestCrossover(t *testing.T) {
	data, err := json.NewRuntimeData("../test.json")
	if err != nil {
		t.Errorf("TestCrossover: Could not make runtime object")
	}

	relationshipTestUp := technicals.Relationship {
		Col1Val: 0.5,
		Col2Val: 0.9,
	}

	relationshipTestDown := technicals.Relationship {
		Col1Val: 100,
		Col2Val: 50,
	}


	testCrossUp := []float64{0.4, 0.3}
	testCrossDown := []float64{50, 75}

	check := technicals.CrossoverRelationship(&relationshipTestUp, &data, testCrossUp)

	check2 := technicals.CrossoverRelationship(&relationshipTestDown,
	&data, testCrossDown)

	if check != 1 || check2 != -1 {
		t.Errorf("Computation failed for TestCrossover")
	}

}