package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/execquery,sql/upsert,intercept,schema/snapshot,sql/lock ./schema
