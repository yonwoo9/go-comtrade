package comtrade

// https://github.com/influxdata/telegraf/blob/master/plugins/parsers/registry.go

// Creator is the function to create a new parser
type Creator func() Parser

// Parsers contains the registry of all known parsers (following the new style)
var Parsers = map[string]Creator{}

// Add adds a parser to the registry. Usually this function is called in the plugin's init function
func Add(name string, creator Creator) {
	Parsers[name] = creator
}
