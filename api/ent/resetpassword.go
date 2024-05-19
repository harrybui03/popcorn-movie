// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/resetpassword"
	"PopcornMovie/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ResetPassword is the model entity for the ResetPassword schema.
type ResetPassword struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResetPasswordQuery when eager-loading is set.
	Edges        ResetPasswordEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ResetPasswordEdges holds the relations/edges for other nodes in the graph.
type ResetPasswordEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResetPasswordEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ResetPassword) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case resetpassword.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case resetpassword.FieldID, resetpassword.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ResetPassword fields.
func (rp *ResetPassword) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resetpassword.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rp.ID = *value
			}
		case resetpassword.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				rp.UserID = *value
			}
		case resetpassword.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rp.CreatedAt = value.Time
			}
		default:
			rp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ResetPassword.
// This includes values selected through modifiers, order, etc.
func (rp *ResetPassword) Value(name string) (ent.Value, error) {
	return rp.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the ResetPassword entity.
func (rp *ResetPassword) QueryUser() *UserQuery {
	return NewResetPasswordClient(rp.config).QueryUser(rp)
}

// Update returns a builder for updating this ResetPassword.
// Note that you need to call ResetPassword.Unwrap() before calling this method if this ResetPassword
// was returned from a transaction, and the transaction was committed or rolled back.
func (rp *ResetPassword) Update() *ResetPasswordUpdateOne {
	return NewResetPasswordClient(rp.config).UpdateOne(rp)
}

// Unwrap unwraps the ResetPassword entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rp *ResetPassword) Unwrap() *ResetPassword {
	_tx, ok := rp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ResetPassword is not a transactional entity")
	}
	rp.config.driver = _tx.drv
	return rp
}

// String implements the fmt.Stringer.
func (rp *ResetPassword) String() string {
	var builder strings.Builder
	builder.WriteString("ResetPassword(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rp.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", rp.UserID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(rp.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ResetPasswords is a parsable slice of ResetPassword.
type ResetPasswords []*ResetPassword
