package main;

import(
	"errors";
	"github.com/lootag/ImageAuGomentationCLI/augmentation";
	"github.com/lootag/ImageAuGomentationCLI/blur";
	"github.com/lootag/ImageAuGomentationCLI/flip";
	"github.com/lootag/ImageAuGomentationCLI/rotate";

)

func AugmentationsFactory(augmentation string) (augmentation.IAugmentationService, error){
	if augmentation == "blur" {
		return blur.BlurService{}, nil
	} else if augmentation == "flip" {
		return flip.FlipService{}, nil
	} else if augmentation == "rotate" {
		return rotate.RotateService{}, nil
	}else {
		return nil, errors.New("augmentation not supported!")
	}
	
}
