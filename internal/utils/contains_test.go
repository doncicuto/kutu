package utils_test

import (
	"testing"

	"github.com/muultipla/kutu/internal/utils"
)

func TestContains(t *testing.T) {
	var sliceTest = []string{"test", "abc", "def"}
	contained := "test"
	notContained := "xyz"

	got := utils.Contains(sliceTest, contained)
	if !got {
		t.Errorf("%v contains %s should return true, got %t", sliceTest, contained, got)
	}

	got = utils.Contains(sliceTest, notContained)
	if got {
		t.Errorf("%v contains %s should return true, got %t", sliceTest, notContained, got)
	}
}
