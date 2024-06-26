// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"PopcornMovie/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// Total applies equality check predicate on the "total" field. It's identical to TotalEQ.
func Total(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTotal, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUserID, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUpdatedAt, v))
}

// TotalEQ applies the EQ predicate on the "total" field.
func TotalEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTotal, v))
}

// TotalNEQ applies the NEQ predicate on the "total" field.
func TotalNEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldTotal, v))
}

// TotalIn applies the In predicate on the "total" field.
func TotalIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldTotal, vs...))
}

// TotalNotIn applies the NotIn predicate on the "total" field.
func TotalNotIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldTotal, vs...))
}

// TotalGT applies the GT predicate on the "total" field.
func TotalGT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldTotal, v))
}

// TotalGTE applies the GTE predicate on the "total" field.
func TotalGTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldTotal, v))
}

// TotalLT applies the LT predicate on the "total" field.
func TotalLT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldTotal, v))
}

// TotalLTE applies the LTE predicate on the "total" field.
func TotalLTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldTotal, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldUserID, vs...))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldCode, v))
}

// CodeIsNil applies the IsNil predicate on the "code" field.
func CodeIsNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldIsNull(FieldCode))
}

// CodeNotNil applies the NotNil predicate on the "code" field.
func CodeNotNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldNotNull(FieldCode))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTickets applies the HasEdge predicate on the "tickets" edge.
func HasTickets() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTicketsWith applies the HasEdge predicate on the "tickets" edge with a given conditions (other predicates).
func HasTicketsWith(preds ...predicate.Ticket) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newTicketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFoodOrderLines applies the HasEdge predicate on the "food_order_lines" edge.
func HasFoodOrderLines() predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FoodOrderLinesTable, FoodOrderLinesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFoodOrderLinesWith applies the HasEdge predicate on the "food_order_lines" edge with a given conditions (other predicates).
func HasFoodOrderLinesWith(preds ...predicate.FoodOrderLine) predicate.Transaction {
	return predicate.Transaction(func(s *sql.Selector) {
		step := newFoodOrderLinesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.NotPredicates(p))
}
