// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/web3eye-io/Web3Eye/nft-meta/pkg/db/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderItemType applies equality check predicate on the "order_item_type" field. It's identical to OrderItemTypeEQ.
func OrderItemType(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderItemType), v))
	})
}

// Contract applies equality check predicate on the "contract" field. It's identical to ContractEQ.
func Contract(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContract), v))
	})
}

// TokenType applies equality check predicate on the "token_type" field. It's identical to TokenTypeEQ.
func TokenType(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenType), v))
	})
}

// TokenID applies equality check predicate on the "token_id" field. It's identical to TokenIDEQ.
func TokenID(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenID), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderItemTypeEQ applies the EQ predicate on the "order_item_type" field.
func OrderItemTypeEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeNEQ applies the NEQ predicate on the "order_item_type" field.
func OrderItemTypeNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeIn applies the In predicate on the "order_item_type" field.
func OrderItemTypeIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderItemType), v...))
	})
}

// OrderItemTypeNotIn applies the NotIn predicate on the "order_item_type" field.
func OrderItemTypeNotIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderItemType), v...))
	})
}

// OrderItemTypeGT applies the GT predicate on the "order_item_type" field.
func OrderItemTypeGT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeGTE applies the GTE predicate on the "order_item_type" field.
func OrderItemTypeGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeLT applies the LT predicate on the "order_item_type" field.
func OrderItemTypeLT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeLTE applies the LTE predicate on the "order_item_type" field.
func OrderItemTypeLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeContains applies the Contains predicate on the "order_item_type" field.
func OrderItemTypeContains(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeHasPrefix applies the HasPrefix predicate on the "order_item_type" field.
func OrderItemTypeHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeHasSuffix applies the HasSuffix predicate on the "order_item_type" field.
func OrderItemTypeHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeEqualFold applies the EqualFold predicate on the "order_item_type" field.
func OrderItemTypeEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderItemType), v))
	})
}

// OrderItemTypeContainsFold applies the ContainsFold predicate on the "order_item_type" field.
func OrderItemTypeContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderItemType), v))
	})
}

// ContractEQ applies the EQ predicate on the "contract" field.
func ContractEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContract), v))
	})
}

// ContractNEQ applies the NEQ predicate on the "contract" field.
func ContractNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContract), v))
	})
}

// ContractIn applies the In predicate on the "contract" field.
func ContractIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContract), v...))
	})
}

// ContractNotIn applies the NotIn predicate on the "contract" field.
func ContractNotIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContract), v...))
	})
}

// ContractGT applies the GT predicate on the "contract" field.
func ContractGT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContract), v))
	})
}

// ContractGTE applies the GTE predicate on the "contract" field.
func ContractGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContract), v))
	})
}

// ContractLT applies the LT predicate on the "contract" field.
func ContractLT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContract), v))
	})
}

// ContractLTE applies the LTE predicate on the "contract" field.
func ContractLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContract), v))
	})
}

// ContractContains applies the Contains predicate on the "contract" field.
func ContractContains(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContract), v))
	})
}

// ContractHasPrefix applies the HasPrefix predicate on the "contract" field.
func ContractHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContract), v))
	})
}

// ContractHasSuffix applies the HasSuffix predicate on the "contract" field.
func ContractHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContract), v))
	})
}

// ContractEqualFold applies the EqualFold predicate on the "contract" field.
func ContractEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContract), v))
	})
}

// ContractContainsFold applies the ContainsFold predicate on the "contract" field.
func ContractContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContract), v))
	})
}

// TokenTypeEQ applies the EQ predicate on the "token_type" field.
func TokenTypeEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenType), v))
	})
}

// TokenTypeNEQ applies the NEQ predicate on the "token_type" field.
func TokenTypeNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenType), v))
	})
}

// TokenTypeIn applies the In predicate on the "token_type" field.
func TokenTypeIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenType), v...))
	})
}

// TokenTypeNotIn applies the NotIn predicate on the "token_type" field.
func TokenTypeNotIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenType), v...))
	})
}

// TokenTypeGT applies the GT predicate on the "token_type" field.
func TokenTypeGT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenType), v))
	})
}

// TokenTypeGTE applies the GTE predicate on the "token_type" field.
func TokenTypeGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenType), v))
	})
}

// TokenTypeLT applies the LT predicate on the "token_type" field.
func TokenTypeLT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenType), v))
	})
}

// TokenTypeLTE applies the LTE predicate on the "token_type" field.
func TokenTypeLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenType), v))
	})
}

// TokenTypeContains applies the Contains predicate on the "token_type" field.
func TokenTypeContains(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenType), v))
	})
}

// TokenTypeHasPrefix applies the HasPrefix predicate on the "token_type" field.
func TokenTypeHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenType), v))
	})
}

// TokenTypeHasSuffix applies the HasSuffix predicate on the "token_type" field.
func TokenTypeHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenType), v))
	})
}

// TokenTypeEqualFold applies the EqualFold predicate on the "token_type" field.
func TokenTypeEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenType), v))
	})
}

// TokenTypeContainsFold applies the ContainsFold predicate on the "token_type" field.
func TokenTypeContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenType), v))
	})
}

// TokenIDEQ applies the EQ predicate on the "token_id" field.
func TokenIDEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTokenID), v))
	})
}

// TokenIDNEQ applies the NEQ predicate on the "token_id" field.
func TokenIDNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTokenID), v))
	})
}

// TokenIDIn applies the In predicate on the "token_id" field.
func TokenIDIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTokenID), v...))
	})
}

// TokenIDNotIn applies the NotIn predicate on the "token_id" field.
func TokenIDNotIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTokenID), v...))
	})
}

// TokenIDGT applies the GT predicate on the "token_id" field.
func TokenIDGT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTokenID), v))
	})
}

// TokenIDGTE applies the GTE predicate on the "token_id" field.
func TokenIDGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTokenID), v))
	})
}

// TokenIDLT applies the LT predicate on the "token_id" field.
func TokenIDLT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTokenID), v))
	})
}

// TokenIDLTE applies the LTE predicate on the "token_id" field.
func TokenIDLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTokenID), v))
	})
}

// TokenIDContains applies the Contains predicate on the "token_id" field.
func TokenIDContains(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTokenID), v))
	})
}

// TokenIDHasPrefix applies the HasPrefix predicate on the "token_id" field.
func TokenIDHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTokenID), v))
	})
}

// TokenIDHasSuffix applies the HasSuffix predicate on the "token_id" field.
func TokenIDHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTokenID), v))
	})
}

// TokenIDEqualFold applies the EqualFold predicate on the "token_id" field.
func TokenIDEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTokenID), v))
	})
}

// TokenIDContainsFold applies the ContainsFold predicate on the "token_id" field.
func TokenIDContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTokenID), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAmount), v))
	})
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAmount), v))
	})
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAmount), v))
	})
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAmount), v))
	})
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAmount), v))
	})
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemark), v))
	})
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemark), v))
	})
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemark), v...))
	})
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.OrderItem {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemark), v...))
	})
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemark), v))
	})
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemark), v))
	})
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemark), v))
	})
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemark), v))
	})
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemark), v))
	})
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemark), v))
	})
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemark), v))
	})
}

// RemarkIsNil applies the IsNil predicate on the "remark" field.
func RemarkIsNil() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRemark)))
	})
}

// RemarkNotNil applies the NotNil predicate on the "remark" field.
func RemarkNotNil() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRemark)))
	})
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemark), v))
	})
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemark), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
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
func Not(p predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		p(s.Not())
	})
}
