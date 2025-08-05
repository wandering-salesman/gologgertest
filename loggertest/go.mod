module loggertest

go 1.23.5

replace loggerpackage => ../loggerpackage

require (
	go.uber.org/zap v1.27.0
	loggerpackage v0.0.0-00010101000000-000000000000
)

require go.uber.org/multierr v1.10.0 // indirect
