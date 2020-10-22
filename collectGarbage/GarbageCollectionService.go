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
package collectGarbage

import (
	"io/ioutil"
	"os"
)

type GarbageCollectionService struct {
}

func (garbageCollectionService GarbageCollectionService) CollectGarbage() {
	annotations := getAugmentedAnnotations()
	images := getAugmentedImages()
	for imageIndex := range images {
		annotationToCheck := images[imageIndex][:len(images[imageIndex])-3] + "xml"
		if !contains(annotations, annotationToCheck) {
			os.Remove("./AugmentedImages/" + images[imageIndex])
		}
	}

}

func getAugmentedAnnotations() []string {
	annotations := []string{}
	root := "./AugmentedAnnotations"
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for fileInfoIndex := range fileInfos {
		annotations = append(annotations, fileInfos[fileInfoIndex].Name())
	}
	return annotations
}

func getAugmentedImages() []string {
	images := []string{}
	root := "./AugmentedImages"
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for fileInfoIndex := range fileInfos {
		images = append(images, fileInfos[fileInfoIndex].Name())
	}
	return images

}

func contains(stringArray1 []string, toCheck string) bool {
	for stringIndex := range stringArray1 {
		if stringArray1[stringIndex] == toCheck {
			return true
		}
	}
	return false
}
