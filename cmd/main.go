package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-yaml/yaml"
	"go-curaprofile-schema-generator/internal/cura"
	v2 "go-curaprofile-schema-generator/internal/cura/definition/v2"
	"go-curaprofile-schema-generator/internal/schemas"
	"log"
	"os"
)

var configuration struct {
	Format  string
	Release string
}

func init() {
	flag.StringVar(&configuration.Format, "f", "json", "Output format (yaml, json)")
	flag.StringVar(&configuration.Release, "r", "latest", "Release value")
}

func main() {

	// get release
	var release string
	switch configuration.Release {
	case "latest":
		r, err := cura.GetLatestRelease()
		if err != nil {
			log.Fatal(err)
		}
		release = r
	default:
		release = configuration.Release
	}

	// generate definitions
	if def, err := cura.GetPrinterDefinition(release); err != nil {
		log.Fatal(err)
	} else if err := generateSchema(def, "printer-schema", release); err != nil {
		log.Fatal(err)
	}

	if def, err := cura.GetExtruderDefinition(release); err != nil {
		log.Fatal(err)
	} else if err := generateSchema(def, "extruder-schema", release); err != nil {
		log.Fatal(err)
	}

}

func generateSchema(def *v2.Definition, prefix string, release string) error {

	schema, err := schemas.GenerateSchema(def)
	if err != nil {
		log.Fatal(err)
	}

	// prepare everything
	fileName := prefix
	var data []byte
	switch configuration.Format {
	case "json":
		fileName = fileName + ".json"
		data, err = json.MarshalIndent(schema, "", "\t")
		if err != nil {
			return err
		}
	case "yaml":
		fileName = fileName + ".yaml"
		data, err = yaml.Marshal(schema)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unexpected format %s", configuration.Format)
	}

	// write file
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}
