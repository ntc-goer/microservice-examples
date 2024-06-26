// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ntc-goer/microservice-examples/registry/sagaorchestation/ent/sagalogs"
)

// SagaLogsCreate is the builder for creating a SagaLogs entity.
type SagaLogsCreate struct {
	config
	mutation *SagaLogsMutation
	hooks    []Hook
}

// SetWorkflowID sets the "workflow_id" field.
func (slc *SagaLogsCreate) SetWorkflowID(s string) *SagaLogsCreate {
	slc.mutation.SetWorkflowID(s)
	return slc
}

// SetRequestID sets the "request_id" field.
func (slc *SagaLogsCreate) SetRequestID(s string) *SagaLogsCreate {
	slc.mutation.SetRequestID(s)
	return slc
}

// SetWorkflowName sets the "workflow_name" field.
func (slc *SagaLogsCreate) SetWorkflowName(s string) *SagaLogsCreate {
	slc.mutation.SetWorkflowName(s)
	return slc
}

// SetStepName sets the "step_name" field.
func (slc *SagaLogsCreate) SetStepName(s string) *SagaLogsCreate {
	slc.mutation.SetStepName(s)
	return slc
}

// SetStepOrder sets the "step_order" field.
func (slc *SagaLogsCreate) SetStepOrder(i int) *SagaLogsCreate {
	slc.mutation.SetStepOrder(i)
	return slc
}

// SetStatus sets the "status" field.
func (slc *SagaLogsCreate) SetStatus(s string) *SagaLogsCreate {
	slc.mutation.SetStatus(s)
	return slc
}

// SetMessage sets the "message" field.
func (slc *SagaLogsCreate) SetMessage(s string) *SagaLogsCreate {
	slc.mutation.SetMessage(s)
	return slc
}

// SetUpdateAt sets the "update_at" field.
func (slc *SagaLogsCreate) SetUpdateAt(t time.Time) *SagaLogsCreate {
	slc.mutation.SetUpdateAt(t)
	return slc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (slc *SagaLogsCreate) SetNillableUpdateAt(t *time.Time) *SagaLogsCreate {
	if t != nil {
		slc.SetUpdateAt(*t)
	}
	return slc
}

// SetCreatedAt sets the "created_at" field.
func (slc *SagaLogsCreate) SetCreatedAt(t time.Time) *SagaLogsCreate {
	slc.mutation.SetCreatedAt(t)
	return slc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (slc *SagaLogsCreate) SetNillableCreatedAt(t *time.Time) *SagaLogsCreate {
	if t != nil {
		slc.SetCreatedAt(*t)
	}
	return slc
}

// SetID sets the "id" field.
func (slc *SagaLogsCreate) SetID(i int) *SagaLogsCreate {
	slc.mutation.SetID(i)
	return slc
}

// Mutation returns the SagaLogsMutation object of the builder.
func (slc *SagaLogsCreate) Mutation() *SagaLogsMutation {
	return slc.mutation
}

// Save creates the SagaLogs in the database.
func (slc *SagaLogsCreate) Save(ctx context.Context) (*SagaLogs, error) {
	slc.defaults()
	return withHooks(ctx, slc.sqlSave, slc.mutation, slc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (slc *SagaLogsCreate) SaveX(ctx context.Context) *SagaLogs {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *SagaLogsCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *SagaLogsCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slc *SagaLogsCreate) defaults() {
	if _, ok := slc.mutation.UpdateAt(); !ok {
		v := sagalogs.DefaultUpdateAt()
		slc.mutation.SetUpdateAt(v)
	}
	if _, ok := slc.mutation.CreatedAt(); !ok {
		v := sagalogs.DefaultCreatedAt()
		slc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (slc *SagaLogsCreate) check() error {
	if _, ok := slc.mutation.WorkflowID(); !ok {
		return &ValidationError{Name: "workflow_id", err: errors.New(`ent: missing required field "SagaLogs.workflow_id"`)}
	}
	if v, ok := slc.mutation.WorkflowID(); ok {
		if err := sagalogs.WorkflowIDValidator(v); err != nil {
			return &ValidationError{Name: "workflow_id", err: fmt.Errorf(`ent: validator failed for field "SagaLogs.workflow_id": %w`, err)}
		}
	}
	if _, ok := slc.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "SagaLogs.request_id"`)}
	}
	if v, ok := slc.mutation.RequestID(); ok {
		if err := sagalogs.RequestIDValidator(v); err != nil {
			return &ValidationError{Name: "request_id", err: fmt.Errorf(`ent: validator failed for field "SagaLogs.request_id": %w`, err)}
		}
	}
	if _, ok := slc.mutation.WorkflowName(); !ok {
		return &ValidationError{Name: "workflow_name", err: errors.New(`ent: missing required field "SagaLogs.workflow_name"`)}
	}
	if v, ok := slc.mutation.WorkflowName(); ok {
		if err := sagalogs.WorkflowNameValidator(v); err != nil {
			return &ValidationError{Name: "workflow_name", err: fmt.Errorf(`ent: validator failed for field "SagaLogs.workflow_name": %w`, err)}
		}
	}
	if _, ok := slc.mutation.StepName(); !ok {
		return &ValidationError{Name: "step_name", err: errors.New(`ent: missing required field "SagaLogs.step_name"`)}
	}
	if v, ok := slc.mutation.StepName(); ok {
		if err := sagalogs.StepNameValidator(v); err != nil {
			return &ValidationError{Name: "step_name", err: fmt.Errorf(`ent: validator failed for field "SagaLogs.step_name": %w`, err)}
		}
	}
	if _, ok := slc.mutation.StepOrder(); !ok {
		return &ValidationError{Name: "step_order", err: errors.New(`ent: missing required field "SagaLogs.step_order"`)}
	}
	if _, ok := slc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "SagaLogs.status"`)}
	}
	if v, ok := slc.mutation.Status(); ok {
		if err := sagalogs.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SagaLogs.status": %w`, err)}
		}
	}
	if _, ok := slc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "SagaLogs.message"`)}
	}
	if v, ok := slc.mutation.Message(); ok {
		if err := sagalogs.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "SagaLogs.message": %w`, err)}
		}
	}
	if _, ok := slc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "SagaLogs.update_at"`)}
	}
	if _, ok := slc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SagaLogs.created_at"`)}
	}
	return nil
}

