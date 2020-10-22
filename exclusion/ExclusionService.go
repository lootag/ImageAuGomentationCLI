/*
This file is part of ImageAuGomentationCLI.

ImageAuGomentationCLI is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

ImageAuGomentationCLI is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with ImageAuGomentationCLI.  If not, see <https://www.gnu.org/licenses/>.
*/
package exclusion

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type ExclusionService struct {
}

func (exclusionService ExclusionService) GetClassesToExclude(exclusionThreshold int,
	userDefinedExclusions []string,
	imageNames []string,
	annotationType entities.AnnotationType) []string {
	var wg sync.WaitGroup
	annotationPaths := make(chan string, len(imageNames))
	annotationsToGroup := make(chan entities.Annotation, len(imageNames))
	wg.Add(2)
	go getAnnotationPathsFromImageNames(imageNames, annotationPaths, &wg)
	go readAnnotations(annotationType, &wg, annotationPaths, annotationsToGroup)
	wg.Wait()
	countMap := getCountMap(annotationsToGroup)
	classesToExclude := excludeClassesWithCountBelowThreshold(exclusionThreshold, countMap)
	for userDefinedIndex := range userDefinedExclusions {
		if !stringArrayContains(classesToExclude, userDefinedExclusions[userDefinedIndex]) {
			classesToExclude = append(classesToExclude, userDefinedExclusions[userDefinedIndex])
		}
	}
	return classesToExclude
}

func getAnnotationPathsFromImageNames(imageNames []string,
	annotationPaths chan string,
	excludeWaitGroup *sync.WaitGroup) {
	defer (*excludeWaitGroup).Done()
	var wg sync.WaitGroup
	for imageNameIndex := range imageNames {
		wg.Add(1)
		go getAnnotationPathsFromImageNamesWorker(imageNames[imageNameIndex], annotationPaths, &wg)
	}
	wg.Wait()
	close(annotationPaths)
}

func readAnnotations(annotationType entities.AnnotationType,
	excludeWaitGroup *sync.WaitGroup,
	annotationPaths chan string,
	annotationsToGroup chan entities.Annotation) {
	defer (*excludeWaitGroup).Done()
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotationPath := range annotationPaths {
		annotation := annotationReader.ReadSync(annotationPath)
		annotationsToGroup <- annotation
	}
	close(annotationsToGroup)
}

func getCountMap(annotationsToGroup chan entities.Annotation) map[string]int {
	countMap := make(map[string]int)
	for annotation := range annotationsToGroup {
		for classIndex := range annotation.Classes {
			if stringArrayContains(getMapKeys(countMap), annotation.Classes[classIndex]) {
				countMap[annotation.Classes[classIndex]] += 1
			} else {
				countMap[annotation.Classes[classIndex]] = 1
			}
		}
	}
	return countMap
}

func excludeClassesWithCountBelowThreshold(exclusionThreshold int, countMap map[string]int) []string {
	classesBelowThreshold := []string{}
	for key, value := range countMap {
		if value < exclusionThreshold {
			classesBelowThreshold = append(classesBelowThreshold, key)
		}
	}
	return classesBelowThreshold

}

func stringArrayContains(stringArray1 []string, toCheck string) bool {
	for stringIndex := range stringArray1 {
		if stringArray1[stringIndex] == toCheck {
			return true
		}
	}
	return false
}

func getMapKeys(stringIntMap map[string]int) []string {
	keys := []string{}
	for key, _ := range stringIntMap {
		keys = append(keys, key)
	}

	return keys
}
