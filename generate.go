package main

//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/annotations.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/http.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/field_behavior.proto
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/httpbody.proto
<<<<<<< HEAD
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/micro.proto
=======
//go:generate protoc -I. -I./api --go_out=paths=source_relative:. ./api/client.proto
>>>>>>> 75c244f1a60f (update)
//go:generate protoc -I. -I./openapiv2 --go_out=paths=source_relative:. ./openapiv2/annotations.proto
//go:generate protoc -I. -I./openapiv2 --go_out=paths=source_relative:. ./openapiv2/openapiv2.proto
//go:generate protoc -I. -I./openapiv3 --go_out=paths=source_relative:. ./openapiv3/annotations.proto
//go:generate protoc -I. -I./openapiv3 --go_out=paths=source_relative:. ./openapiv3/openapiv3.proto
//go:generate protoc -I. -I./tag --go_out=paths=source_relative:. ./tag/tag.proto
