package sorter

// SortFlag provides a type for sort flags
type SortFlag uint

const (
	// SortByValue - sort data by value using type-stratified comparison:
	// nil < bool < numeric < string < other. This is the default/zero value.
	SortByValue SortFlag = 0

	// SortByKey - sort hash data by key.
	SortByKey SortFlag = 1 << 0

	// SortAsc - sort data in ascending order. This is the default.
	SortAsc SortFlag = 1 << 1

	// SortDesc - sort data in descending order.
	SortDesc SortFlag = 1 << 2

	// SortAsString - when sorting data use string comparison via cast. This is
	// a tiebreak for SortByValue and the default for SortByKey.
	SortAsString SortFlag = 1 << 3

	// SortReverse - reverse the order of the data set after sorting.
	SortReverse SortFlag = 1 << 4
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
