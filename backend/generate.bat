@echo off

set schema=orders

oapi-codegen -generate "types" -package "spec" ../api/openapi/%schema%/openapi.yaml > ./spec/schema.gen.go
oapi-codegen -generate "server" -package "spec" ../api/openapi/%schema%/openapi.yaml > ./spec/api.gen.go
oapi-codegen -generate "spec" -package "spec" ../api/openapi/%schema%/openapi.yaml > ./spec/spec.gen.go
