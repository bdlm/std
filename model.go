package std

/*
ModelType defines a type for use by models.
*/
type ModelType int

const (
	// ModelTypeList causes the model to behave as a list (keys are unsigned,
	// contiguous integers beginning at 0).
	ModelTypeList ModelType = iota
	// ModelTypeHash causes the model to behave as a hash (keys are strings,
	// order is static).
	ModelTypeHash
)

/*
Model is a list or a map of Values.

This interface defines data storage and access methods for data models in
order to provide a consistent interface for communicating messages between
instances. This allows several abstractions on and recursions into
multidimensional untyped data structures.
*/
type Model interface {
	// Delete removes a value from this model.
	Delete(key interface{}) error
	// Filter filters elements of the data using a callback function and
	// returns the result.
	Filter(callback func(Value) Model) Model
	// Get returns the specified data value in this model.
	Get(key interface{}) (Value, error)
	// GetID returns returns this model's id.
	GetID() interface{}
	// GetType returns the model type.
	GetType() ModelType
	// Has tests to see of a specified data element exists in this model.
	Has(key interface{}) bool
	// Lock marks this model as read-only.
	Lock()
	// Map applies a callback to all elements in this model and returns the
	// result.
	Map(callback func(Value) Model) Model
	// Merge merges data from any Model into this Model.
	Merge(Model) error
	// Push a value to the end of the internal data store.
	Push(value interface{}) error
	// Reduce iteratively reduces the data to a single value using a
	// callback function and returns the result.
	Reduce(callback func(Value) bool) Value
	// Reverse reverses the order of the data store.
	Reverse()
	// Set stores a value in the internal data store. All values must be
	// identified by key.
	Set(key interface{}, value interface{}) error
	// SetID sets this Model's identifier property.
	SetID(id interface{})
	// SetType sets the model type. If any data is stored in this model,
	// this property becomes read-only.
	SetType(typ ModelType) error
	// Sort sorts the model data.
	Sort(flags uintptr) error
}

/*
Value represents a value stored in a node in a Model.
*/
type Value interface {
	// Bool returns the boolean representation of the value of this node, or
	// an error if the type conversion is not possible.
	Bool() (bool, error)
	// Float returns the float64 representation of the value of this node,
	// or an error if the type conversion is not possible.
	Float() (float64, error)
	// Float32 returns the float32 representation of the value of this node,
	// or an error if the type conversion is not possible.
	Float32() (float32, error)
	// Float64 returns the float64 representation of the value of this node,
	// or an error if the type conversion is not possible.
	Float64() (float64, error)
	// Int returns the int representation of the value of this node, or an
	// error if the type conversion is not possible.
	Int() (int, error)
	// List returns the array of Values stored in this node, or an error if
	// the type conversion is not possible.
	List() ([]Value, error)
	// Map returns the map[string]Value data stored in this node, or an
	// error if the type conversion is not possible.
	Map() (map[string]Value, error)
	// Model returns the Model stored at this node, or an error if the value
	// does not implement Model.
	Model() (Model, error)
	// String returns the boolean representation of the value, or an error
	// if the type conversion is not possible.
	String() (string, error)
	// Value returns the untyped value.
	Value() interface{}
}
