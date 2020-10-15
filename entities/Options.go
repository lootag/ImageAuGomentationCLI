package entities

type Options struct {
	Size           int
	Side           Direction
	Sigma          float64
	Xml            bool
	BatchSize      int
	InAnnotationType AnnotationType
	OutAnnotationType AnnotationType
}
