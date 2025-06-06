package main

//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/annotations.proto
//go:generate protoc -I. -I./api --python_out=. ./api/annotations.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/http.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/field_behavior.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/httpbody.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/micro.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/client.proto
//go:generate protoc -I. -I./openapiv2 --go_out=paths=source_relative:. ./openapiv2/annotations.proto
//go:generate protoc -I. -I./openapiv2 --go_out=paths=source_relative:. ./openapiv2/openapiv2.proto
//go:generate protoc -I. -I./openapiv3 --go_out=paths=source_relative:. ./openapiv3/annotations.proto
//go:generate protoc -I. -I./openapiv3 --go_out=paths=source_relative:. ./openapiv3/openapiv3.proto
//go:generate protoc -I. -I./tag --go_out=paths=source_relative:. ./tag/tag.proto
//go:generate protoc -I. -I./tag --go_out=paths=source_relative:. ./errors/errors.proto
//go:generate protoc -I. -I./tag --go_out=paths=source_relative:. ./codec/frame.proto
//go:generate protoc -I. -I./graphql --go_out=paths=source_relative:. ./graphql/graphql.proto
