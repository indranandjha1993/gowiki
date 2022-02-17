module example.com/main

go 1.17

replace example.com/wiki => ../wiki

replace example.com/handler => ../handler

require (
	example.com/handler v0.0.0-00010101000000-000000000000
	example.com/wiki v0.0.0-00010101000000-000000000000
)
