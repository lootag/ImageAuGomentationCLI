package annotationWriters

import (
	"encoding/xml"
	"github.com/lootag/ImageAuGomentationCLI/annotationDtos"
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"io/ioutil"
	"os"
	"sync"
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
		object.Name = annotation.Classes[classIndex]
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
