module github.com/jncmaguire/example-proto-custom-resolver/resolver

go 1.17

replace github.com/jncmaguire/example-proto-custom-resolver/pet => ../pet

require (
	github.com/jncmaguire/example-proto-custom-resolver/pet v0.0.0-20220418021637-b7f1aed15b81
	github.com/stretchr/testify v1.7.1
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/jncmaguire/example-proto-types v0.0.0-20220418020303-3a4e2e32667b // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0 // indirect
)
