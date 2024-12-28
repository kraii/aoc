package nineteen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	workflows, parts := parse("nineteen/19-ex.txt")
	fmt.Printf("wf %+v \n", workflows)
	fmt.Printf("parts %+v \n", parts)
}

func TestSumAcceptedExample(t *testing.T) {
	assert.Equal(t, 19114, sumAccepted("nineteen/19-ex.txt"))
}

func TestSumAccepted(t *testing.T) {
	println(sumAccepted("nineteen/input.txt"))
}

func TestSumPermutationsExample(t *testing.T) {
	assert.Equal(t, int64(167409079868000), sumPermutations("nineteen/19-ex.txt"))
}

func TestSumPermutationsExample2(t *testing.T) {
	assert.Equal(t, int64(1), sumPermutations("nineteen/19-ex2.txt"))
}

func TestSumPermutations(t *testing.T) {
	println(sumPermutations("nineteen/input.txt"))
}