func (slc *SagaLogsCreate) sqlSave(ctx context.Context) (*SagaLogs, error) {
	if err := slc.check(); err != nil {
		return nil, err
	}
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	slc.mutation.id = &_node.ID
	slc.mutation.done = true
	return _node, nil
}

func (slc *SagaLogsCreate) createSpec() (*SagaLogs, *sqlgraph.CreateSpec) {
	var (
		_node = &SagaLogs{config: slc.config}
		_spec = sqlgraph.NewCreateSpec(sagalogs.Table, sqlgraph.NewFieldSpec(sagalogs.FieldID, field.TypeInt))
	)
	if id, ok := slc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := slc.mutation.WorkflowID(); ok {
		_spec.SetField(sagalogs.FieldWorkflowID, field.TypeString, value)
		_node.WorkflowID = value
	}
	if value, ok := slc.mutation.RequestID(); ok {
		_spec.SetField(sagalogs.FieldRequestID, field.TypeString, value)
		_node.RequestID = value
	}
	if value, ok := slc.mutation.WorkflowName(); ok {
		_spec.SetField(sagalogs.FieldWorkflowName, field.TypeString, value)
		_node.WorkflowName = value
	}
	if value, ok := slc.mutation.StepName(); ok {
		_spec.SetField(sagalogs.FieldStepName, field.TypeString, value)
		_node.StepName = value
	}
	if value, ok := slc.mutation.StepOrder(); ok {
		_spec.SetField(sagalogs.FieldStepOrder, field.TypeInt, value)
		_node.StepOrder = value
	}
	if value, ok := slc.mutation.Status(); ok {
		_spec.SetField(sagalogs.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := slc.mutation.Message(); ok {
		_spec.SetField(sagalogs.FieldMessage, field.TypeString, value)
		_node.Message = value
	}
	if value, ok := slc.mutation.UpdateAt(); ok {
		_spec.SetField(sagalogs.FieldUpdateAt, field.TypeTime, value)
		_node.UpdateAt = value
	}
	if value, ok := slc.mutation.CreatedAt(); ok {
		_spec.SetField(sagalogs.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// SagaLogsCreateBulk is the builder for creating many SagaLogs entities in bulk.
type SagaLogsCreateBulk struct {
	config
	err      error
	builders []*SagaLogsCreate
}

// Save creates the SagaLogs entities in the database.
func (slcb *SagaLogsCreateBulk) Save(ctx context.Context) ([]*SagaLogs, error) {
	if slcb.err != nil {
		return nil, slcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*SagaLogs, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SagaLogsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *SagaLogsCreateBulk) SaveX(ctx context.Context) []*SagaLogs {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *SagaLogsCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *SagaLogsCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}
