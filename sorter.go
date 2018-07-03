package std

type SortFlag int

const (
	// SortByKey - sort hash data by key
	SortByKey SortFlag = iota
)

type Sorter interface {
	// Sort sorts the model data.
	Sort(flag SortFlag) error
}
