package main

//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/annotations.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/http.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/micro.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/field_behavior.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/httpbody.proto
//go:generate protoc -I. -I./openapiv2 --go_out=paths=source_relative:. ./openapiv2/annotations.proto
//go:generate protoc -I. -I./openapiv2 --go_out=paths=source_relative:. ./openapiv2/openapiv2.proto
