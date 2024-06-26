// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SagaLogsColumns holds the columns for the "saga_logs" table.
	SagaLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "workflow_id", Type: field.TypeString},
		{Name: "request_id", Type: field.TypeString},
		{Name: "workflow_name", Type: field.TypeString},
		{Name: "step_name", Type: field.TypeString},
		{Name: "step_order", Type: field.TypeInt},
		{Name: "status", Type: field.TypeString},
		{Name: "message", Type: field.TypeString},
		{Name: "update_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
	}
	// SagaLogsTable holds the schema information for the "saga_logs" table.
	SagaLogsTable = &schema.Table{
		Name:       "saga_logs",
		Columns:    SagaLogsColumns,
		PrimaryKey: []*schema.Column{SagaLogsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SagaLogsTable,
	}
)

func init() {
	SagaLogsTable.Annotation = &entsql.Annotation{
		Table: "saga_logs",
	}
}
