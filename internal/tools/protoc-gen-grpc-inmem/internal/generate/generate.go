package generate

import (
	"embed"
	"fmt"
	"path/filepath"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

func Generate(gp *protogen.Plugin) error {
	g := newGenerator(gp)
	return g.generate()
}

type generator struct {
	p           *protogen.Plugin
	directories map[string]pkgInfo
}

func newGenerator(gp *protogen.Plugin) *generator {
	return &generator{
		p:           gp,
		directories: make(map[string]pkgInfo),
	}
}

type pkgInfo struct {
	impPath protogen.GoImportPath
	pkgName protogen.GoPackageName
}

func (g *generator) generate() error {
	for _, file := range g.p.Files {
		if !file.Generate {
			continue
		}

		if len(file.Services) < 1 {
			continue
		}

		err := g.generateFile(file)
		if err != nil {
			return fmt.Errorf("Failed to generate file %q: %w", file.Proto.GetName(), err)
		}
	}

	for dir, info := range g.directories {
		genFile := g.p.NewGeneratedFile(filepath.Join(dir, "inmem_grpc_stream.gen.go"), info.impPath)
		inmemTemplate.ExecuteTemplate(genFile, "inmem-stream.tmpl", map[string]string{"GoPackage": string(info.pkgName)})
	}
	return nil
}

func (g *generator) generateFile(file *protogen.File) error {
	tdata := &templateData{
		PackageName: string(file.GoPackageName),
	}

	filename := file.GeneratedFilenamePrefix + "_inmem_grpc.gen.go"
	genFile := g.p.NewGeneratedFile(filename, file.GoImportPath)

	for _, svc := range file.Services {
		svcTypes := &inmemServiceTypes{
			ClientTypeName:        genFile.QualifiedGoIdent(protogen.GoIdent{GoName: svc.GoName + "Client", GoImportPath: file.GoImportPath}),
			ServerTypeName:        genFile.QualifiedGoIdent(protogen.GoIdent{GoName: svc.GoName + "Server", GoImportPath: file.GoImportPath}),
			InmemTypeName:         genFile.QualifiedGoIdent(protogen.GoIdent{GoName: "Inmem" + svc.GoName + "Client", GoImportPath: file.GoImportPath}),
			CloningClientTypeName: genFile.QualifiedGoIdent(protogen.GoIdent{GoName: "Cloning" + svc.GoName + "Client", GoImportPath: file.GoImportPath}),
			ServiceName:           svc.GoName,
		}

		tsvc := inmemService{
			inmemServiceTypes: svcTypes,
		}

		for _, method := range svc.Methods {
			if method.Desc.IsStreamingClient() {
				// when we need these we can implement this
				panic("client streams are unsupported")
			}

			if method.Desc.IsStreamingServer() {
				tsvc.ServerStreamMethods = append(tsvc.ServerStreamMethods, &inmemMethod{
					inmemServiceTypes: svcTypes,
					Method:            method,
				})

				// record that we need to also generate the inmem stream client code
				// into this directory
				g.directories[filepath.Dir(filename)] = pkgInfo{impPath: file.GoImportPath, pkgName: file.GoPackageName}
			} else {
				tsvc.UnaryMethods = append(tsvc.UnaryMethods, &inmemMethod{
					inmemServiceTypes: svcTypes,
					Method:            method,
				})
			}
		}

		tdata.Services = append(tdata.Services, &tsvc)
	}

	err := inmemTemplate.ExecuteTemplate(genFile, "file.tmpl", &tdata)
	if err != nil {
		return fmt.Errorf("Error rendering template: %w", err)
	}

	return nil

}

type templateData struct {
	PackageName string
	Services    []*inmemService
}

type inmemService struct {
	UnaryMethods []*inmemMethod
	// ClientStreamMethods      []*protogen.Method
	ServerStreamMethods []*inmemMethod
	// BidirectionStreamMethods []*protogen.Method
	*inmemServiceTypes
}

type inmemServiceTypes struct {
	ClientTypeName        string
	ServerTypeName        string
	InmemTypeName         string
	ServiceName           string
	CloningClientTypeName string
}

type inmemMethod struct {
	Method *protogen.Method
	*inmemServiceTypes
}

var (
	//go:embed templates
	templates embed.FS

	inmemTemplate = template.Must(template.ParseFS(templates, "templates/*"))
)
