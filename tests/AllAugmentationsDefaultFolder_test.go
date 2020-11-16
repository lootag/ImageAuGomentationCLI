package tests

import (
	"testing"
	"os/exec"

	"github.com/lootag/ImageAuGomentationCLI/tests/helpers"
)


func TestAllAugmentationsDefaultFolderPascalInOut(t *testing.T) {
	defer helpers.ExecuteSequentially()()
	
	//Arrange
	annotationTypeIn := "PascalVoc"
	annotationTypeOut := "PascalVoc"
	action := "All"
	isFolderDefault := true
	helpers.Setup(isFolderDefault, annotationTypeIn)
	command := "augoment"
	arg0 := "-batch_size=5"
	arg1 := "-blur=3"
	arg2 := "-exclusion_threshold=1"
	cmd := exec.Command(command, arg0, arg1, arg2)

	//Act
	cmd.Run()

	//Assert
	helpers.AssertNumberOfAugmentedImages(t, action, annotationTypeOut)
	helpers.AssertNumberOfAugmentedAnnotations(t, action, annotationTypeOut)
	helpers.Destroy()
}



