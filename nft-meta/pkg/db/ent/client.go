// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/migrate"

	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/block"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/endpoint"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/synctask"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Block is the client for interacting with the Block builders.
	Block *BlockClient
	// Contract is the client for interacting with the Contract builders.
	Contract *ContractClient
	// Endpoint is the client for interacting with the Endpoint builders.
	Endpoint *EndpointClient
	// Snapshot is the client for interacting with the Snapshot builders.
	Snapshot *SnapshotClient
	// SyncTask is the client for interacting with the SyncTask builders.
	SyncTask *SyncTaskClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// Transfer is the client for interacting with the Transfer builders.
	Transfer *TransferClient
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
	c.Block = NewBlockClient(c.config)
	c.Contract = NewContractClient(c.config)
	c.Endpoint = NewEndpointClient(c.config)
	c.Snapshot = NewSnapshotClient(c.config)
	c.SyncTask = NewSyncTaskClient(c.config)
	c.Token = NewTokenClient(c.config)
	c.Transfer = NewTransferClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		Block:    NewBlockClient(cfg),
		Contract: NewContractClient(cfg),
		Endpoint: NewEndpointClient(cfg),
		Snapshot: NewSnapshotClient(cfg),
		SyncTask: NewSyncTaskClient(cfg),
		Token:    NewTokenClient(cfg),
		Transfer: NewTransferClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		Block:    NewBlockClient(cfg),
		Contract: NewContractClient(cfg),
		Endpoint: NewEndpointClient(cfg),
		Snapshot: NewSnapshotClient(cfg),
		SyncTask: NewSyncTaskClient(cfg),
		Token:    NewTokenClient(cfg),
		Transfer: NewTransferClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Block.
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
	c.Block.Use(hooks...)
	c.Contract.Use(hooks...)
	c.Endpoint.Use(hooks...)
	c.Snapshot.Use(hooks...)
	c.SyncTask.Use(hooks...)
	c.Token.Use(hooks...)
	c.Transfer.Use(hooks...)
}

// BlockClient is a client for the Block schema.
type BlockClient struct {
	config
}

