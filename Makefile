#
.RECIPEPREFIX = >

rundev:
> go run ./cmd/app/main.go

build:
> go build -o workoverlord ./cmd/app
