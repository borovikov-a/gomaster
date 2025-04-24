module gomaster/modules/main

go 1.24.1

replace gomaster/modules/ypmodule => ../ypmodule

replace gomaster/modules/somemodule => ../somemodule

require (
	gomaster/modules/somemodule v0.0.0-00010101000000-000000000000
	gomaster/modules/ypmodule v0.0.0-00010101000000-000000000000
)
