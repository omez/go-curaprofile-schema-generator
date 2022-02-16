package cura

import v2 "go-curaprofile-schema-generator/internal/cura/definition/v2"

const printerDefinitionPath = "resources/definitions/fdmprinter.def.json"
const extruderDefinitionPath = "resources/definitions/fdmextruder.def.json"

func GetPrinterDefinition(release string) (*v2.Definition, error) {
	def := &v2.Definition{}
	err := getDefinitionFromCuraGithubRepo(def, printerDefinitionPath, release)
	return def, err
}

func GetLatestPrinterDefinition() (*v2.Definition, error) {
	def := &v2.Definition{}
	err := getLatestDefinitionFromCuraGithubRepo(def, printerDefinitionPath)
	return def, err
}

func GetExtruderDefinition(release string) (*v2.Definition, error) {
	def := &v2.Definition{}
	err := getDefinitionFromCuraGithubRepo(def, extruderDefinitionPath, release)
	return def, err
}

func GetLatestExtruderDefinition() (*v2.Definition, error) {
	def := &v2.Definition{}
	err := getLatestDefinitionFromCuraGithubRepo(def, extruderDefinitionPath)
	return def, err
}
