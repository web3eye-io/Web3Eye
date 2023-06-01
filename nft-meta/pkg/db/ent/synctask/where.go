// Code generated by ent, DO NOT EDIT.

package synctask

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// ChainType applies equality check predicate on the "chain_type" field. It's identical to ChainTypeEQ.
func ChainType(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainType), v))
	})
}

// ChainID applies equality check predicate on the "chain_id" field. It's identical to ChainIDEQ.
func ChainID(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainID), v))
	})
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// Current applies equality check predicate on the "current" field. It's identical to CurrentEQ.
func Current(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrent), v))
	})
}

// Topic applies equality check predicate on the "topic" field. It's identical to TopicEQ.
func Topic(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopic), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// SyncState applies equality check predicate on the "sync_state" field. It's identical to SyncStateEQ.
func SyncState(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncState), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// ChainTypeEQ applies the EQ predicate on the "chain_type" field.
func ChainTypeEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainType), v))
	})
}

// ChainTypeNEQ applies the NEQ predicate on the "chain_type" field.
func ChainTypeNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainType), v))
	})
}

// ChainTypeIn applies the In predicate on the "chain_type" field.
func ChainTypeIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChainType), v...))
	})
}

// ChainTypeNotIn applies the NotIn predicate on the "chain_type" field.
func ChainTypeNotIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChainType), v...))
	})
}

// ChainTypeGT applies the GT predicate on the "chain_type" field.
func ChainTypeGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainType), v))
	})
}

// ChainTypeGTE applies the GTE predicate on the "chain_type" field.
func ChainTypeGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainType), v))
	})
}

// ChainTypeLT applies the LT predicate on the "chain_type" field.
func ChainTypeLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainType), v))
	})
}

// ChainTypeLTE applies the LTE predicate on the "chain_type" field.
func ChainTypeLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainType), v))
	})
}

// ChainTypeContains applies the Contains predicate on the "chain_type" field.
func ChainTypeContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChainType), v))
	})
}

// ChainTypeHasPrefix applies the HasPrefix predicate on the "chain_type" field.
func ChainTypeHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChainType), v))
	})
}

// ChainTypeHasSuffix applies the HasSuffix predicate on the "chain_type" field.
func ChainTypeHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChainType), v))
	})
}

// ChainTypeIsNil applies the IsNil predicate on the "chain_type" field.
func ChainTypeIsNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChainType)))
	})
}

// ChainTypeNotNil applies the NotNil predicate on the "chain_type" field.
func ChainTypeNotNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChainType)))
	})
}

// ChainTypeEqualFold applies the EqualFold predicate on the "chain_type" field.
func ChainTypeEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChainType), v))
	})
}

// ChainTypeContainsFold applies the ContainsFold predicate on the "chain_type" field.
func ChainTypeContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChainType), v))
	})
}

// ChainIDEQ applies the EQ predicate on the "chain_id" field.
func ChainIDEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChainID), v))
	})
}

// ChainIDNEQ applies the NEQ predicate on the "chain_id" field.
func ChainIDNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChainID), v))
	})
}

// ChainIDIn applies the In predicate on the "chain_id" field.
func ChainIDIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChainID), v...))
	})
}

// ChainIDNotIn applies the NotIn predicate on the "chain_id" field.
func ChainIDNotIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChainID), v...))
	})
}

// ChainIDGT applies the GT predicate on the "chain_id" field.
func ChainIDGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChainID), v))
	})
}

// ChainIDGTE applies the GTE predicate on the "chain_id" field.
func ChainIDGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChainID), v))
	})
}

// ChainIDLT applies the LT predicate on the "chain_id" field.
func ChainIDLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChainID), v))
	})
}

// ChainIDLTE applies the LTE predicate on the "chain_id" field.
func ChainIDLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChainID), v))
	})
}

// ChainIDContains applies the Contains predicate on the "chain_id" field.
func ChainIDContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChainID), v))
	})
}

// ChainIDHasPrefix applies the HasPrefix predicate on the "chain_id" field.
func ChainIDHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChainID), v))
	})
}

// ChainIDHasSuffix applies the HasSuffix predicate on the "chain_id" field.
func ChainIDHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChainID), v))
	})
}

// ChainIDEqualFold applies the EqualFold predicate on the "chain_id" field.
func ChainIDEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChainID), v))
	})
}

// ChainIDContainsFold applies the ContainsFold predicate on the "chain_id" field.
func ChainIDContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChainID), v))
	})
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStart), v))
	})
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStart), v))
	})
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...uint64) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStart), v...))
	})
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...uint64) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStart), v...))
	})
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStart), v))
	})
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStart), v))
	})
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStart), v))
	})
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStart), v))
	})
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnd), v))
	})
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnd), v))
	})
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...uint64) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEnd), v...))
	})
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...uint64) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEnd), v...))
	})
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEnd), v))
	})
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEnd), v))
	})
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEnd), v))
	})
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEnd), v))
	})
}

