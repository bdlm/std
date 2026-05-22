package sorter

// SortFlag provides a type for sort flags
type SortFlag uint

const (
	// SortByKey - sort hash data by key.
	SortByKey SortFlag = 1 << iota

	// SortByValue - sort data by value using type-stratified comparison:
	// nil < bool < numeric < string < other.
	SortByValue

	// SortAsc - sort data in ascending order. This is the default.
	SortAsc

	// SortDesc - sort data in descending order.
	SortDesc

	// SortAsString - when sorting data use string comparison via cast. This is
	// a tiebreak for SortByValue and the default for SortByKey.
	SortAsString

	// SortReverse - reverse the order of the data set after sorting.
	SortReverse
)

// Sorter describes a sorter.
type Sorter interface {
	// Reverse reverses the order of the data set.
	Reverse() error
	// Sort sorts the model data.
	Sort(SortFlag) error
	// Len returns the number of items stored in this model.
	Len() int
}
