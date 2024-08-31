@echo off
@rem 事前にフォルダを作っておくこと

set schema=orders

oapi-codegen -generate "types" -package "apispec" ../api/openapi/%schema%/openapi.yaml > ./spec/apispec/schema.gen.go
oapi-codegen -generate "server" -package "apispec" ../api/openapi/%schema%/openapi.yaml > ./spec/apispec/api.gen.go
@REM oapi-codegen -generate "spec" -package "apispec" ../api/openapi/%schema%/openapi.yaml > ./spec/apispec/spec.gen.go
