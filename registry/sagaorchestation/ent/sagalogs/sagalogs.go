// Code generated by ent, DO NOT EDIT.

package sagalogs

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the sagalogs type in the database.
	Label = "saga_logs"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWorkflowID holds the string denoting the workflow_id field in the database.
	FieldWorkflowID = "workflow_id"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldWorkflowName holds the string denoting the workflow_name field in the database.
	FieldWorkflowName = "workflow_name"
	// FieldStepName holds the string denoting the step_name field in the database.
	FieldStepName = "step_name"
	// FieldStepOrder holds the string denoting the step_order field in the database.
	FieldStepOrder = "step_order"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the sagalogs in the database.
	Table = "saga_logs"
)

// Columns holds all SQL columns for sagalogs fields.
var Columns = []string{
	FieldID,
	FieldWorkflowID,
	FieldRequestID,
	FieldWorkflowName,
	FieldStepName,
	FieldStepOrder,
	FieldStatus,
	FieldMessage,
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
	// WorkflowIDValidator is a validator for the "workflow_id" field. It is called by the builders before save.
	WorkflowIDValidator func(string) error
	// RequestIDValidator is a validator for the "request_id" field. It is called by the builders before save.
	RequestIDValidator func(string) error
	// WorkflowNameValidator is a validator for the "workflow_name" field. It is called by the builders before save.
	WorkflowNameValidator func(string) error
	// StepNameValidator is a validator for the "step_name" field. It is called by the builders before save.
	StepNameValidator func(string) error
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
	// MessageValidator is a validator for the "message" field. It is called by the builders before save.
	MessageValidator func(string) error
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the SagaLogs queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByWorkflowID orders the results by the workflow_id field.
func ByWorkflowID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowID, opts...).ToFunc()
}

// ByRequestID orders the results by the request_id field.
func ByRequestID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequestID, opts...).ToFunc()
}

// ByWorkflowName orders the results by the workflow_name field.
func ByWorkflowName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkflowName, opts...).ToFunc()
}

// ByStepName orders the results by the step_name field.
func ByStepName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStepName, opts...).ToFunc()
}

// ByStepOrder orders the results by the step_order field.
func ByStepOrder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStepOrder, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByMessage orders the results by the message field.
func ByMessage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMessage, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
