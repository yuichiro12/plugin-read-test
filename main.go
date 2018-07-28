package read

import (
	"plugin"
	"log"
)

func save(p *plugin.Plugin, m map[string]interface{}) {
	save, err := p.Lookup("Save")
	if err != nil {
		log.Fatal(err)
	}
	save.(func(map[string]interface{}))(m)
}

func apply(p *plugin.Plugin, m map[string]interface{}) {
	apply, err := p.Lookup("Apply")
	if err != nil {
		log.Fatal(err)
	}
	apply.(func(map[string]interface{}))(m)
}