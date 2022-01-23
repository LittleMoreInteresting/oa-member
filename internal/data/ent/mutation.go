// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"oa-member/internal/data/ent/member"
	"oa-member/internal/data/ent/predicate"
	"sync"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMember = "Member"
)

// MemberMutation represents an operation that mutates the Member nodes in the graph.
type MemberMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	mobile        *string
	_Email        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Member, error)
	predicates    []predicate.Member
}

var _ ent.Mutation = (*MemberMutation)(nil)

// memberOption allows management of the mutation configuration using functional options.
type memberOption func(*MemberMutation)

// newMemberMutation creates new mutation for the Member entity.
func newMemberMutation(c config, op Op, opts ...memberOption) *MemberMutation {
	m := &MemberMutation{
		config:        c,
		op:            op,
		typ:           TypeMember,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMemberID sets the ID field of the mutation.
func withMemberID(id int) memberOption {
	return func(m *MemberMutation) {
		var (
			err   error
			once  sync.Once
			value *Member
		)
		m.oldValue = func(ctx context.Context) (*Member, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Member.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMember sets the old Member of the mutation.
func withMember(node *Member) memberOption {
	return func(m *MemberMutation) {
		m.oldValue = func(context.Context) (*Member, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MemberMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MemberMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MemberMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *MemberMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *MemberMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *MemberMutation) ResetName() {
	m.name = nil
}

// SetMobile sets the "mobile" field.
func (m *MemberMutation) SetMobile(s string) {
	m.mobile = &s
}

// Mobile returns the value of the "mobile" field in the mutation.
func (m *MemberMutation) Mobile() (r string, exists bool) {
	v := m.mobile
	if v == nil {
		return
	}
	return *v, true
}

// OldMobile returns the old "mobile" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldMobile(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMobile is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMobile requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMobile: %w", err)
	}
	return oldValue.Mobile, nil
}

// ResetMobile resets all changes to the "mobile" field.
func (m *MemberMutation) ResetMobile() {
	m.mobile = nil
}

// SetEmail sets the "Email" field.
func (m *MemberMutation) SetEmail(s string) {
	m._Email = &s
}

// Email returns the value of the "Email" field in the mutation.
func (m *MemberMutation) Email() (r string, exists bool) {
	v := m._Email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "Email" field's value of the Member entity.
// If the Member object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MemberMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "Email" field.
func (m *MemberMutation) ResetEmail() {
	m._Email = nil
}

// Where appends a list predicates to the MemberMutation builder.
func (m *MemberMutation) Where(ps ...predicate.Member) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *MemberMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Member).
func (m *MemberMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MemberMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.name != nil {
		fields = append(fields, member.FieldName)
	}
	if m.mobile != nil {
		fields = append(fields, member.FieldMobile)
	}
	if m._Email != nil {
		fields = append(fields, member.FieldEmail)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MemberMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case member.FieldName:
		return m.Name()
	case member.FieldMobile:
		return m.Mobile()
	case member.FieldEmail:
		return m.Email()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MemberMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case member.FieldName:
		return m.OldName(ctx)
	case member.FieldMobile:
		return m.OldMobile(ctx)
	case member.FieldEmail:
		return m.OldEmail(ctx)
	}
	return nil, fmt.Errorf("unknown Member field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) SetField(name string, value ent.Value) error {
	switch name {
	case member.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case member.FieldMobile:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMobile(v)
		return nil
	case member.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MemberMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MemberMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MemberMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Member numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MemberMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MemberMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MemberMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Member nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MemberMutation) ResetField(name string) error {
	switch name {
	case member.FieldName:
		m.ResetName()
		return nil
	case member.FieldMobile:
		m.ResetMobile()
		return nil
	case member.FieldEmail:
		m.ResetEmail()
		return nil
	}
	return fmt.Errorf("unknown Member field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MemberMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MemberMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MemberMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MemberMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MemberMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MemberMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MemberMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Member unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MemberMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Member edge %s", name)
}
