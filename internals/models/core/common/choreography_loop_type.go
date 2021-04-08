package common


type ChoreographyLoopType int

// Declare typed constants each with type of ChoreographyLoopType
const (
    none ChoreographyLoopType = iota
    standard
	multiInstanceSequential
	multiInstanceParallel
)

// String returns the string value of the ChoreographyLoopType
func (s ChoreographyLoopType) String() string {
    strings := [...]string{"none", "Standard", "MultiInstanceSequential", "MultiInstanceParallel"}

    return strings[s]
}