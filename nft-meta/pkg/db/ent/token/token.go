// Code generated by ent, DO NOT EDIT.

package token

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldChainType holds the string denoting the chain_type field in the database.
	FieldChainType = "chain_type"
	// FieldChainID holds the string denoting the chain_id field in the database.
	FieldChainID = "chain_id"
	// FieldContract holds the string denoting the contract field in the database.
	FieldContract = "contract"
	// FieldTokenType holds the string denoting the token_type field in the database.
	FieldTokenType = "token_type"
	// FieldTokenID holds the string denoting the token_id field in the database.
	FieldTokenID = "token_id"
	// FieldOwner holds the string denoting the owner field in the database.
	FieldOwner = "owner"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// FieldURIType holds the string denoting the uri_type field in the database.
	FieldURIType = "uri_type"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldVideoURL holds the string denoting the video_url field in the database.
	FieldVideoURL = "video_url"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldVectorID holds the string denoting the vector_id field in the database.
	FieldVectorID = "vector_id"
	// FieldVectorState holds the string denoting the vector_state field in the database.
	FieldVectorState = "vector_state"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldIpfsImageURL holds the string denoting the ipfs_image_url field in the database.
	FieldIpfsImageURL = "ipfs_image_url"
	// FieldImageSnapshotID holds the string denoting the image_snapshot_id field in the database.
	FieldImageSnapshotID = "image_snapshot_id"
	// Table holds the table name of the token in the database.
	Table = "tokens"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldEntID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldChainType,
	FieldChainID,
	FieldContract,
	FieldTokenType,
	FieldTokenID,
	FieldOwner,
	FieldURI,
	FieldURIType,
	FieldImageURL,
	FieldVideoURL,
	FieldDescription,
	FieldName,
	FieldVectorID,
	FieldVectorState,
	FieldRemark,
	FieldIpfsImageURL,
	FieldImageSnapshotID,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultVectorState holds the default value on creation for the "vector_state" field.
	DefaultVectorState string
)
