package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/viant/neatly"
	"github.com/viant/toolbox/data"
	"github.com/viant/toolbox/url"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
)

func init() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.String("i", "", "<path nearly document> ")
	flag.String("f", "json", "<output format> json or yaml")

}

func printJSON(aMap map[string]interface{}) {
	buf, err := json.MarshalIndent(aMap, "", "\t")
	if err != nil {
		log.Fatal("failed to build JSON")
	}
	fmt.Printf("%s", buf)
}

func printYAML(aMap map[string]interface{}) {
	buf, err := yaml.Marshal(aMap)
	if err != nil {
		log.Fatal("failed to build JSON")
	}
	fmt.Printf("%s", buf)
}

func main() {
	flag.Parse()
	flagset := make(map[string]string)
	flag.Visit(func(f *flag.Flag) {
		flagset[f.Name] = f.Value.String()
	})
	input, ok := flagset["i"]
	if !ok {
		flag.PrintDefaults()
		return
	}
	var context = data.NewMap()
	var neatlyDocument = make(map[string]interface{})
	dao := neatly.NewDao("", "", "", nil)
	err := dao.Load(context, url.NewResource(input), &neatlyDocument)
	if err != nil {
		log.Fatal("failed to loead neatly document: %v %v\n", input, err)
	}
	switch strings.ToLower(flag.Lookup("f").Value.String()) {
	case "json":
		printJSON(neatlyDocument)
	case "yaml":
		printYAML(neatlyDocument)
	default:
		fmt.Printf("unsupported output format: %v", flagset["f"])
	}

}