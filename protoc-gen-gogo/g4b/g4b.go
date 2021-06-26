package g4b

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"strings"
)

const (
	syncTypeActv  = "actv"
	syncTypeUnion = "union"
	syncTypeGame  = "game"

	syncTypeRedis = "redis"
	syncTypeDB    = "db"
)

func init() {
	generator.RegisterPlugin(new(g4b))
}

type g4b struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "g4b".
func (g *g4b) Name() string {
	return "g4b"
}

// Init initializes the plugin.
func (g *g4b) Init(gen *generator.Generator) {
	g.gen = gen
}

// P forwards to g.gen.P.
func (g *g4b) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the messages in the given file.
func (g *g4b) Generate(file *generator.FileDescriptor) {
	if len(file.Messages()) == 0 {
		return
	}

	for _, msg := range file.Messages() {
		subStruct := *msg.Name + "_Cache"
		// generate sub struct
		g.P("type ", subStruct, " struct {")
		for _, f := range msg.Field {
			names := strings.Split(*f.JsonName, ",")
			for _, n := range names {
				switch n {
				case syncTypeActv:
				case syncTypeGame:
				case syncTypeUnion:
				case syncTypeRedis:
				case syncTypeDB:
				default:
					//log.Printf("invalid sync type:%s, %v", n, names)
				}
			}
		}
		g.P("}")
		// generate a mock method for sub struct
		g.P("func (m *", subStruct, ") GetMockedCache() {}")
	}
}

// GenerateImports generates the import declaration for this file.
func (g *g4b) GenerateImports(file *generator.FileDescriptor) {}
