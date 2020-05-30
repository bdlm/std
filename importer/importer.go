package importer

// Importer is the interface that wraps the standard Import method.
type Importer interface {
	// Import imports the given data into the reciever's data structure. It
	// accepts the empty interface and extracts data from there, returning
	// an error if the import fails.
	//
	// Implementations must not retain d. Implementations should not retain
	// any imported data if returning an error.
	Import(d interface{}) error
}
