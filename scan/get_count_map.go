package scan

import (

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

func(this ScanningService) getCountMap(annotationsToGroup []entities.Annotation) map[string]int {
	countMap := make(map[string]int)
	for annotationIndex := range annotationsToGroup {
		for classIndex := range annotationsToGroup[annotationIndex].Classes {
			if stringArrayContains(this.getMapKeys(countMap), annotationsToGroup[annotationIndex].Classes[classIndex]) {
				countMap[annotationsToGroup[annotationIndex].Classes[classIndex]] += 1
			} else {
				countMap[annotationsToGroup[annotationIndex].Classes[classIndex]] = 1
			}
		}
	}
	return countMap
}