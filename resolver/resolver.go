package resolver

import (
	"github.com/jncmaguire/example-proto-custom-resolver/pet"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func New() (*protoregistry.Types, error) {
	var (
		err   error
		types = protoregistry.GlobalTypes // unclear whether it's equivalent to new(Type), so using this for example purposes
	)

	// Add your messages to this list
	for _, message := range []proto.Message{
		new(pet.Blank),
		new(pet.Pet),
	} {

		pr := message.ProtoReflect()

		if _, err = types.FindMessageByName(pr.Descriptor().FullName()); err == nil {
			continue
		}

		if err = types.RegisterMessage(pr.Type()); err != nil {
			break
		}
	}

	return types, err
}
