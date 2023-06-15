// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ContractsColumns holds the columns for the "contracts" table.
	ContractsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "chain_type", Type: field.TypeString},
		{Name: "chain_id", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "symbol", Type: field.TypeString},
		{Name: "creator", Type: field.TypeString, Nullable: true},
		{Name: "block_num", Type: field.TypeUint64, Nullable: true},
		{Name: "tx_hash", Type: field.TypeString, Nullable: true},
		{Name: "tx_time", Type: field.TypeUint32, Nullable: true},
		{Name: "profile_url", Type: field.TypeString, Nullable: true},
		{Name: "base_url", Type: field.TypeString, Nullable: true},
		{Name: "banner_url", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// ContractsTable holds the schema information for the "contracts" table.
	ContractsTable = &schema.Table{
		Name:       "contracts",
		Columns:    ContractsColumns,
		PrimaryKey: []*schema.Column{ContractsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "contract_chain_type_chain_id_address",
				Unique:  true,
				Columns: []*schema.Column{ContractsColumns[4], ContractsColumns[5], ContractsColumns[6]},
			},
		},
	}
	// SnapshotsColumns holds the columns for the "snapshots" table.
	SnapshotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "index", Type: field.TypeUint64},
		{Name: "snapshot_comm_p", Type: field.TypeString},
		{Name: "snapshot_root", Type: field.TypeString},
		{Name: "snapshot_uri", Type: field.TypeString},
		{Name: "backup_state", Type: field.TypeString},
	}
	// SnapshotsTable holds the schema information for the "snapshots" table.
	SnapshotsTable = &schema.Table{
		Name:       "snapshots",
		Columns:    SnapshotsColumns,
		PrimaryKey: []*schema.Column{SnapshotsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "snapshot_index_backup_state",
				Unique:  true,
				Columns: []*schema.Column{SnapshotsColumns[4], SnapshotsColumns[8]},
			},
		},
	}
	// SyncTasksColumns holds the columns for the "sync_tasks" table.
	SyncTasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "chain_type", Type: field.TypeString, Nullable: true, Default: "ChainUnkonwn"},
		{Name: "chain_id", Type: field.TypeString},
		{Name: "start", Type: field.TypeUint64},
		{Name: "end", Type: field.TypeUint64},
		{Name: "current", Type: field.TypeUint64},
		{Name: "topic", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "sync_state", Type: field.TypeString, Nullable: true, Default: "Default"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// SyncTasksTable holds the schema information for the "sync_tasks" table.
	SyncTasksTable = &schema.Table{
		Name:       "sync_tasks",
		Columns:    SyncTasksColumns,
		PrimaryKey: []*schema.Column{SyncTasksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "synctask_topic",
				Unique:  false,
				Columns: []*schema.Column{SyncTasksColumns[9]},
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "chain_type", Type: field.TypeString},
		{Name: "chain_id", Type: field.TypeString},
		{Name: "contract", Type: field.TypeString},
		{Name: "token_type", Type: field.TypeString},
		{Name: "token_id", Type: field.TypeString},
		{Name: "owner", Type: field.TypeString, Nullable: true},
		{Name: "uri", Type: field.TypeString, Nullable: true},
		{Name: "uri_type", Type: field.TypeString, Nullable: true},
		{Name: "image_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "video_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "vector_id", Type: field.TypeInt64, Nullable: true},
		{Name: "vector_state", Type: field.TypeString, Nullable: true, Default: "Default"},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "ipfs_image_url", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "image_cid", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "token_contract_token_id",
				Unique:  true,
				Columns: []*schema.Column{TokensColumns[6], TokensColumns[8]},
			},
		},
	}
	// TransfersColumns holds the columns for the "transfers" table.
	TransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "chain_type", Type: field.TypeString},
		{Name: "chain_id", Type: field.TypeString},
		{Name: "contract", Type: field.TypeString},
		{Name: "token_type", Type: field.TypeString},
		{Name: "token_id", Type: field.TypeString},
		{Name: "from", Type: field.TypeString},
		{Name: "to", Type: field.TypeString},
		{Name: "amount", Type: field.TypeUint64},
		{Name: "block_number", Type: field.TypeUint64},
		{Name: "tx_hash", Type: field.TypeString},
		{Name: "block_hash", Type: field.TypeString},
		{Name: "tx_time", Type: field.TypeUint32, Nullable: true},
		{Name: "remark", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// TransfersTable holds the schema information for the "transfers" table.
	TransfersTable = &schema.Table{
		Name:       "transfers",
		Columns:    TransfersColumns,
		PrimaryKey: []*schema.Column{TransfersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "transfer_contract_token_id",
				Unique:  false,
				Columns: []*schema.Column{TransfersColumns[6], TransfersColumns[8]},
			},
			{
				Name:    "transfer_tx_hash_token_id",
				Unique:  false,
				Columns: []*schema.Column{TransfersColumns[13], TransfersColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ContractsTable,
		SnapshotsTable,
		SyncTasksTable,
		TokensTable,
		TransfersTable,
	}
)

func init() {
}
