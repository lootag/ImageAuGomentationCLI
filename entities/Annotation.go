package entities;

type Annotation struct{
	FileName string;
	NewName string;
	Width int;
	Height int; 
	Depth int;
	BoundingBoxes[] BoundingBox;
	Classes[] string;
}