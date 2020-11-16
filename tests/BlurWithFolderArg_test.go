package tests

import (
	"testing"
	"os/exec"

	"github.com/lootag/ImageAuGomentationCLI/tests/helpers"
)

func TestBlurWithFolderArgPascalInOut(t *testing.T) {
	defer helpers.ExecuteSequentially()()
	
	//Arrange
	t.Logf("Executing test 3")
	annotationTypeIn := "PascalVoc"
	annotationTypeOut := "PascalVoc"
	action := "Blur"
	isFolderDefault := false
	helpers.Setup(isFolderDefault, annotationTypeIn)
	command := "augoment"
	arg0 := "-batch_size=5"
	arg1 := "-blur=3"
	arg2 := "-exclusion_threshold=1"
	arg3 := "-rotate=skip"
	arg4 := "-folder=resources/PascalVoc"

	cmd := exec.Command(command, arg0, arg1, arg2, arg3, arg4)

	//Act
	cmd.Run()

	//Assert
	helpers.AssertNumberOfAugmentedImages(t, action, annotationTypeOut)
	helpers.AssertNumberOfAugmentedAnnotations(t, action, annotationTypeOut)
	helpers.Destroy()
}
