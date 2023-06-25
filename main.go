package main

import (
	"bytes"
	"flag"
	"html/template"

	"github.com/lcmaguire/protoc-gen-lazy-proxy/pkg"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		// this enables optional fields to be supported.
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

		return nil
	})
}

func Generate(gen *protogen.Plugin) error {

	serviceInfo := make([]pkg.LazyProxyServiceInfo, 0)
	methodInformation := make([]pkg.LazyProxyMethodInfo, 0)

	for _, file := range gen.Files {
		for _, service := range file.Services {
			serviceName := string(service.Desc.Name())

			sInfo := pkg.LazyProxyServiceInfo{
				ServiceName: serviceName,
				Pkg:         service.GoName,
			}
			serviceInfo = append(serviceInfo, sInfo)

			for _, method := range service.Methods {
				// todo check if streaming and skip.

				//requestType := //getParamPKG(v.Input.GoIdent.GoImportPath.String()) + "." + v.Input.GoIdent.GoName
				//responseType := getParamPKG(v.Output.GoIdent.GoImportPath.String()) + "." + v.Output.GoIdent.GoName

				// get import path + pkg,
				requestType := method.Input.Desc.Name()
				responseType := method.Output.Desc.Name()

				mInfo := pkg.LazyProxyMethodInfo{
					ServiceName:  serviceName,
					MethodName:   string(method.Desc.Name()),
					RequestName:  string(requestType),
					ResponseName: string(responseType),
				}

				methodInformation = append(methodInformation, mInfo)
			}
		}
	}

	gf := gen.NewGeneratedFile("lazyproxy/main.go", "")
	gf.P("package main")

	for _, service := range serviceInfo {
		str := ExecuteTemplate(pkg.LazyProxyService, service)
		gf.P(str)
	}

	for _, method := range methodInformation {
		str := ExecuteTemplate(pkg.LazyProxyService, method)
		gf.P(str)
	}

	return nil
}

// ExecuteTemplate something to implement templates.
func ExecuteTemplate(tplate string, data any) string {
	// todo read more about template library, see if it may be better to have one Template struct and re use it.
	templ, err := template.New("").Parse(tplate)
	if err != nil {
		panic(err)
	}

	buffy := bytes.NewBuffer([]byte{})
	if err := templ.Execute(buffy, data); err != nil {
		panic(err)
	}
	return buffy.String()
}
