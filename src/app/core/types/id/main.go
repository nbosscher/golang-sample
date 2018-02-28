// Package id is a collection of entity id types. They're used
// to prevent accidental usage of the wrong id.
// (e.g. mis-ordering of parameters to func(UserID int64, TenantID int64))
package id

type IDType uint64

type User IDType
type Tenant IDType
type Location IDType