// CurrentEQ applies the EQ predicate on the "current" field.
func CurrentEQ(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrent), v))
	})
}

// CurrentNEQ applies the NEQ predicate on the "current" field.
func CurrentNEQ(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCurrent), v))
	})
}

// CurrentIn applies the In predicate on the "current" field.
func CurrentIn(vs ...uint64) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCurrent), v...))
	})
}

// CurrentNotIn applies the NotIn predicate on the "current" field.
func CurrentNotIn(vs ...uint64) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCurrent), v...))
	})
}

// CurrentGT applies the GT predicate on the "current" field.
func CurrentGT(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCurrent), v))
	})
}

// CurrentGTE applies the GTE predicate on the "current" field.
func CurrentGTE(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCurrent), v))
	})
}

// CurrentLT applies the LT predicate on the "current" field.
func CurrentLT(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCurrent), v))
	})
}

// CurrentLTE applies the LTE predicate on the "current" field.
func CurrentLTE(v uint64) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCurrent), v))
	})
}

// TopicEQ applies the EQ predicate on the "topic" field.
func TopicEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopic), v))
	})
}

// TopicNEQ applies the NEQ predicate on the "topic" field.
func TopicNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTopic), v))
	})
}

// TopicIn applies the In predicate on the "topic" field.
func TopicIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTopic), v...))
	})
}

// TopicNotIn applies the NotIn predicate on the "topic" field.
func TopicNotIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTopic), v...))
	})
}

// TopicGT applies the GT predicate on the "topic" field.
func TopicGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTopic), v))
	})
}

// TopicGTE applies the GTE predicate on the "topic" field.
func TopicGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTopic), v))
	})
}

// TopicLT applies the LT predicate on the "topic" field.
func TopicLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTopic), v))
	})
}

// TopicLTE applies the LTE predicate on the "topic" field.
func TopicLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTopic), v))
	})
}

// TopicContains applies the Contains predicate on the "topic" field.
func TopicContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTopic), v))
	})
}

// TopicHasPrefix applies the HasPrefix predicate on the "topic" field.
func TopicHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTopic), v))
	})
}

// TopicHasSuffix applies the HasSuffix predicate on the "topic" field.
func TopicHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTopic), v))
	})
}

// TopicEqualFold applies the EqualFold predicate on the "topic" field.
func TopicEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTopic), v))
	})
}

// TopicContainsFold applies the ContainsFold predicate on the "topic" field.
func TopicContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTopic), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// SyncStateEQ applies the EQ predicate on the "sync_state" field.
func SyncStateEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSyncState), v))
	})
}

// SyncStateNEQ applies the NEQ predicate on the "sync_state" field.
func SyncStateNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSyncState), v))
	})
}

// SyncStateIn applies the In predicate on the "sync_state" field.
func SyncStateIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSyncState), v...))
	})
}

// SyncStateNotIn applies the NotIn predicate on the "sync_state" field.
func SyncStateNotIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSyncState), v...))
	})
}

// SyncStateGT applies the GT predicate on the "sync_state" field.
func SyncStateGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSyncState), v))
	})
}

// SyncStateGTE applies the GTE predicate on the "sync_state" field.
func SyncStateGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSyncState), v))
	})
}

// SyncStateLT applies the LT predicate on the "sync_state" field.
func SyncStateLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSyncState), v))
	})
}

// SyncStateLTE applies the LTE predicate on the "sync_state" field.
func SyncStateLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSyncState), v))
	})
}

// SyncStateContains applies the Contains predicate on the "sync_state" field.
func SyncStateContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSyncState), v))
	})
}

// SyncStateHasPrefix applies the HasPrefix predicate on the "sync_state" field.
func SyncStateHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSyncState), v))
	})
}

// SyncStateHasSuffix applies the HasSuffix predicate on the "sync_state" field.
func SyncStateHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSyncState), v))
	})
}

// SyncStateIsNil applies the IsNil predicate on the "sync_state" field.
func SyncStateIsNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSyncState)))
	})
}

// SyncStateNotNil applies the NotNil predicate on the "sync_state" field.
func SyncStateNotNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSyncState)))
	})
}

// SyncStateEqualFold applies the EqualFold predicate on the "sync_state" field.
func SyncStateEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSyncState), v))
	})
}

// SyncStateContainsFold applies the ContainsFold predicate on the "sync_state" field.
func SyncStateContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSyncState), v))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.SyncTask {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemark)))
	})
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemark)))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SyncTask) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SyncTask) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SyncTask) predicate.SyncTask {
	return predicate.SyncTask(func(s *sql.Selector) {
		p(s.Not())
	})
}
