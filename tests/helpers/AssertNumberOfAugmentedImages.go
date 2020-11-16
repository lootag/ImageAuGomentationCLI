package helpers

import (
	"testing"
	"io/ioutil"
	"strconv"
	"path/filepath"
)

func AssertNumberOfAugmentedImages(t *testing.T, action string, annotationType string) {
	targetTestFolder := filepath.Join("resources", annotationType, "TestAugmentedImages" + action)
	expectedAugmentedImages, err := ioutil.ReadDir(targetTestFolder)
	if err != nil {
		panic(err)
	}
	actualAugmentedImages, err := ioutil.ReadDir("AugmentedImages")
	if err != nil {
		panic(err)
	}

	if len(expectedAugmentedImages) != len(actualAugmentedImages) {
		t.Errorf("Was expecting " + strconv.Itoa(len(expectedAugmentedImages)) + " augmented images but got " + strconv.Itoa(len(actualAugmentedImages)))
	}

}