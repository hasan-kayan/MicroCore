package domain

type TenantID string

func (t TenantID) String() string { return string(t) }
