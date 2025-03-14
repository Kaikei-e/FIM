module federation_orchestrator

go 1.23.2

require (
	buf_schema v0.0.0-00010101000000-000000000000
	connectrpc.com/connect v1.18.1
	github.com/doyensec/safeurl v0.2.1
	github.com/gorilla/mux v1.8.1
	github.com/labstack/echo/v4 v4.13.3
)

require (
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace buf_schema => ../../buf_schema
