// Code generated by ent, DO NOT EDIT.

package orderpair

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// TxHash applies equality check predicate on the "tx_hash" field. It's identical to TxHashEQ.
func TxHash(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxHash), v))
	})
}

// Recipient applies equality check predicate on the "recipient" field. It's identical to RecipientEQ.
func Recipient(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// TargetID applies equality check predicate on the "target_id" field. It's identical to TargetIDEQ.
func TargetID(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetID), v))
	})
}

// BarterID applies equality check predicate on the "barter_id" field. It's identical to BarterIDEQ.
func BarterID(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBarterID), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// TxHashEQ applies the EQ predicate on the "tx_hash" field.
func TxHashEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTxHash), v))
	})
}

// TxHashNEQ applies the NEQ predicate on the "tx_hash" field.
func TxHashNEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTxHash), v))
	})
}

// TxHashIn applies the In predicate on the "tx_hash" field.
func TxHashIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTxHash), v...))
	})
}

// TxHashNotIn applies the NotIn predicate on the "tx_hash" field.
func TxHashNotIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTxHash), v...))
	})
}

// TxHashGT applies the GT predicate on the "tx_hash" field.
func TxHashGT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTxHash), v))
	})
}

// TxHashGTE applies the GTE predicate on the "tx_hash" field.
func TxHashGTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTxHash), v))
	})
}

// TxHashLT applies the LT predicate on the "tx_hash" field.
func TxHashLT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTxHash), v))
	})
}

// TxHashLTE applies the LTE predicate on the "tx_hash" field.
func TxHashLTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTxHash), v))
	})
}

// TxHashContains applies the Contains predicate on the "tx_hash" field.
func TxHashContains(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTxHash), v))
	})
}

// TxHashHasPrefix applies the HasPrefix predicate on the "tx_hash" field.
func TxHashHasPrefix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTxHash), v))
	})
}

// TxHashHasSuffix applies the HasSuffix predicate on the "tx_hash" field.
func TxHashHasSuffix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTxHash), v))
	})
}

// TxHashEqualFold applies the EqualFold predicate on the "tx_hash" field.
func TxHashEqualFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTxHash), v))
	})
}

// TxHashContainsFold applies the ContainsFold predicate on the "tx_hash" field.
func TxHashContainsFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTxHash), v))
	})
}

// RecipientEQ applies the EQ predicate on the "recipient" field.
func RecipientEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecipient), v))
	})
}

// RecipientNEQ applies the NEQ predicate on the "recipient" field.
func RecipientNEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecipient), v))
	})
}

// RecipientIn applies the In predicate on the "recipient" field.
func RecipientIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecipient), v...))
	})
}

// RecipientNotIn applies the NotIn predicate on the "recipient" field.
func RecipientNotIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecipient), v...))
	})
}

// RecipientGT applies the GT predicate on the "recipient" field.
func RecipientGT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecipient), v))
	})
}

// RecipientGTE applies the GTE predicate on the "recipient" field.
func RecipientGTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecipient), v))
	})
}

// RecipientLT applies the LT predicate on the "recipient" field.
func RecipientLT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecipient), v))
	})
}

// RecipientLTE applies the LTE predicate on the "recipient" field.
func RecipientLTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecipient), v))
	})
}

// RecipientContains applies the Contains predicate on the "recipient" field.
func RecipientContains(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRecipient), v))
	})
}

// RecipientHasPrefix applies the HasPrefix predicate on the "recipient" field.
func RecipientHasPrefix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRecipient), v))
	})
}

// RecipientHasSuffix applies the HasSuffix predicate on the "recipient" field.
func RecipientHasSuffix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRecipient), v))
	})
}

// RecipientEqualFold applies the EqualFold predicate on the "recipient" field.
func RecipientEqualFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRecipient), v))
	})
}

// RecipientContainsFold applies the ContainsFold predicate on the "recipient" field.
func RecipientContainsFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRecipient), v))
	})
}

// TargetIDEQ applies the EQ predicate on the "target_id" field.
func TargetIDEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTargetID), v))
	})
}

// TargetIDNEQ applies the NEQ predicate on the "target_id" field.
func TargetIDNEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTargetID), v))
	})
}

// TargetIDIn applies the In predicate on the "target_id" field.
func TargetIDIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTargetID), v...))
	})
}

// TargetIDNotIn applies the NotIn predicate on the "target_id" field.
func TargetIDNotIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTargetID), v...))
	})
}

// TargetIDGT applies the GT predicate on the "target_id" field.
func TargetIDGT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTargetID), v))
	})
}

// TargetIDGTE applies the GTE predicate on the "target_id" field.
func TargetIDGTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTargetID), v))
	})
}

// TargetIDLT applies the LT predicate on the "target_id" field.
func TargetIDLT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTargetID), v))
	})
}

// TargetIDLTE applies the LTE predicate on the "target_id" field.
func TargetIDLTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTargetID), v))
	})
}

// TargetIDContains applies the Contains predicate on the "target_id" field.
func TargetIDContains(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTargetID), v))
	})
}

// TargetIDHasPrefix applies the HasPrefix predicate on the "target_id" field.
func TargetIDHasPrefix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTargetID), v))
	})
}

// TargetIDHasSuffix applies the HasSuffix predicate on the "target_id" field.
func TargetIDHasSuffix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTargetID), v))
	})
}

// TargetIDEqualFold applies the EqualFold predicate on the "target_id" field.
func TargetIDEqualFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTargetID), v))
	})
}

// TargetIDContainsFold applies the ContainsFold predicate on the "target_id" field.
func TargetIDContainsFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTargetID), v))
	})
}

// BarterIDEQ applies the EQ predicate on the "barter_id" field.
func BarterIDEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBarterID), v))
	})
}

// BarterIDNEQ applies the NEQ predicate on the "barter_id" field.
func BarterIDNEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBarterID), v))
	})
}

// BarterIDIn applies the In predicate on the "barter_id" field.
func BarterIDIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBarterID), v...))
	})
}

// BarterIDNotIn applies the NotIn predicate on the "barter_id" field.
func BarterIDNotIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBarterID), v...))
	})
}

// BarterIDGT applies the GT predicate on the "barter_id" field.
func BarterIDGT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBarterID), v))
	})
}

// BarterIDGTE applies the GTE predicate on the "barter_id" field.
func BarterIDGTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBarterID), v))
	})
}

// BarterIDLT applies the LT predicate on the "barter_id" field.
func BarterIDLT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBarterID), v))
	})
}

// BarterIDLTE applies the LTE predicate on the "barter_id" field.
func BarterIDLTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBarterID), v))
	})
}

// BarterIDContains applies the Contains predicate on the "barter_id" field.
func BarterIDContains(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBarterID), v))
	})
}

// BarterIDHasPrefix applies the HasPrefix predicate on the "barter_id" field.
func BarterIDHasPrefix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBarterID), v))
	})
}

// BarterIDHasSuffix applies the HasSuffix predicate on the "barter_id" field.
func BarterIDHasSuffix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBarterID), v))
	})
}

// BarterIDEqualFold applies the EqualFold predicate on the "barter_id" field.
func BarterIDEqualFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBarterID), v))
	})
}

// BarterIDContainsFold applies the ContainsFold predicate on the "barter_id" field.
func BarterIDContainsFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBarterID), v))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.OrderPair {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemark)))
	})
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemark)))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderPair) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderPair) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
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
func Not(p predicate.OrderPair) predicate.OrderPair {
	return predicate.OrderPair(func(s *sql.Selector) {
		p(s.Not())
	})
}
