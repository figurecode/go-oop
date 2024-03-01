module main

go 1.22

replace (
	circular => ../circular
	delivery => ../delivery
	timer => ../timer
)

require (
	circular v0.0.0-00010101000000-000000000000
	delivery v0.0.0-00010101000000-000000000000
	timer v0.0.0-00010101000000-000000000000
)
