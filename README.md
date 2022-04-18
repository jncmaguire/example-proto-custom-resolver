# Example: Proto Custom Resolver

This is a repository holding exploratory code given the following niche requirements:

* JSON objects must originate from proto messages
* JSON objects should be deserialized back into their proto types for use in code
* The place in which they will be used in the code will take various types of proto types.
* The consumer and the multiple producers of the proto types are maintained by different groups and intentionally decoupled.

As such, this solution provides the following:
* Semi-independent yet centralized message management with Go submodules
* A centralized resolver

As a major caveat, the generic nature of consumption in itself goes against the high specificity intended for protobufs (at least to me). But it was fun to write, so here it is.

## Adding a new type

This contains the directions for adding a new type.

### Make a submodule for your collection

* Identify your object collection. If there's not a folder for it, create one and turn it into a submodule using `go mod init`

### Add your new collection.

There are some general guidelines.

#### Create a tools file for your imports (optional)

If you have dependencies to import, and are in a Go module, create a tools file with a blank identifier for the imports you need. I did this for ease of versioning and laziness.

```go
// +build tools

package example

import (
    _ "github.com/jncmaguire/example-proto-types"
)

```

### Generate the proto code

You can run `gen.sh <package name>`.

### Add the proto message to the resolver

In [resolver](./resolver/resolver.go#L15), add your message as a pointer to the list of `proto.Message`. This helps code that uses these contracts resolve the definitions when they're send as protobuf Any objects.

To resolve the import, you will need to do add this to the `resolver` go.mod `replace github.com/jncmaguire/example-proto-custom-resolver/<your package name> => ../<your package name>`.

### Voilà! You've committed a software sin.

Merge your work and you're good to go on this godforsaken path.
