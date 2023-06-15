// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/contract"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/snapshot"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/synctask"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/token"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/transfer"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 5)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   contract.Table,
			Columns: contract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: contract.FieldID,
			},
		},
		Type: "Contract",
		Fields: map[string]*sqlgraph.FieldSpec{
			contract.FieldCreatedAt:   {Type: field.TypeUint32, Column: contract.FieldCreatedAt},
			contract.FieldUpdatedAt:   {Type: field.TypeUint32, Column: contract.FieldUpdatedAt},
			contract.FieldDeletedAt:   {Type: field.TypeUint32, Column: contract.FieldDeletedAt},
			contract.FieldChainType:   {Type: field.TypeString, Column: contract.FieldChainType},
			contract.FieldChainID:     {Type: field.TypeString, Column: contract.FieldChainID},
			contract.FieldAddress:     {Type: field.TypeString, Column: contract.FieldAddress},
			contract.FieldName:        {Type: field.TypeString, Column: contract.FieldName},
			contract.FieldSymbol:      {Type: field.TypeString, Column: contract.FieldSymbol},
			contract.FieldCreator:     {Type: field.TypeString, Column: contract.FieldCreator},
			contract.FieldBlockNum:    {Type: field.TypeUint64, Column: contract.FieldBlockNum},
			contract.FieldTxHash:      {Type: field.TypeString, Column: contract.FieldTxHash},
			contract.FieldTxTime:      {Type: field.TypeUint32, Column: contract.FieldTxTime},
			contract.FieldProfileURL:  {Type: field.TypeString, Column: contract.FieldProfileURL},
			contract.FieldBaseURL:     {Type: field.TypeString, Column: contract.FieldBaseURL},
			contract.FieldBannerURL:   {Type: field.TypeString, Column: contract.FieldBannerURL},
			contract.FieldDescription: {Type: field.TypeString, Column: contract.FieldDescription},
			contract.FieldRemark:      {Type: field.TypeString, Column: contract.FieldRemark},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   snapshot.Table,
			Columns: snapshot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: snapshot.FieldID,
			},
		},
		Type: "Snapshot",
		Fields: map[string]*sqlgraph.FieldSpec{
			snapshot.FieldCreatedAt:     {Type: field.TypeUint32, Column: snapshot.FieldCreatedAt},
			snapshot.FieldUpdatedAt:     {Type: field.TypeUint32, Column: snapshot.FieldUpdatedAt},
			snapshot.FieldDeletedAt:     {Type: field.TypeUint32, Column: snapshot.FieldDeletedAt},
			snapshot.FieldIndex:         {Type: field.TypeUint64, Column: snapshot.FieldIndex},
			snapshot.FieldSnapshotCommP: {Type: field.TypeString, Column: snapshot.FieldSnapshotCommP},
			snapshot.FieldSnapshotRoot:  {Type: field.TypeString, Column: snapshot.FieldSnapshotRoot},
			snapshot.FieldSnapshotURI:   {Type: field.TypeString, Column: snapshot.FieldSnapshotURI},
			snapshot.FieldBackupState:   {Type: field.TypeString, Column: snapshot.FieldBackupState},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   synctask.Table,
			Columns: synctask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: synctask.FieldID,
			},
		},
		Type: "SyncTask",
		Fields: map[string]*sqlgraph.FieldSpec{
			synctask.FieldCreatedAt:   {Type: field.TypeUint32, Column: synctask.FieldCreatedAt},
			synctask.FieldUpdatedAt:   {Type: field.TypeUint32, Column: synctask.FieldUpdatedAt},
			synctask.FieldDeletedAt:   {Type: field.TypeUint32, Column: synctask.FieldDeletedAt},
			synctask.FieldChainType:   {Type: field.TypeString, Column: synctask.FieldChainType},
			synctask.FieldChainID:     {Type: field.TypeString, Column: synctask.FieldChainID},
			synctask.FieldStart:       {Type: field.TypeUint64, Column: synctask.FieldStart},
			synctask.FieldEnd:         {Type: field.TypeUint64, Column: synctask.FieldEnd},
			synctask.FieldCurrent:     {Type: field.TypeUint64, Column: synctask.FieldCurrent},
			synctask.FieldTopic:       {Type: field.TypeString, Column: synctask.FieldTopic},
			synctask.FieldDescription: {Type: field.TypeString, Column: synctask.FieldDescription},
			synctask.FieldSyncState:   {Type: field.TypeString, Column: synctask.FieldSyncState},
			synctask.FieldRemark:      {Type: field.TypeString, Column: synctask.FieldRemark},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		},
		Type: "Token",
		Fields: map[string]*sqlgraph.FieldSpec{
			token.FieldCreatedAt:    {Type: field.TypeUint32, Column: token.FieldCreatedAt},
			token.FieldUpdatedAt:    {Type: field.TypeUint32, Column: token.FieldUpdatedAt},
			token.FieldDeletedAt:    {Type: field.TypeUint32, Column: token.FieldDeletedAt},
			token.FieldChainType:    {Type: field.TypeString, Column: token.FieldChainType},
			token.FieldChainID:      {Type: field.TypeString, Column: token.FieldChainID},
			token.FieldContract:     {Type: field.TypeString, Column: token.FieldContract},
			token.FieldTokenType:    {Type: field.TypeString, Column: token.FieldTokenType},
			token.FieldTokenID:      {Type: field.TypeString, Column: token.FieldTokenID},
			token.FieldOwner:        {Type: field.TypeString, Column: token.FieldOwner},
			token.FieldURI:          {Type: field.TypeString, Column: token.FieldURI},
			token.FieldURIType:      {Type: field.TypeString, Column: token.FieldURIType},
			token.FieldImageURL:     {Type: field.TypeString, Column: token.FieldImageURL},
			token.FieldVideoURL:     {Type: field.TypeString, Column: token.FieldVideoURL},
			token.FieldDescription:  {Type: field.TypeString, Column: token.FieldDescription},
			token.FieldName:         {Type: field.TypeString, Column: token.FieldName},
			token.FieldVectorID:     {Type: field.TypeInt64, Column: token.FieldVectorID},
			token.FieldVectorState:  {Type: field.TypeString, Column: token.FieldVectorState},
			token.FieldRemark:       {Type: field.TypeString, Column: token.FieldRemark},
			token.FieldIpfsImageURL: {Type: field.TypeString, Column: token.FieldIpfsImageURL},
			token.FieldImageCid:     {Type: field.TypeString, Column: token.FieldImageCid},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   transfer.Table,
			Columns: transfer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: transfer.FieldID,
			},
		},
		Type: "Transfer",
		Fields: map[string]*sqlgraph.FieldSpec{
			transfer.FieldCreatedAt:   {Type: field.TypeUint32, Column: transfer.FieldCreatedAt},
			transfer.FieldUpdatedAt:   {Type: field.TypeUint32, Column: transfer.FieldUpdatedAt},
			transfer.FieldDeletedAt:   {Type: field.TypeUint32, Column: transfer.FieldDeletedAt},
			transfer.FieldChainType:   {Type: field.TypeString, Column: transfer.FieldChainType},
			transfer.FieldChainID:     {Type: field.TypeString, Column: transfer.FieldChainID},
			transfer.FieldContract:    {Type: field.TypeString, Column: transfer.FieldContract},
			transfer.FieldTokenType:   {Type: field.TypeString, Column: transfer.FieldTokenType},
			transfer.FieldTokenID:     {Type: field.TypeString, Column: transfer.FieldTokenID},
			transfer.FieldFrom:        {Type: field.TypeString, Column: transfer.FieldFrom},
			transfer.FieldTo:          {Type: field.TypeString, Column: transfer.FieldTo},
			transfer.FieldAmount:      {Type: field.TypeUint64, Column: transfer.FieldAmount},
			transfer.FieldBlockNumber: {Type: field.TypeUint64, Column: transfer.FieldBlockNumber},
			transfer.FieldTxHash:      {Type: field.TypeString, Column: transfer.FieldTxHash},
			transfer.FieldBlockHash:   {Type: field.TypeString, Column: transfer.FieldBlockHash},
			transfer.FieldTxTime:      {Type: field.TypeUint32, Column: transfer.FieldTxTime},
			transfer.FieldRemark:      {Type: field.TypeString, Column: transfer.FieldRemark},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *ContractQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ContractQuery builder.
func (cq *ContractQuery) Filter() *ContractFilter {
	return &ContractFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *ContractMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ContractMutation builder.
func (m *ContractMutation) Filter() *ContractFilter {
	return &ContractFilter{config: m.config, predicateAdder: m}
}

// ContractFilter provides a generic filtering capability at runtime for ContractQuery.
type ContractFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *ContractFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *ContractFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(contract.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *ContractFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(contract.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *ContractFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(contract.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *ContractFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(contract.FieldDeletedAt))
}

// WhereChainType applies the entql string predicate on the chain_type field.
func (f *ContractFilter) WhereChainType(p entql.StringP) {
	f.Where(p.Field(contract.FieldChainType))
}

// WhereChainID applies the entql string predicate on the chain_id field.
func (f *ContractFilter) WhereChainID(p entql.StringP) {
	f.Where(p.Field(contract.FieldChainID))
}

// WhereAddress applies the entql string predicate on the address field.
func (f *ContractFilter) WhereAddress(p entql.StringP) {
	f.Where(p.Field(contract.FieldAddress))
}

// WhereName applies the entql string predicate on the name field.
func (f *ContractFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(contract.FieldName))
}

// WhereSymbol applies the entql string predicate on the symbol field.
func (f *ContractFilter) WhereSymbol(p entql.StringP) {
	f.Where(p.Field(contract.FieldSymbol))
}

// WhereCreator applies the entql string predicate on the creator field.
func (f *ContractFilter) WhereCreator(p entql.StringP) {
	f.Where(p.Field(contract.FieldCreator))
}

// WhereBlockNum applies the entql uint64 predicate on the block_num field.
func (f *ContractFilter) WhereBlockNum(p entql.Uint64P) {
	f.Where(p.Field(contract.FieldBlockNum))
}

// WhereTxHash applies the entql string predicate on the tx_hash field.
func (f *ContractFilter) WhereTxHash(p entql.StringP) {
	f.Where(p.Field(contract.FieldTxHash))
}

// WhereTxTime applies the entql uint32 predicate on the tx_time field.
func (f *ContractFilter) WhereTxTime(p entql.Uint32P) {
	f.Where(p.Field(contract.FieldTxTime))
}

// WhereProfileURL applies the entql string predicate on the profile_url field.
func (f *ContractFilter) WhereProfileURL(p entql.StringP) {
	f.Where(p.Field(contract.FieldProfileURL))
}

// WhereBaseURL applies the entql string predicate on the base_url field.
func (f *ContractFilter) WhereBaseURL(p entql.StringP) {
	f.Where(p.Field(contract.FieldBaseURL))
}

// WhereBannerURL applies the entql string predicate on the banner_url field.
func (f *ContractFilter) WhereBannerURL(p entql.StringP) {
	f.Where(p.Field(contract.FieldBannerURL))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *ContractFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(contract.FieldDescription))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *ContractFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(contract.FieldRemark))
}

