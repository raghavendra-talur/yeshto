package ast

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

type ModuleInfo struct {
	Packages []*packages.Package
}

// BuildModuleInfo is the entry point for the AST package
func BuildModuleInfo(path string) (*ModuleInfo, error) {
	var modInfo ModuleInfo

	mode := packages.NeedName |
		packages.NeedFiles |
		packages.NeedImports |
		packages.NeedDeps |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo |
		packages.NeedModule

	cfg := &packages.Config{
		Mode:  mode,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg, path+"/...")
	if err != nil {
		return nil, fmt.Errorf("error loading packages: %v", err)
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found")
	}

	modInfo.Packages = pkgs

	return &modInfo, nil
}
