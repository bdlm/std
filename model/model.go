package model

// ModelType defines a type for use by models.
type ModelType int

const (
	// HASH causes the model to behave as a hash (keys are strings,
	// order is static).
	HASH ModelType = iota

	// LIST causes the model to behave as a list (keys are unsigned,
	// contiguous integers beginning at 0).
	LIST
)

// Model is a list or a map of Values.
//
// This interface defines data storage and access methods for data models in
// order to provide a consistent interface for communicating messages between
// instances. This allows several abstractions on and recursions into
// multidimensional untyped data structures.
type Model interface {
	// Delete removes a value from this model.
	Delete(key interface{}) error
	// Filter filters elements of the data using a callback function and
	// returns the result.
	Filter(callback func(Value) bool) Model
	// Get returns the specified data value in this model.
	Get(key interface{}) (Value, error)
	// GetData returns snapshots of the current data set and indexes.
	//
	// Implementations must not expose internal slice or map storage directly;
	// the returned slice and maps must be safe for the caller to read and
	// modify without affecting the model's internal state.
	//
	// The return values are:
	//   - data: a snapshot of the model's values in model order.
	//   - keyToIndex: a snapshot of hash keys to indexes for HASH models; nil
	//     for LIST models.
	//   - indexToKey: a snapshot of indexes to hash keys for HASH models; nil
	//     for LIST models.
	GetData() ([]interface{}, map[string]int, map[int]string)
	// GetID returns this model's id.
	GetID() interface{}
	// GetType returns the model type.
	GetType() ModelType
	// Has tests to see of a specified data element exists in this model.
	Has(key interface{}) bool
	// Lock marks this model as read-only. There is no Unlock.
	Lock()
	// Map applies a callback to all elements in this model and returns the
	// result.
	Map(callback func(Value) Value) Model
	// Merge merges data from any Model into this Model.
	Merge(Model) error
	// Push a value to the end of the internal data store.
	Push(value interface{}) error
	// Reduce iteratively reduces the data to a single value using a callback
	// function and returns the result.
	//
	// The callback function takes two `Value` arguments, the first being the
	// carry value from the previous iteration (or the first element for the
	// first iteration) and the second being the current element. The callback
	// returns a `Value` which becomes the carry value for the next iteration.
	// After all iterations, Reduce returns the final carry value.
	//
	// For a non-empty model, the first element is used as the initial carry,
	// so the callback is first invoked with that carry and the second
	// element. For an empty model, Reduce returns nil and does not invoke the
	// callback.
	Reduce(callback func(carry, cur Value) Value) Value
	// Set stores a value in the internal data store. All values must be
	// identified by key.
	Set(key interface{}, value interface{}) error
	// SetData replaces the current data stored in the model with the
	// provided data.
	SetData(data interface{}) error
	// Reverse reverses the order of the data store.
	Reverse() error
	// SetID sets this Model's identifier property.
	//
	// It returns an error if the model is read-only (for example, after
	// Lock has been called) or if id is invalid or of an unsupported type
	// for the implementation.
	SetID(id interface{}) error
	// SetType sets the model type. If any data is stored in this model,
	// this property becomes read-only.
	SetType(typ ModelType) error
}
