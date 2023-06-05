// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"healthyshopper/ent/migrate"

	"healthyshopper/ent/address"
	"healthyshopper/ent/product"
	"healthyshopper/ent/user"
	"healthyshopper/ent/useraddress"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Address is the client for interacting with the Address builders.
	Address *AddressClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserAddress is the client for interacting with the UserAddress builders.
	UserAddress *UserAddressClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Address = NewAddressClient(c.config)
	c.Product = NewProductClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserAddress = NewUserAddressClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Address:     NewAddressClient(cfg),
		Product:     NewProductClient(cfg),
		User:        NewUserClient(cfg),
		UserAddress: NewUserAddressClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Address:     NewAddressClient(cfg),
		Product:     NewProductClient(cfg),
		User:        NewUserClient(cfg),
		UserAddress: NewUserAddressClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Address.
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
	c.Address.Use(hooks...)
	c.Product.Use(hooks...)
	c.User.Use(hooks...)
	c.UserAddress.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Address.Intercept(interceptors...)
	c.Product.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.UserAddress.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AddressMutation:
		return c.Address.mutate(ctx, m)
	case *ProductMutation:
		return c.Product.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *UserAddressMutation:
		return c.UserAddress.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AddressClient is a client for the Address schema.
type AddressClient struct {
	config
}

