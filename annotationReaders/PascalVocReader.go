package annotationReaders;

import(
	"sync";
	"os";
	"github.com/lootag/ImageAuGomentationCLI/entities";
	"github.com/lootag/ImageAuGomentationCLI/annotationDtos";
	"io/ioutil";
	"encoding/xml";
	
)

type PascalVocReader struct{

}

func (pascalVocReader PascalVocReader) Read(fileName string, 
	inputAnnotations chan entities.Annotation, 
	aumentationWaitGroup *sync.WaitGroup){
	defer (*aumentationWaitGroup).Done();
	var xmlAnnotation annotationDtos.PascalVoc;
	fullPath :=  fileName;
	xmlFile, err := os.Open(fullPath);
	if err != nil{
		panic(err);
	}
	defer xmlFile.Close();
	byteValue, _ := ioutil.ReadAll(xmlFile);
	xml.Unmarshal(byteValue, &xmlAnnotation);
	var annotation entities.Annotation;
	annotation.Classes = []string{};
	annotation.BoundingBoxes = []entities.BoundingBox{};
	for objectIndex:= 0; objectIndex < len(xmlAnnotation.Objects); objectIndex++{
		var boundingBox entities.BoundingBox;
		boundingBox.Xmin = xmlAnnotation.Objects[objectIndex].BndBox.Xmin;
		boundingBox.Ymin = xmlAnnotation.Objects[objectIndex].BndBox.Ymin;
		boundingBox.Xmax = xmlAnnotation.Objects[objectIndex].BndBox.Xmax;
		boundingBox.Ymax = xmlAnnotation.Objects[objectIndex].BndBox.Ymax;
		annotation.Classes = append(annotation.Classes, xmlAnnotation.Objects[objectIndex].Name);
		annotation.BoundingBoxes = append(annotation.BoundingBoxes, boundingBox);
	}
	annotation.FileName = xmlAnnotation.FileName;
	annotation.Width = xmlAnnotation.Size.Width;
	annotation.Height = xmlAnnotation.Size.Height;
	annotation.Depth = xmlAnnotation.Size.Depth;
	inputAnnotations <- annotation;
}