// NewBlockClient returns a client for the Block from the given config.
func NewBlockClient(c config) *BlockClient {
	return &BlockClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `block.Hooks(f(g(h())))`.
func (c *BlockClient) Use(hooks ...Hook) {
	c.hooks.Block = append(c.hooks.Block, hooks...)
}

// Create returns a builder for creating a Block entity.
func (c *BlockClient) Create() *BlockCreate {
	mutation := newBlockMutation(c.config, OpCreate)
	return &BlockCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Block entities.
func (c *BlockClient) CreateBulk(builders ...*BlockCreate) *BlockCreateBulk {
	return &BlockCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Block.
func (c *BlockClient) Update() *BlockUpdate {
	mutation := newBlockMutation(c.config, OpUpdate)
	return &BlockUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlockClient) UpdateOne(b *Block) *BlockUpdateOne {
	mutation := newBlockMutation(c.config, OpUpdateOne, withBlock(b))
	return &BlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlockClient) UpdateOneID(id uuid.UUID) *BlockUpdateOne {
	mutation := newBlockMutation(c.config, OpUpdateOne, withBlockID(id))
	return &BlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Block.
func (c *BlockClient) Delete() *BlockDelete {
	mutation := newBlockMutation(c.config, OpDelete)
	return &BlockDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlockClient) DeleteOne(b *Block) *BlockDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BlockClient) DeleteOneID(id uuid.UUID) *BlockDeleteOne {
	builder := c.Delete().Where(block.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlockDeleteOne{builder}
}

// Query returns a query builder for Block.
func (c *BlockClient) Query() *BlockQuery {
	return &BlockQuery{
		config: c.config,
	}
}

// Get returns a Block entity by its id.
func (c *BlockClient) Get(ctx context.Context, id uuid.UUID) (*Block, error) {
	return c.Query().Where(block.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlockClient) GetX(ctx context.Context, id uuid.UUID) *Block {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BlockClient) Hooks() []Hook {
	hooks := c.hooks.Block
	return append(hooks[:len(hooks):len(hooks)], block.Hooks[:]...)
}

// ContractClient is a client for the Contract schema.
type ContractClient struct {
	config
}

// NewContractClient returns a client for the Contract from the given config.
func NewContractClient(c config) *ContractClient {
	return &ContractClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contract.Hooks(f(g(h())))`.
func (c *ContractClient) Use(hooks ...Hook) {
	c.hooks.Contract = append(c.hooks.Contract, hooks...)
}

// Create returns a builder for creating a Contract entity.
func (c *ContractClient) Create() *ContractCreate {
	mutation := newContractMutation(c.config, OpCreate)
	return &ContractCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contract entities.
func (c *ContractClient) CreateBulk(builders ...*ContractCreate) *ContractCreateBulk {
	return &ContractCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contract.
func (c *ContractClient) Update() *ContractUpdate {
	mutation := newContractMutation(c.config, OpUpdate)
	return &ContractUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContractClient) UpdateOne(co *Contract) *ContractUpdateOne {
	mutation := newContractMutation(c.config, OpUpdateOne, withContract(co))
	return &ContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContractClient) UpdateOneID(id uuid.UUID) *ContractUpdateOne {
	mutation := newContractMutation(c.config, OpUpdateOne, withContractID(id))
	return &ContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contract.
func (c *ContractClient) Delete() *ContractDelete {
	mutation := newContractMutation(c.config, OpDelete)
	return &ContractDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContractClient) DeleteOne(co *Contract) *ContractDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ContractClient) DeleteOneID(id uuid.UUID) *ContractDeleteOne {
	builder := c.Delete().Where(contract.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContractDeleteOne{builder}
}

// Query returns a query builder for Contract.
func (c *ContractClient) Query() *ContractQuery {
	return &ContractQuery{
		config: c.config,
	}
}

// Get returns a Contract entity by its id.
func (c *ContractClient) Get(ctx context.Context, id uuid.UUID) (*Contract, error) {
	return c.Query().Where(contract.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContractClient) GetX(ctx context.Context, id uuid.UUID) *Contract {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ContractClient) Hooks() []Hook {
	hooks := c.hooks.Contract
	return append(hooks[:len(hooks):len(hooks)], contract.Hooks[:]...)
}

// EndpointClient is a client for the Endpoint schema.
type EndpointClient struct {
	config
}

// NewEndpointClient returns a client for the Endpoint from the given config.
func NewEndpointClient(c config) *EndpointClient {
	return &EndpointClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `endpoint.Hooks(f(g(h())))`.
func (c *EndpointClient) Use(hooks ...Hook) {
	c.hooks.Endpoint = append(c.hooks.Endpoint, hooks...)
}

// Create returns a builder for creating a Endpoint entity.
func (c *EndpointClient) Create() *EndpointCreate {
	mutation := newEndpointMutation(c.config, OpCreate)
	return &EndpointCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Endpoint entities.
func (c *EndpointClient) CreateBulk(builders ...*EndpointCreate) *EndpointCreateBulk {
	return &EndpointCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Endpoint.
func (c *EndpointClient) Update() *EndpointUpdate {
	mutation := newEndpointMutation(c.config, OpUpdate)
	return &EndpointUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EndpointClient) UpdateOne(e *Endpoint) *EndpointUpdateOne {
	mutation := newEndpointMutation(c.config, OpUpdateOne, withEndpoint(e))
	return &EndpointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EndpointClient) UpdateOneID(id uuid.UUID) *EndpointUpdateOne {
	mutation := newEndpointMutation(c.config, OpUpdateOne, withEndpointID(id))
	return &EndpointUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Endpoint.
func (c *EndpointClient) Delete() *EndpointDelete {
	mutation := newEndpointMutation(c.config, OpDelete)
	return &EndpointDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EndpointClient) DeleteOne(e *Endpoint) *EndpointDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *EndpointClient) DeleteOneID(id uuid.UUID) *EndpointDeleteOne {
	builder := c.Delete().Where(endpoint.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EndpointDeleteOne{builder}
}

// Query returns a query builder for Endpoint.
func (c *EndpointClient) Query() *EndpointQuery {
	return &EndpointQuery{
		config: c.config,
	}
}

// Get returns a Endpoint entity by its id.
func (c *EndpointClient) Get(ctx context.Context, id uuid.UUID) (*Endpoint, error) {
	return c.Query().Where(endpoint.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EndpointClient) GetX(ctx context.Context, id uuid.UUID) *Endpoint {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *EndpointClient) Hooks() []Hook {
	hooks := c.hooks.Endpoint
	return append(hooks[:len(hooks):len(hooks)], endpoint.Hooks[:]...)
}

// SnapshotClient is a client for the Snapshot schema.
type SnapshotClient struct {
	config
}

// NewSnapshotClient returns a client for the Snapshot from the given config.
func NewSnapshotClient(c config) *SnapshotClient {
	return &SnapshotClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `snapshot.Hooks(f(g(h())))`.
func (c *SnapshotClient) Use(hooks ...Hook) {
	c.hooks.Snapshot = append(c.hooks.Snapshot, hooks...)
}

// Create returns a builder for creating a Snapshot entity.
func (c *SnapshotClient) Create() *SnapshotCreate {
	mutation := newSnapshotMutation(c.config, OpCreate)
	return &SnapshotCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Snapshot entities.
func (c *SnapshotClient) CreateBulk(builders ...*SnapshotCreate) *SnapshotCreateBulk {
	return &SnapshotCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Snapshot.
func (c *SnapshotClient) Update() *SnapshotUpdate {
	mutation := newSnapshotMutation(c.config, OpUpdate)
	return &SnapshotUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SnapshotClient) UpdateOne(s *Snapshot) *SnapshotUpdateOne {
	mutation := newSnapshotMutation(c.config, OpUpdateOne, withSnapshot(s))
	return &SnapshotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SnapshotClient) UpdateOneID(id uuid.UUID) *SnapshotUpdateOne {
	mutation := newSnapshotMutation(c.config, OpUpdateOne, withSnapshotID(id))
	return &SnapshotUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Snapshot.
func (c *SnapshotClient) Delete() *SnapshotDelete {
	mutation := newSnapshotMutation(c.config, OpDelete)
	return &SnapshotDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SnapshotClient) DeleteOne(s *Snapshot) *SnapshotDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SnapshotClient) DeleteOneID(id uuid.UUID) *SnapshotDeleteOne {
	builder := c.Delete().Where(snapshot.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SnapshotDeleteOne{builder}
}

// Query returns a query builder for Snapshot.
func (c *SnapshotClient) Query() *SnapshotQuery {
	return &SnapshotQuery{
		config: c.config,
	}
}

// Get returns a Snapshot entity by its id.
func (c *SnapshotClient) Get(ctx context.Context, id uuid.UUID) (*Snapshot, error) {
	return c.Query().Where(snapshot.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SnapshotClient) GetX(ctx context.Context, id uuid.UUID) *Snapshot {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SnapshotClient) Hooks() []Hook {
	hooks := c.hooks.Snapshot
	return append(hooks[:len(hooks):len(hooks)], snapshot.Hooks[:]...)
}

// SyncTaskClient is a client for the SyncTask schema.
type SyncTaskClient struct {
	config
}

// NewSyncTaskClient returns a client for the SyncTask from the given config.
func NewSyncTaskClient(c config) *SyncTaskClient {
	return &SyncTaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `synctask.Hooks(f(g(h())))`.
func (c *SyncTaskClient) Use(hooks ...Hook) {
	c.hooks.SyncTask = append(c.hooks.SyncTask, hooks...)
}

// Create returns a builder for creating a SyncTask entity.
func (c *SyncTaskClient) Create() *SyncTaskCreate {
	mutation := newSyncTaskMutation(c.config, OpCreate)
	return &SyncTaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SyncTask entities.
func (c *SyncTaskClient) CreateBulk(builders ...*SyncTaskCreate) *SyncTaskCreateBulk {
	return &SyncTaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SyncTask.
func (c *SyncTaskClient) Update() *SyncTaskUpdate {
	mutation := newSyncTaskMutation(c.config, OpUpdate)
	return &SyncTaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SyncTaskClient) UpdateOne(st *SyncTask) *SyncTaskUpdateOne {
	mutation := newSyncTaskMutation(c.config, OpUpdateOne, withSyncTask(st))
	return &SyncTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SyncTaskClient) UpdateOneID(id uuid.UUID) *SyncTaskUpdateOne {
	mutation := newSyncTaskMutation(c.config, OpUpdateOne, withSyncTaskID(id))
	return &SyncTaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SyncTask.
func (c *SyncTaskClient) Delete() *SyncTaskDelete {
	mutation := newSyncTaskMutation(c.config, OpDelete)
	return &SyncTaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SyncTaskClient) DeleteOne(st *SyncTask) *SyncTaskDeleteOne {
	return c.DeleteOneID(st.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SyncTaskClient) DeleteOneID(id uuid.UUID) *SyncTaskDeleteOne {
	builder := c.Delete().Where(synctask.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SyncTaskDeleteOne{builder}
}

// Query returns a query builder for SyncTask.
func (c *SyncTaskClient) Query() *SyncTaskQuery {
	return &SyncTaskQuery{
		config: c.config,
	}
}

// Get returns a SyncTask entity by its id.
func (c *SyncTaskClient) Get(ctx context.Context, id uuid.UUID) (*SyncTask, error) {
	return c.Query().Where(synctask.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SyncTaskClient) GetX(ctx context.Context, id uuid.UUID) *SyncTask {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SyncTaskClient) Hooks() []Hook {
	hooks := c.hooks.SyncTask
	return append(hooks[:len(hooks):len(hooks)], synctask.Hooks[:]...)
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id uuid.UUID) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id uuid.UUID) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id uuid.UUID) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id uuid.UUID) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	hooks := c.hooks.Token
	return append(hooks[:len(hooks):len(hooks)], token.Hooks[:]...)
}

// TransferClient is a client for the Transfer schema.
type TransferClient struct {
	config
}

// NewTransferClient returns a client for the Transfer from the given config.
func NewTransferClient(c config) *TransferClient {
	return &TransferClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `transfer.Hooks(f(g(h())))`.
func (c *TransferClient) Use(hooks ...Hook) {
	c.hooks.Transfer = append(c.hooks.Transfer, hooks...)
}

// Create returns a builder for creating a Transfer entity.
func (c *TransferClient) Create() *TransferCreate {
	mutation := newTransferMutation(c.config, OpCreate)
	return &TransferCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Transfer entities.
func (c *TransferClient) CreateBulk(builders ...*TransferCreate) *TransferCreateBulk {
	return &TransferCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Transfer.
func (c *TransferClient) Update() *TransferUpdate {
	mutation := newTransferMutation(c.config, OpUpdate)
	return &TransferUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TransferClient) UpdateOne(t *Transfer) *TransferUpdateOne {
	mutation := newTransferMutation(c.config, OpUpdateOne, withTransfer(t))
	return &TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TransferClient) UpdateOneID(id uuid.UUID) *TransferUpdateOne {
	mutation := newTransferMutation(c.config, OpUpdateOne, withTransferID(id))
	return &TransferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Transfer.
func (c *TransferClient) Delete() *TransferDelete {
	mutation := newTransferMutation(c.config, OpDelete)
	return &TransferDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TransferClient) DeleteOne(t *Transfer) *TransferDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TransferClient) DeleteOneID(id uuid.UUID) *TransferDeleteOne {
	builder := c.Delete().Where(transfer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TransferDeleteOne{builder}
}

// Query returns a query builder for Transfer.
func (c *TransferClient) Query() *TransferQuery {
	return &TransferQuery{
		config: c.config,
	}
}

// Get returns a Transfer entity by its id.
func (c *TransferClient) Get(ctx context.Context, id uuid.UUID) (*Transfer, error) {
	return c.Query().Where(transfer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TransferClient) GetX(ctx context.Context, id uuid.UUID) *Transfer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TransferClient) Hooks() []Hook {
	hooks := c.hooks.Transfer
	return append(hooks[:len(hooks):len(hooks)], transfer.Hooks[:]...)
}
