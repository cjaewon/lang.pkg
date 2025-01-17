// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"lang.pkg/ent/migrate"

	"lang.pkg/ent/book"
	"lang.pkg/ent/user"
	"lang.pkg/ent/voca"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Book is the client for interacting with the Book builders.
	Book *BookClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Voca is the client for interacting with the Voca builders.
	Voca *VocaClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Book = NewBookClient(c.config)
	c.User = NewUserClient(c.config)
	c.Voca = NewVocaClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Book:   NewBookClient(cfg),
		User:   NewUserClient(cfg),
		Voca:   NewVocaClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Book:   NewBookClient(cfg),
		User:   NewUserClient(cfg),
		Voca:   NewVocaClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Book.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Book.Use(hooks...)
	c.User.Use(hooks...)
	c.Voca.Use(hooks...)
}

// BookClient is a client for the Book schema.
type BookClient struct {
	config
}

// NewBookClient returns a client for the Book from the given config.
func NewBookClient(c config) *BookClient {
	return &BookClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `book.Hooks(f(g(h())))`.
func (c *BookClient) Use(hooks ...Hook) {
	c.hooks.Book = append(c.hooks.Book, hooks...)
}

// Create returns a create builder for Book.
func (c *BookClient) Create() *BookCreate {
	mutation := newBookMutation(c.config, OpCreate)
	return &BookCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Book entities.
func (c *BookClient) CreateBulk(builders ...*BookCreate) *BookCreateBulk {
	return &BookCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Book.
func (c *BookClient) Update() *BookUpdate {
	mutation := newBookMutation(c.config, OpUpdate)
	return &BookUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookClient) UpdateOne(b *Book) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBook(b))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookClient) UpdateOneID(id int) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBookID(id))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Book.
func (c *BookClient) Delete() *BookDelete {
	mutation := newBookMutation(c.config, OpDelete)
	return &BookDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BookClient) DeleteOne(b *Book) *BookDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BookClient) DeleteOneID(id int) *BookDeleteOne {
	builder := c.Delete().Where(book.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookDeleteOne{builder}
}

// Query returns a query builder for Book.
func (c *BookClient) Query() *BookQuery {
	return &BookQuery{config: c.config}
}

// Get returns a Book entity by its id.
func (c *BookClient) Get(ctx context.Context, id int) (*Book, error) {
	return c.Query().Where(book.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookClient) GetX(ctx context.Context, id int) *Book {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryVocas queries the vocas edge of a Book.
func (c *BookClient) QueryVocas(b *Book) *VocaQuery {
	query := &VocaQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(voca.Table, voca.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, book.VocasTable, book.VocasColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Book.
func (c *BookClient) QueryOwner(b *Book) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, book.OwnerTable, book.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookClient) Hooks() []Hook {
	return c.hooks.Book
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

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of User entities.
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBooks queries the books edge of a User.
func (c *UserClient) QueryBooks(u *User) *BookQuery {
	query := &BookQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(book.Table, book.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BooksTable, user.BooksColumn),
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

// VocaClient is a client for the Voca schema.
type VocaClient struct {
	config
}

// NewVocaClient returns a client for the Voca from the given config.
func NewVocaClient(c config) *VocaClient {
	return &VocaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `voca.Hooks(f(g(h())))`.
func (c *VocaClient) Use(hooks ...Hook) {
	c.hooks.Voca = append(c.hooks.Voca, hooks...)
}

// Create returns a create builder for Voca.
func (c *VocaClient) Create() *VocaCreate {
	mutation := newVocaMutation(c.config, OpCreate)
	return &VocaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Voca entities.
func (c *VocaClient) CreateBulk(builders ...*VocaCreate) *VocaCreateBulk {
	return &VocaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Voca.
func (c *VocaClient) Update() *VocaUpdate {
	mutation := newVocaMutation(c.config, OpUpdate)
	return &VocaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VocaClient) UpdateOne(v *Voca) *VocaUpdateOne {
	mutation := newVocaMutation(c.config, OpUpdateOne, withVoca(v))
	return &VocaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VocaClient) UpdateOneID(id uuid.UUID) *VocaUpdateOne {
	mutation := newVocaMutation(c.config, OpUpdateOne, withVocaID(id))
	return &VocaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Voca.
func (c *VocaClient) Delete() *VocaDelete {
	mutation := newVocaMutation(c.config, OpDelete)
	return &VocaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VocaClient) DeleteOne(v *Voca) *VocaDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VocaClient) DeleteOneID(id uuid.UUID) *VocaDeleteOne {
	builder := c.Delete().Where(voca.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VocaDeleteOne{builder}
}

// Query returns a query builder for Voca.
func (c *VocaClient) Query() *VocaQuery {
	return &VocaQuery{config: c.config}
}

// Get returns a Voca entity by its id.
func (c *VocaClient) Get(ctx context.Context, id uuid.UUID) (*Voca, error) {
	return c.Query().Where(voca.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VocaClient) GetX(ctx context.Context, id uuid.UUID) *Voca {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VocaClient) Hooks() []Hook {
	return c.hooks.Voca
}
