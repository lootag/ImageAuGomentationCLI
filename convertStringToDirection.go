package main

import (
	"errors"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func convertStringToDirection(argument string) (entities.Direction, error) {
	switch argument {
	case "all":
		return entities.ALL, nil
	case "left":
		return entities.LEFT, nil
	case "right":
		return entities.RIGHT, nil
	case "flip":
		return entities.FLIP, nil
	case "skip":
		return entities.NIL_DIRECTION, nil
	}
	return entities.NIL_DIRECTION, errors.New("The rotation you've specified is not supported.")
}