// addPredicate implements the predicateAdder interface.
func (sq *SnapshotQuery) addPredicate(pred func(s *sql.Selector)) {
	sq.predicates = append(sq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SnapshotQuery builder.
func (sq *SnapshotQuery) Filter() *SnapshotFilter {
	return &SnapshotFilter{config: sq.config, predicateAdder: sq}
}

// addPredicate implements the predicateAdder interface.
func (m *SnapshotMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SnapshotMutation builder.
func (m *SnapshotMutation) Filter() *SnapshotFilter {
	return &SnapshotFilter{config: m.config, predicateAdder: m}
}

// SnapshotFilter provides a generic filtering capability at runtime for SnapshotQuery.
type SnapshotFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SnapshotFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *SnapshotFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(snapshot.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *SnapshotFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(snapshot.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *SnapshotFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(snapshot.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *SnapshotFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(snapshot.FieldDeletedAt))
}

// WhereIndex applies the entql uint64 predicate on the index field.
func (f *SnapshotFilter) WhereIndex(p entql.Uint64P) {
	f.Where(p.Field(snapshot.FieldIndex))
}

// WhereSnapshotCommP applies the entql string predicate on the snapshot_comm_p field.
func (f *SnapshotFilter) WhereSnapshotCommP(p entql.StringP) {
	f.Where(p.Field(snapshot.FieldSnapshotCommP))
}

// WhereSnapshotRoot applies the entql string predicate on the snapshot_root field.
func (f *SnapshotFilter) WhereSnapshotRoot(p entql.StringP) {
	f.Where(p.Field(snapshot.FieldSnapshotRoot))
}

// WhereSnapshotURI applies the entql string predicate on the snapshot_uri field.
func (f *SnapshotFilter) WhereSnapshotURI(p entql.StringP) {
	f.Where(p.Field(snapshot.FieldSnapshotURI))
}

// WhereBackupState applies the entql string predicate on the backup_state field.
func (f *SnapshotFilter) WhereBackupState(p entql.StringP) {
	f.Where(p.Field(snapshot.FieldBackupState))
}

// addPredicate implements the predicateAdder interface.
func (stq *SyncTaskQuery) addPredicate(pred func(s *sql.Selector)) {
	stq.predicates = append(stq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the SyncTaskQuery builder.
func (stq *SyncTaskQuery) Filter() *SyncTaskFilter {
	return &SyncTaskFilter{config: stq.config, predicateAdder: stq}
}

// addPredicate implements the predicateAdder interface.
func (m *SyncTaskMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the SyncTaskMutation builder.
func (m *SyncTaskMutation) Filter() *SyncTaskFilter {
	return &SyncTaskFilter{config: m.config, predicateAdder: m}
}

// SyncTaskFilter provides a generic filtering capability at runtime for SyncTaskQuery.
type SyncTaskFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *SyncTaskFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *SyncTaskFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(synctask.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *SyncTaskFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(synctask.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *SyncTaskFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(synctask.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *SyncTaskFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(synctask.FieldDeletedAt))
}

// WhereChainType applies the entql string predicate on the chain_type field.
func (f *SyncTaskFilter) WhereChainType(p entql.StringP) {
	f.Where(p.Field(synctask.FieldChainType))
}

// WhereChainID applies the entql string predicate on the chain_id field.
func (f *SyncTaskFilter) WhereChainID(p entql.StringP) {
	f.Where(p.Field(synctask.FieldChainID))
}

// WhereStart applies the entql uint64 predicate on the start field.
func (f *SyncTaskFilter) WhereStart(p entql.Uint64P) {
	f.Where(p.Field(synctask.FieldStart))
}

// WhereEnd applies the entql uint64 predicate on the end field.
func (f *SyncTaskFilter) WhereEnd(p entql.Uint64P) {
	f.Where(p.Field(synctask.FieldEnd))
}

// WhereCurrent applies the entql uint64 predicate on the current field.
func (f *SyncTaskFilter) WhereCurrent(p entql.Uint64P) {
	f.Where(p.Field(synctask.FieldCurrent))
}

// WhereTopic applies the entql string predicate on the topic field.
func (f *SyncTaskFilter) WhereTopic(p entql.StringP) {
	f.Where(p.Field(synctask.FieldTopic))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *SyncTaskFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(synctask.FieldDescription))
}

// WhereSyncState applies the entql string predicate on the sync_state field.
func (f *SyncTaskFilter) WhereSyncState(p entql.StringP) {
	f.Where(p.Field(synctask.FieldSyncState))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *SyncTaskFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(synctask.FieldRemark))
}

// addPredicate implements the predicateAdder interface.
func (tq *TokenQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TokenQuery builder.
func (tq *TokenQuery) Filter() *TokenFilter {
	return &TokenFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TokenMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TokenMutation builder.
func (m *TokenMutation) Filter() *TokenFilter {
	return &TokenFilter{config: m.config, predicateAdder: m}
}

// TokenFilter provides a generic filtering capability at runtime for TokenQuery.
type TokenFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TokenFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TokenFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(token.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TokenFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(token.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TokenFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(token.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TokenFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(token.FieldDeletedAt))
}

// WhereChainType applies the entql string predicate on the chain_type field.
func (f *TokenFilter) WhereChainType(p entql.StringP) {
	f.Where(p.Field(token.FieldChainType))
}

// WhereChainID applies the entql string predicate on the chain_id field.
func (f *TokenFilter) WhereChainID(p entql.StringP) {
	f.Where(p.Field(token.FieldChainID))
}

// WhereContract applies the entql string predicate on the contract field.
func (f *TokenFilter) WhereContract(p entql.StringP) {
	f.Where(p.Field(token.FieldContract))
}

// WhereTokenType applies the entql string predicate on the token_type field.
func (f *TokenFilter) WhereTokenType(p entql.StringP) {
	f.Where(p.Field(token.FieldTokenType))
}

// WhereTokenID applies the entql string predicate on the token_id field.
func (f *TokenFilter) WhereTokenID(p entql.StringP) {
	f.Where(p.Field(token.FieldTokenID))
}

// WhereOwner applies the entql string predicate on the owner field.
func (f *TokenFilter) WhereOwner(p entql.StringP) {
	f.Where(p.Field(token.FieldOwner))
}

// WhereURI applies the entql string predicate on the uri field.
func (f *TokenFilter) WhereURI(p entql.StringP) {
	f.Where(p.Field(token.FieldURI))
}

// WhereURIType applies the entql string predicate on the uri_type field.
func (f *TokenFilter) WhereURIType(p entql.StringP) {
	f.Where(p.Field(token.FieldURIType))
}

// WhereImageURL applies the entql string predicate on the image_url field.
func (f *TokenFilter) WhereImageURL(p entql.StringP) {
	f.Where(p.Field(token.FieldImageURL))
}

// WhereVideoURL applies the entql string predicate on the video_url field.
func (f *TokenFilter) WhereVideoURL(p entql.StringP) {
	f.Where(p.Field(token.FieldVideoURL))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *TokenFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(token.FieldDescription))
}

// WhereName applies the entql string predicate on the name field.
func (f *TokenFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(token.FieldName))
}

// WhereVectorID applies the entql int64 predicate on the vector_id field.
func (f *TokenFilter) WhereVectorID(p entql.Int64P) {
	f.Where(p.Field(token.FieldVectorID))
}

// WhereVectorState applies the entql string predicate on the vector_state field.
func (f *TokenFilter) WhereVectorState(p entql.StringP) {
	f.Where(p.Field(token.FieldVectorState))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *TokenFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(token.FieldRemark))
}

// WhereIpfsImageURL applies the entql string predicate on the ipfs_image_url field.
func (f *TokenFilter) WhereIpfsImageURL(p entql.StringP) {
	f.Where(p.Field(token.FieldIpfsImageURL))
}

// WhereImageCid applies the entql string predicate on the image_cid field.
func (f *TokenFilter) WhereImageCid(p entql.StringP) {
	f.Where(p.Field(token.FieldImageCid))
}

// addPredicate implements the predicateAdder interface.
func (tq *TransferQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TransferQuery builder.
func (tq *TransferQuery) Filter() *TransferFilter {
	return &TransferFilter{config: tq.config, predicateAdder: tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TransferMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TransferMutation builder.
func (m *TransferMutation) Filter() *TransferFilter {
	return &TransferFilter{config: m.config, predicateAdder: m}
}

// TransferFilter provides a generic filtering capability at runtime for TransferQuery.
type TransferFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *TransferFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *TransferFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(transfer.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *TransferFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(transfer.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *TransferFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(transfer.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *TransferFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(transfer.FieldDeletedAt))
}

// WhereChainType applies the entql string predicate on the chain_type field.
func (f *TransferFilter) WhereChainType(p entql.StringP) {
	f.Where(p.Field(transfer.FieldChainType))
}

// WhereChainID applies the entql string predicate on the chain_id field.
func (f *TransferFilter) WhereChainID(p entql.StringP) {
	f.Where(p.Field(transfer.FieldChainID))
}

// WhereContract applies the entql string predicate on the contract field.
func (f *TransferFilter) WhereContract(p entql.StringP) {
	f.Where(p.Field(transfer.FieldContract))
}

// WhereTokenType applies the entql string predicate on the token_type field.
func (f *TransferFilter) WhereTokenType(p entql.StringP) {
	f.Where(p.Field(transfer.FieldTokenType))
}

// WhereTokenID applies the entql string predicate on the token_id field.
func (f *TransferFilter) WhereTokenID(p entql.StringP) {
	f.Where(p.Field(transfer.FieldTokenID))
}

// WhereFrom applies the entql string predicate on the from field.
func (f *TransferFilter) WhereFrom(p entql.StringP) {
	f.Where(p.Field(transfer.FieldFrom))
}

// WhereTo applies the entql string predicate on the to field.
func (f *TransferFilter) WhereTo(p entql.StringP) {
	f.Where(p.Field(transfer.FieldTo))
}

// WhereAmount applies the entql uint64 predicate on the amount field.
func (f *TransferFilter) WhereAmount(p entql.Uint64P) {
	f.Where(p.Field(transfer.FieldAmount))
}

// WhereBlockNumber applies the entql uint64 predicate on the block_number field.
func (f *TransferFilter) WhereBlockNumber(p entql.Uint64P) {
	f.Where(p.Field(transfer.FieldBlockNumber))
}

// WhereTxHash applies the entql string predicate on the tx_hash field.
func (f *TransferFilter) WhereTxHash(p entql.StringP) {
	f.Where(p.Field(transfer.FieldTxHash))
}

// WhereBlockHash applies the entql string predicate on the block_hash field.
func (f *TransferFilter) WhereBlockHash(p entql.StringP) {
	f.Where(p.Field(transfer.FieldBlockHash))
}

// WhereTxTime applies the entql uint32 predicate on the tx_time field.
func (f *TransferFilter) WhereTxTime(p entql.Uint32P) {
	f.Where(p.Field(transfer.FieldTxTime))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *TransferFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(transfer.FieldRemark))
}
