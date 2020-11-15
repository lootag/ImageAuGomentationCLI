package tests

import (
	"testing"
	"os/exec"

	"github.com/lootag/ImageAuGomentationCLI/tests/helpers"

)

func TestAllAugmentationsWithFolderArgPascalInOut(t *testing.T) {
	defer helpers.ExecuteSequentially()()
	
	//Arrange
	t.Logf("executing test 2")
	annotationTypeIn := "PascalVoc"
	annotationTypeOut := "PascalVoc"
	action := "All"
	isFolderDefault := false
	helpers.Setup(isFolderDefault, annotationTypeIn)
	command := "augoment"
	arg0 := "-folder=resources/PascalVoc"
	arg1 := "-batch_size=5"
	arg2 := "-blur=3"
	arg3 := "-exclusion_threshold=1"
	cmd := exec.Command(command, arg0, arg1, arg2, arg3)

	//Act
	cmd.Run()

	//Assert
	helpers.AssertNumberOfAugmentedImages(t, action, annotationTypeOut)
	helpers.AssertNumberOfAugmentedAnnotations(t, action, annotationTypeOut)
	helpers.Destroy()
}
