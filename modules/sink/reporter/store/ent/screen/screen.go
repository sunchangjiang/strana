// Code generated by entc, DO NOT EDIT.

package screen

const (
	// Label holds the string label denoting the screen type in the database.
	Label = "screen"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"

	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"

	// Table holds the table name of the screen in the database.
	Table = "screens"
	// EventsTable is the table the holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "event_screen"
)

// Columns holds all SQL columns for screen fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCategory,
}