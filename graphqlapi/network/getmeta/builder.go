package getmeta

import (
	"fmt"

	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/graphql-go/graphql"
)

// New GetMeta Builder to build PeerFields
func New(peerName string, schema schema.Schema) *Builder {
	return &Builder{
		peerName: peerName,
		schema:   schema,
	}
}

// Builder for Network -> GetMeta
type Builder struct {
	peerName string
	schema   schema.Schema
}

// PeerField for Network -> GetMeta -> <Peer>
func (b *Builder) PeerField() (*graphql.Field, error) {
	kinds, err := b.buildKinds()
	if err != nil {
		return nil, fmt.Errorf("could not build kinds for peer '%s': %s", b.peerName, err)
	}

	if len(kinds) == 0 {
		// if we didn't find a single class for all kinds, it's essentially the
		// same as if this peer didn't exist
		return nil, nil
	}

	object := graphql.NewObject(graphql.ObjectConfig{
		Name:        fmt.Sprintf("WeaviateNetworkGetMeta%sObj", b.peerName),
		Fields:      kinds,
		Description: fmt.Sprintf("%s%s", descriptions.NetworkGetMetaWeaviateObjDesc, b.peerName),
	})

	field := &graphql.Field{
		Name:        fmt.Sprintf("%s%s", "Meta", b.peerName),
		Description: fmt.Sprintf("%s%s", descriptions.NetworkWeaviateDesc, b.peerName),
		Type:        object,
		Resolve:     Resolve,
	}
	return field, nil
}

func (b *Builder) buildKinds() (graphql.Fields, error) {
	fields := graphql.Fields{}

	if b.schema.Actions != nil && len(b.schema.Actions.Classes) > 0 {
		actions, err := b.buildKind(kind.ACTION_KIND)
		if err != nil {
			return nil, fmt.Errorf("could not build 'action' kind: %s", err)
		}

		fields["Actions"] = newActionsField(actions)
	}

	if b.schema.Things != nil && len(b.schema.Things.Classes) > 0 {
		things, err := b.buildKind(kind.THING_KIND)
		if err != nil {
			return nil, fmt.Errorf("could not build 'thing' kind: %s", err)
		}

		fields["Things"] = newThingsField(things)
	}

	return fields, nil
}

func newActionsField(actions *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name:        "WeaviateNetworkGetMetaActions",
		Description: descriptions.NetworkGetMetaActionsDesc,
		Type:        actions,
	}
}

func newThingsField(things *graphql.Object) *graphql.Field {
	return &graphql.Field{
		Name:        "WeaviateNetworkGetMetaThings",
		Description: descriptions.NetworkGetMetaThingsDesc,
		Type:        things,
	}
}

func (b *Builder) buildKind(k kind.Kind) (*graphql.Object, error) {
	// from here on we have legacy (unrefactored code). This method is the
	// transition

	switch k {
	case kind.ACTION_KIND:
		return ClassFieldsFromSchema(b.schema.Actions.Classes, true, b.peerName)
	case kind.THING_KIND:
		return ClassFieldsFromSchema(b.schema.Things.Classes, false, b.peerName)
	}

	return nil, fmt.Errorf("unrecognized kind '%s'", k)

}

func passThroughResolver(p graphql.ResolveParams) (interface{}, error) {
	return p.Source, nil
}