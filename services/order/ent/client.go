// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/dish"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/order"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Dish is the client for interacting with the Dish builders.
	Dish *DishClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Dish = NewDishClient(c.config)
	c.Order = NewOrderClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Dish:   NewDishClient(cfg),
		Order:  NewOrderClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Dish:   NewDishClient(cfg),
		Order:  NewOrderClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Dish.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Dish.Use(hooks...)
	c.Order.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Dish.Intercept(interceptors...)
	c.Order.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DishMutation:
		return c.Dish.mutate(ctx, m)
	case *OrderMutation:
		return c.Order.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DishClient is a client for the Dish schema.
type DishClient struct {
	config
}

// NewDishClient returns a client for the Dish from the given config.
func NewDishClient(c config) *DishClient {
	return &DishClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dish.Hooks(f(g(h())))`.
func (c *DishClient) Use(hooks ...Hook) {
	c.hooks.Dish = append(c.hooks.Dish, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `dish.Intercept(f(g(h())))`.
func (c *DishClient) Intercept(interceptors ...Interceptor) {
	c.inters.Dish = append(c.inters.Dish, interceptors...)
}

// Create returns a builder for creating a Dish entity.
func (c *DishClient) Create() *DishCreate {
	mutation := newDishMutation(c.config, OpCreate)
	return &DishCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Dish entities.
func (c *DishClient) CreateBulk(builders ...*DishCreate) *DishCreateBulk {
	return &DishCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DishClient) MapCreateBulk(slice any, setFunc func(*DishCreate, int)) *DishCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DishCreateBulk{err: fmt.Errorf("calling to DishClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DishCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DishCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Dish.
func (c *DishClient) Update() *DishUpdate {
	mutation := newDishMutation(c.config, OpUpdate)
	return &DishUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DishClient) UpdateOne(d *Dish) *DishUpdateOne {
	mutation := newDishMutation(c.config, OpUpdateOne, withDish(d))
	return &DishUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DishClient) UpdateOneID(id uuid.UUID) *DishUpdateOne {
	mutation := newDishMutation(c.config, OpUpdateOne, withDishID(id))
	return &DishUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Dish.
func (c *DishClient) Delete() *DishDelete {
	mutation := newDishMutation(c.config, OpDelete)
	return &DishDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DishClient) DeleteOne(d *Dish) *DishDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DishClient) DeleteOneID(id uuid.UUID) *DishDeleteOne {
	builder := c.Delete().Where(dish.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DishDeleteOne{builder}
}

// Query returns a query builder for Dish.
func (c *DishClient) Query() *DishQuery {
	return &DishQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDish},
		inters: c.Interceptors(),
	}
}

// Get returns a Dish entity by its id.
func (c *DishClient) Get(ctx context.Context, id uuid.UUID) (*Dish, error) {
	return c.Query().Where(dish.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DishClient) GetX(ctx context.Context, id uuid.UUID) *Dish {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DishClient) Hooks() []Hook {
	return c.hooks.Dish
}

// Interceptors returns the client interceptors.
func (c *DishClient) Interceptors() []Interceptor {
	return c.inters.Dish
}

func (c *DishClient) mutate(ctx context.Context, m *DishMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DishCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DishUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DishUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DishDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Dish mutation op: %q", m.Op())
	}
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `order.Intercept(f(g(h())))`.
func (c *OrderClient) Intercept(interceptors ...Interceptor) {
	c.inters.Order = append(c.inters.Order, interceptors...)
}

// Create returns a builder for creating a Order entity.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderClient) MapCreateBulk(slice any, setFunc func(*OrderCreate, int)) *OrderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderCreateBulk{err: fmt.Errorf("calling to OrderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id uuid.UUID) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderClient) DeleteOneID(id uuid.UUID) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrder},
		inters: c.Interceptors(),
	}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id uuid.UUID) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id uuid.UUID) *Order {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// Interceptors returns the client interceptors.
func (c *OrderClient) Interceptors() []Interceptor {
	return c.inters.Order
}

func (c *OrderClient) mutate(ctx context.Context, m *OrderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Order mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Dish, Order []ent.Hook
	}
	inters struct {
		Dish, Order []ent.Interceptor
	}
)
