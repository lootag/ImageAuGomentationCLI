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
package scan

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/annotationReaders"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type ScanningService struct {
}

func (scanningService ScanningService) Scan(annotationType entities.AnnotationType,
	folderToScan string) {
	annotationsToRead := getAnnotationPaths(folderToScan)
	annotationsToGroup := readAnnotations(annotationType, annotationsToRead)
	countMap := getCountMap(annotationsToGroup)
	fmt.Println("Here's a scan of your data: ")
	for key, value := range countMap {
		fmt.Println(key + ", " + strconv.Itoa(value) + " instances")
	}
}

func getAnnotationPaths(folderToScan string) []string {
	annotationsToRead := []string{}
	root := folderToScan + "/Annotations"
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for fileInfoIndex := range fileInfos {
		annotationsToRead = append(annotationsToRead, folderToScan+"/Annotations/"+fileInfos[fileInfoIndex].Name())
	}
	return annotationsToRead
}

func readAnnotations(annotationType entities.AnnotationType,
	annotationsToRead []string) []entities.Annotation {
	annotationsToGroup := []entities.Annotation{}
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotationPathIndex := range annotationsToRead {
		annotation := annotationReader.ReadSync(annotationsToRead[annotationPathIndex])
		annotationsToGroup = append(annotationsToGroup, annotation)
	}
	return annotationsToGroup
}

func getCountMap(annotationsToGroup []entities.Annotation) map[string]int {
	countMap := make(map[string]int)
	for annotationIndex := range annotationsToGroup {
		for classIndex := range annotationsToGroup[annotationIndex].Classes {
			if stringArrayContains(getMapKeys(countMap), annotationsToGroup[annotationIndex].Classes[classIndex]) {
				countMap[annotationsToGroup[annotationIndex].Classes[classIndex]] += 1
			} else {
				countMap[annotationsToGroup[annotationIndex].Classes[classIndex]] = 1
			}
		}
	}
	return countMap
}

func addAnnotationPathToChannel(annotationPath string,
	annotationsToRead chan string,
	getAnnotationPathsWaitGroup *sync.WaitGroup) {
	defer (*getAnnotationPathsWaitGroup).Done()
	annotationsToRead <- annotationPath
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
