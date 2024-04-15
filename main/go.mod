module main

go 1.22

replace (
	circular => ../circular
	delivery => ../delivery
	embedded => ../embedded
	loggerLevel => ../loggerLevel
	timer => ../timer
)

require (
	circular v0.0.0-00010101000000-000000000000
	delivery v0.0.0-00010101000000-000000000000
	embedded v0.0.0-00010101000000-000000000000
	loggerLevel v0.0.0-00010101000000-000000000000
	timer v0.0.0-00010101000000-000000000000
)
