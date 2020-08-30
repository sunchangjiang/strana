// Code generated by entc, DO NOT EDIT.

package extra

const (
	// Label holds the string label denoting the extra type in the database.
	Label = "extra"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValues holds the string denoting the values field in the database.
	FieldValues = "values"

	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"

	// Table holds the table name of the extra in the database.
	Table = "extras"
	// EventsTable is the table the holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "event_extra"
)

// Columns holds all SQL columns for extra fields.
var Columns = []string{
	FieldID,
	FieldValues,
}