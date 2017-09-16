package types

// Auditable interface of a model that is auditable
type Auditable interface {
	// AuditableName auditable naming of a model
	AuditableName() string
}

// UserAuditable interface of a model that is auditable as a User
type UserAuditable interface {
	Auditable
	AuditableUserID() uint
}
