// Code generated by ent, DO NOT EDIT.

package ticket

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the ticket type in the database.
	Label = "ticket"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the ticket in the database.
	Table = "tickets"
)

// Columns holds all SQL columns for ticket fields.
var Columns = []string{
	FieldID,
	FieldOrderID,
	FieldRequestID,
	FieldStatus,
	FieldUpdateAt,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// OrderIDValidator is a validator for the "order_id" field. It is called by the builders before save.
	OrderIDValidator func(string) error
	// RequestIDValidator is a validator for the "request_id" field. It is called by the builders before save.
	RequestIDValidator func(string) error
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusCREATE_PENDING      Status = "CREATE_PENDING"
	StatusAWAITING_ACCEPTANCE Status = "AWAITING_ACCEPTANCE"
	StatusCANCELED            Status = "CANCELED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusCREATE_PENDING, StatusAWAITING_ACCEPTANCE, StatusCANCELED:
		return nil
	default:
		return fmt.Errorf("ticket: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Ticket queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByRequestID orders the results by the request_id field.
func ByRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
