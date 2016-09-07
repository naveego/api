package cli

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
)

// Package represents a buildable publisher or subscriber.
type Package struct {
	CommandImport string
	TypeName      string
	TypeImport    string
}

// NewPublisherPackage creates a package for building a publisher
func NewPublisherPackage(typeName, typeImport string) Package {
	return Package{
		CommandImport: "github.com/naveego/api/pipeline/cli/pub",
		TypeName:      typeName,
		TypeImport:    typeImport,
	}
}

// NewSubscriberPackage creates a package for building a subscriber
func NewSubscriberPackage(typeName, typeImport string) Package {
	return Package{
		CommandImport: "github.com/naveego/api/pipeline/cli/sub",
		TypeName:      typeName,
		TypeImport:    typeImport,
	}
}

// BuildPackage generates the package executable from a package definition
func BuildPackage(pkg Package) (string, string, error) {
	dir, err := ioutil.TempDir("", "pipeline")
	if err != nil {
		return "", "", err
	}
	//defer os.RemoveAll(dir)

	req := harnessFileRequest{
		Directory:   dir,
		FileName:    "main.go",
		Name:        pkg.TypeName,
		ImportPath:  pkg.TypeImport,
		CommandPath: pkg.CommandImport,
	}

	err = generateHarnessFile(req)
	if err != nil {
		return "", "", err
	}

	c := NewCommand(dir, "go", "build", "main.go").Execute()

	tmpPkg := filepath.Join(dir, "main.exe")

	// Not using os.rename due to issue if the temp Directory
	// is on a different disk than the output
	in, err := os.Open(tmpPkg)
	if err != nil {
		return "", c.Output, err
	}
	defer in.Close()

	out, err := os.Create("main.exe")
	if err != nil {
		return "", c.Output, err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return "", c.Output, err
	}

	return "main.exe", c.Output, c.Error
}

// RunPackage executes the package and returns the output
func RunPackage(pkg Package, args ...string) (string, error) {
	dir, err := ioutil.TempDir("", "pipeline")
	if err != nil {
		return "", err
	}

	if logrus.GetLevel() != logrus.DebugLevel {
		defer os.RemoveAll(dir)
	}

	req := harnessFileRequest{
		Directory:   dir,
		FileName:    "main.go",
		Name:        pkg.TypeName,
		ImportPath:  pkg.TypeImport,
		CommandPath: pkg.CommandImport,
	}

	err = generateHarnessFile(req)
	if err != nil {
		return "", err
	}

	cmdArgs := []string{"run", "main.go"}
	for _, v := range args {
		cmdArgs = append(cmdArgs, v)
	}
	c := NewCommand(dir, "go", cmdArgs...).Execute()
	return c.Output, c.Error
}

type harnessFileRequest struct {
	Directory   string
	FileName    string
	Name        string
	ImportPath  string
	CommandPath string
}

func generateHarnessFile(request harnessFileRequest) error {
	content := strings.Replace(harnessTemplate, "@importPath", request.ImportPath, -1)
	content = strings.Replace(content, "@name", request.Name, -1)
	content = strings.Replace(content, "@commandPath", request.CommandPath, -1)
	tmpfn := filepath.Join(request.Directory, request.FileName)
	err := ioutil.WriteFile(tmpfn, []byte(content), 0666)
	return err
}

var harnessTemplate = `package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	cmd "@commandPath"
	_ "@importPath"
)

func main() {
	
	cmd.TypeName = "@name"
	
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

}`
