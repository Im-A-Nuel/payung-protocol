// Package migrations embeds the SQL migration files so cmd/migrate can run
// without needing the source tree present at runtime.
package migrations

import "embed"

//go:embed *.sql
var Files embed.FS