// NewAddressClient returns a client for the Address from the given config.
func NewAddressClient(c config) *AddressClient {
	return &AddressClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `address.Hooks(f(g(h())))`.
func (c *AddressClient) Use(hooks ...Hook) {
	c.hooks.Address = append(c.hooks.Address, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `address.Intercept(f(g(h())))`.
func (c *AddressClient) Intercept(interceptors ...Interceptor) {
	c.inters.Address = append(c.inters.Address, interceptors...)
}

// Create returns a builder for creating a Address entity.
func (c *AddressClient) Create() *AddressCreate {
	mutation := newAddressMutation(c.config, OpCreate)
	return &AddressCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Address entities.
func (c *AddressClient) CreateBulk(builders ...*AddressCreate) *AddressCreateBulk {
	return &AddressCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Address.
func (c *AddressClient) Update() *AddressUpdate {
	mutation := newAddressMutation(c.config, OpUpdate)
	return &AddressUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AddressClient) UpdateOne(a *Address) *AddressUpdateOne {
	mutation := newAddressMutation(c.config, OpUpdateOne, withAddress(a))
	return &AddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AddressClient) UpdateOneID(id int) *AddressUpdateOne {
	mutation := newAddressMutation(c.config, OpUpdateOne, withAddressID(id))
	return &AddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Address.
func (c *AddressClient) Delete() *AddressDelete {
	mutation := newAddressMutation(c.config, OpDelete)
	return &AddressDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AddressClient) DeleteOne(a *Address) *AddressDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AddressClient) DeleteOneID(id int) *AddressDeleteOne {
	builder := c.Delete().Where(address.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AddressDeleteOne{builder}
}

// Query returns a query builder for Address.
func (c *AddressClient) Query() *AddressQuery {
	return &AddressQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAddress},
		inters: c.Interceptors(),
	}
}

// Get returns a Address entity by its id.
func (c *AddressClient) Get(ctx context.Context, id int) (*Address, error) {
	return c.Query().Where(address.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AddressClient) GetX(ctx context.Context, id int) *Address {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserAddress queries the user_address edge of a Address.
func (c *AddressClient) QueryUserAddress(a *Address) *UserAddressQuery {
	query := (&UserAddressClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(address.Table, address.FieldID, id),
			sqlgraph.To(useraddress.Table, useraddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, address.UserAddressTable, address.UserAddressColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AddressClient) Hooks() []Hook {
	return c.hooks.Address
}

// Interceptors returns the client interceptors.
func (c *AddressClient) Interceptors() []Interceptor {
	return c.inters.Address
}

func (c *AddressClient) mutate(ctx context.Context, m *AddressMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AddressCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AddressUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AddressDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Address mutation op: %q", m.Op())
	}
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `product.Intercept(f(g(h())))`.
func (c *ProductClient) Intercept(interceptors ...Interceptor) {
	c.inters.Product = append(c.inters.Product, interceptors...)
}

// Create returns a builder for creating a Product entity.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id int) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProductClient) DeleteOneID(id int) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProduct},
		inters: c.Interceptors(),
	}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id int) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id int) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// Interceptors returns the client interceptors.
func (c *ProductClient) Interceptors() []Interceptor {
	return c.inters.Product
}

func (c *ProductClient) mutate(ctx context.Context, m *ProductMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProductCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProductDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Product mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserAddress queries the user_address edge of a User.
func (c *UserClient) QueryUserAddress(u *User) *UserAddressQuery {
	query := (&UserAddressClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(useraddress.Table, useraddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UserAddressTable, user.UserAddressColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// UserAddressClient is a client for the UserAddress schema.
type UserAddressClient struct {
	config
}

// NewUserAddressClient returns a client for the UserAddress from the given config.
func NewUserAddressClient(c config) *UserAddressClient {
	return &UserAddressClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `useraddress.Hooks(f(g(h())))`.
func (c *UserAddressClient) Use(hooks ...Hook) {
	c.hooks.UserAddress = append(c.hooks.UserAddress, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `useraddress.Intercept(f(g(h())))`.
func (c *UserAddressClient) Intercept(interceptors ...Interceptor) {
	c.inters.UserAddress = append(c.inters.UserAddress, interceptors...)
}

// Create returns a builder for creating a UserAddress entity.
func (c *UserAddressClient) Create() *UserAddressCreate {
	mutation := newUserAddressMutation(c.config, OpCreate)
	return &UserAddressCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserAddress entities.
func (c *UserAddressClient) CreateBulk(builders ...*UserAddressCreate) *UserAddressCreateBulk {
	return &UserAddressCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserAddress.
func (c *UserAddressClient) Update() *UserAddressUpdate {
	mutation := newUserAddressMutation(c.config, OpUpdate)
	return &UserAddressUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserAddressClient) UpdateOne(ua *UserAddress) *UserAddressUpdateOne {
	mutation := newUserAddressMutation(c.config, OpUpdateOne, withUserAddress(ua))
	return &UserAddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserAddressClient) UpdateOneID(id int) *UserAddressUpdateOne {
	mutation := newUserAddressMutation(c.config, OpUpdateOne, withUserAddressID(id))
	return &UserAddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserAddress.
func (c *UserAddressClient) Delete() *UserAddressDelete {
	mutation := newUserAddressMutation(c.config, OpDelete)
	return &UserAddressDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserAddressClient) DeleteOne(ua *UserAddress) *UserAddressDeleteOne {
	return c.DeleteOneID(ua.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserAddressClient) DeleteOneID(id int) *UserAddressDeleteOne {
	builder := c.Delete().Where(useraddress.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserAddressDeleteOne{builder}
}

// Query returns a query builder for UserAddress.
func (c *UserAddressClient) Query() *UserAddressQuery {
	return &UserAddressQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUserAddress},
		inters: c.Interceptors(),
	}
}

// Get returns a UserAddress entity by its id.
func (c *UserAddressClient) Get(ctx context.Context, id int) (*UserAddress, error) {
	return c.Query().Where(useraddress.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserAddressClient) GetX(ctx context.Context, id int) *UserAddress {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UserAddress.
func (c *UserAddressClient) QueryUser(ua *UserAddress) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ua.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(useraddress.Table, useraddress.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, useraddress.UserTable, useraddress.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ua.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAddress queries the address edge of a UserAddress.
func (c *UserAddressClient) QueryAddress(ua *UserAddress) *AddressQuery {
	query := (&AddressClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ua.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(useraddress.Table, useraddress.FieldID, id),
			sqlgraph.To(address.Table, address.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, useraddress.AddressTable, useraddress.AddressColumn),
		)
		fromV = sqlgraph.Neighbors(ua.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserAddressClient) Hooks() []Hook {
	return c.hooks.UserAddress
}

// Interceptors returns the client interceptors.
func (c *UserAddressClient) Interceptors() []Interceptor {
	return c.inters.UserAddress
}

func (c *UserAddressClient) mutate(ctx context.Context, m *UserAddressMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserAddressCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserAddressUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserAddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserAddressDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UserAddress mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Address, Product, User, UserAddress []ent.Hook
	}
	inters struct {
		Address, Product, User, UserAddress []ent.Interceptor
	}
)
