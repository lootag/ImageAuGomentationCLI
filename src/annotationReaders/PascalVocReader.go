package annotationReaders

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sync"
	"path"
	
	"github.com/lootag/ImageAuGomentationCLI/commons"
	"github.com/lootag/ImageAuGomentationCLI/annotationDtos"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type PascalVocReader struct {
}

//In the Read function I set the "FileName" by composing the image's actual name retrieved from fs, because the xml's FileName value could be different from the file's name
func (this PascalVocReader) Read(annotationPath string,
	inputAnnotations chan entities.Annotation,
	aumentationWaitGroup *sync.WaitGroup) {
	defer (*aumentationWaitGroup).Done()
	var xmlAnnotation annotationDtos.PascalVoc
	xmlFile, err := os.Open(annotationPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &xmlAnnotation)
	var annotation entities.Annotation
	annotation.Classes = []string{}
	annotation.BoundingBoxes = []entities.BoundingBox{}
	for objectIndex := 0; objectIndex < len(xmlAnnotation.Objects); objectIndex++ {
		var boundingBox entities.BoundingBox
		boundingBox.Xmin = xmlAnnotation.Objects[objectIndex].BndBox.Xmin
		boundingBox.Ymin = xmlAnnotation.Objects[objectIndex].BndBox.Ymin
		boundingBox.Xmax = xmlAnnotation.Objects[objectIndex].BndBox.Xmax
		boundingBox.Ymax = xmlAnnotation.Objects[objectIndex].BndBox.Ymax
		annotation.Classes = append(annotation.Classes, xmlAnnotation.Objects[objectIndex].Name)
		annotation.BoundingBoxes = append(annotation.BoundingBoxes, boundingBox)
	}
	_, annotationName := path.Split(annotationPath)
	annotation.FileName = annotationName[:len(annotationName)-3] + commons.GetFileExtension(xmlAnnotation.FileName)
	annotation.Width = xmlAnnotation.Size.Width
	annotation.Height = xmlAnnotation.Size.Height
	annotation.Depth = xmlAnnotation.Size.Depth
	inputAnnotations <- annotation
}

func (this PascalVocReader) ReadSync(annotationPath string) entities.Annotation {
	var xmlAnnotation annotationDtos.PascalVoc
	xmlFile, err := os.Open(annotationPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &xmlAnnotation)
	var annotation entities.Annotation
	annotation.Classes = []string{}
	annotation.BoundingBoxes = []entities.BoundingBox{}
	for objectIndex := 0; objectIndex < len(xmlAnnotation.Objects); objectIndex++ {
		var boundingBox entities.BoundingBox
		boundingBox.Xmin = xmlAnnotation.Objects[objectIndex].BndBox.Xmin
		boundingBox.Ymin = xmlAnnotation.Objects[objectIndex].BndBox.Ymin
		boundingBox.Xmax = xmlAnnotation.Objects[objectIndex].BndBox.Xmax
		boundingBox.Ymax = xmlAnnotation.Objects[objectIndex].BndBox.Ymax
		annotation.Classes = append(annotation.Classes, xmlAnnotation.Objects[objectIndex].Name)
		annotation.BoundingBoxes = append(annotation.BoundingBoxes, boundingBox)
	}
	_, annotationName := path.Split(annotationPath)
	annotation.FileName = annotationName[:len(annotationName)-3] + commons.GetFileExtension(xmlAnnotation.FileName)
	annotation.Width = xmlAnnotation.Size.Width
	annotation.Height = xmlAnnotation.Size.Height
	annotation.Depth = xmlAnnotation.Size.Depth
	return annotation
}
