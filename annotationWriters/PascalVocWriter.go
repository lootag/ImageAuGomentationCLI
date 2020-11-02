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
package annotationWriters

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sync"
	"strings"

	"github.com/lootag/ImageAuGomentationCLI/annotationDtos"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type PascalVocWriter struct {
}

func (pascalVocWriter PascalVocWriter) Write(annotation entities.Annotation,
	augmentationWaitGroup *sync.WaitGroup) {
	defer (*augmentationWaitGroup).Done()
	var xmlAnnotation annotationDtos.PascalVoc
	xmlAnnotation.Folder = "./AugmentedAnnotations"
	xmlAnnotation.FileName = annotation.FileName
	xmlAnnotation.Source.Database = "Unknown"
	xmlAnnotation.Objects = []annotationDtos.Object{}
	for classIndex := 0; classIndex < len(annotation.Classes); classIndex++ {
		var object annotationDtos.Object
		object.Name = strings.ToLower(annotation.Classes[classIndex])
		object.Pose = "Unspecified"
		object.Truncated = 0
		object.Difficult = 0
		object.BndBox.Xmin = annotation.BoundingBoxes[classIndex].Xmin
		object.BndBox.Ymin = annotation.BoundingBoxes[classIndex].Ymin
		object.BndBox.Xmax = annotation.BoundingBoxes[classIndex].Xmax
		object.BndBox.Ymax = annotation.BoundingBoxes[classIndex].Ymax
		xmlAnnotation.Objects = append(xmlAnnotation.Objects, object)
	}
	xmlAnnotation.Size.Width = annotation.Width
	xmlAnnotation.Size.Height = annotation.Height
	xmlAnnotation.Size.Depth = annotation.Depth
	serializedAnnotation, _ := xml.MarshalIndent(&xmlAnnotation, "", "	")
	permissions := 0644
	err := ioutil.WriteFile("./AugmentedAnnotations/"+annotation.NewName,
		serializedAnnotation,
		os.FileMode(permissions))
	if err != nil {
		panic(err)
	}
}
