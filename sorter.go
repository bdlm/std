package std

type SortFlag int

const (
	// SortByKey - sort hash data by key
	SortByKey SortFlag = iota
	SortStrings
)

/*
Sorter
*/
type Sorter interface {
	// Reverse reverses the order of the data set.
	Reverse(SortFlag) error
	// Sort sorts the model data.
	Sort(SortFlag) error
	// Len returns the number of items stored in this model.
	Len() int
}
