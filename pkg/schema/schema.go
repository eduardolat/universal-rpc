package schema

/*
	The Schema struct is the common and most basic representation of
	a schema file regardless of the format it is written in.

	There should be a parser for each format that can convert the
	format into this struct.
*/

type UnparsedSchema map[string]any

type Schema struct {
	Version   int
	Routers   map[string]Router
	Queries   map[string]Endpoint
	Mutations map[string]Endpoint
}

type Router struct {
	Routers   map[string]Router
	Queries   map[string]Endpoint
	Mutations map[string]Endpoint
}

type Endpoint struct {
	Params  map[string]Field
	Returns map[string]Field
}

type Field struct {
	Type       string
	Required   bool // <- Must be required by default (set by the parser)
	Items      *Field
	Properties map[string]*Field
}
