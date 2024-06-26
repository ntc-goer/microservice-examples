// Code generated by ent, DO NOT EDIT.

package dish

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the dish type in the database.
	Label = "dish"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldDishID holds the string denoting the dish_id field in the database.
	FieldDishID = "dish_id"
	// FieldDishName holds the string denoting the dish_name field in the database.
	FieldDishName = "dish_name"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the dish in the database.
	Table = "dishes"
)

// Columns holds all SQL columns for dish fields.
var Columns = []string{
	FieldID,
	FieldOrderID,
	FieldDishID,
	FieldDishName,
	FieldQuantity,
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
	// DishIDValidator is a validator for the "dish_id" field. It is called by the builders before save.
	DishIDValidator func(string) error
	// DishNameValidator is a validator for the "dish_name" field. It is called by the builders before save.
	DishNameValidator func(string) error
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Dish queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByDishID orders the results by the dish_id field.
func ByDishID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDishID, opts...).ToFunc()
}

// ByDishName orders the results by the dish_name field.
func ByDishName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDishName, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByUpdateAt orders the results by the update_at field.
func ByUpdateAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateAt, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
