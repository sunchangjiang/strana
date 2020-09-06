// Code generated by entc, DO NOT EDIT.

package action

const (
	// Label holds the string label denoting the action type in the database.
	Label = "action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldActionLabel holds the string denoting the action_label field in the database.
	FieldActionLabel = "label"
	// FieldProperty holds the string denoting the property field in the database.
	FieldProperty = "property"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"

	// EdgeEvent holds the string denoting the event edge name in mutations.
	EdgeEvent = "event"

	// Table holds the table name of the action in the database.
	Table = "actions"
	// EventTable is the table the holds the event relation/edge.
	EventTable = "actions"
	// EventInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventInverseTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_action"
)

// Columns holds all SQL columns for action fields.
var Columns = []string{
	FieldID,
	FieldAction,
	FieldCategory,
	FieldActionLabel,
	FieldProperty,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Action type.
var ForeignKeys = []string{
	"event_action",
}
