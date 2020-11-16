package helpers

import (
	"testing"
	"io/ioutil"
	"strconv"
	"path/filepath"
)

func AssertNumberOfAugmentedAnnotations(t *testing.T, action string, annotationType string) {
	targetTestFolder :=  filepath.Join("resources", annotationType, "TestAugmentedAnnotations" + action)
	expectedAugmentedAnnotations, err := ioutil.ReadDir(targetTestFolder)
	if err != nil {
		panic(err)
	}
	actualAugmentedAnnotations, err := ioutil.ReadDir("AugmentedAnnotations")
	if err != nil {
		panic(err)
	}
	if len(expectedAugmentedAnnotations) != len(actualAugmentedAnnotations) {
		t.Errorf("Was expecting " + strconv.Itoa(len(expectedAugmentedAnnotations)) + " augmented annotations but got " + strconv.Itoa(len(actualAugmentedAnnotations)))
	}
}