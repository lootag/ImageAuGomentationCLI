package annotationDtos

import (
	"encoding/xml"
)

type PascalVoc struct {
	XMLName   xml.Name `xml:"annotation"`
	Folder    string   `xml:"folder"`
	FileName  string   `xml:"filename"`
	Path      string   `xml:"path"`
	Source    Source   `xml:"source"`
	Size      Size     `xml:"size"`
	Segmented int      `xml:"segmented"`
	Objects   []Object `xml:"object"`
}

type Size struct {
	XMLName xml.Name `xml:"size"`
	Width   int      `xml:"width"`
	Height  int      `xml:"height"`
	Depth   int      `xml:"depth"`
}

type Source struct {
	XMLName  xml.Name `xml:"source"`
	Database string   `xml:"database"`
}

type Object struct {
	XMLName   xml.Name `xml:"object"`
	Name      string   `xml:"name"`
	Pose      string   `xml:"pose"`
	Truncated int      `xml:"truncated"`
	Difficult int      `xml:"difficult"`
	BndBox    BndBox   `xml:"bndbox"`
}

type BndBox struct {
	XMLName xml.Name `xml:"bndbox"`
	Xmin    int      `xml:"xmin"`
	Ymin    int      `xml:"ymin"`
	Xmax    int      `xml:"xmax"`
	Ymax    int      `xml:"ymax"`
}
