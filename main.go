package main

import (
	"bytes"
	"flag"
	"sort"
	"strings"
	"text/template"

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
		return Generate(gen)
	})
}

func Generate(gen *protogen.Plugin) error {

	serviceInfo := make([]pkg.LazyProxyServiceInfo, 0)
	methodInformation := make([]pkg.LazyProxyMethodInfo, 0)
	imports := map[string]string{}
	gf := gen.NewGeneratedFile("lazyproxy/main.go", protogen.GoImportPath("."))
	gf.P("package main")

	for _, file := range gen.Files {
		pkgName := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())
		connectPkgName := "/" + pkgName + "connect"
		connectImport := protogen.GoIdent{GoImportPath: file.GoDescriptorIdent.GoImportPath + protogen.GoImportPath(connectPkgName)}.GoImportPath.String()
		protoImport := protogen.GoIdent{GoImportPath: file.GoDescriptorIdent.GoImportPath}.GoImportPath.String()

		imports[connectImport] = connectImport
		imports[protoImport] = protoImport

		for _, service := range file.Services {
			serviceName := string(service.Desc.Name())
			sInfo := pkg.LazyProxyServiceInfo{
				ServiceName: serviceName,
				Pkg:         pkgName,
			}
			serviceInfo = append(serviceInfo, sInfo)

			for _, method := range service.Methods {
				// todo have this condition not add info.
				if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
					break
				}

				// Note need to import req / res if it is different
				requestType := getParamPKG(method.Input.GoIdent.GoImportPath.String()) + "." + string(method.Input.Desc.Name())
				responseType := getParamPKG(method.Output.GoIdent.GoImportPath.String()) + "." + string(method.Output.Desc.Name())

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

	// todo have imports be sorted
	importArr := make([]string, 0, len(imports))
	for _, v := range imports {
		importArr = append(importArr, v)
	}
	sort.SliceStable(importArr, func(i, j int) bool {
		return importArr[i] > importArr[j]
	})

	// todo move to array to import all.
	//gf.P(`import "github.com/bufbuild/connect-go"`)
	//gf.P(`import "google.golang.org/grpc"`)
	//gf.P(`import "context"`)

	serverInfo := pkg.LazyProxyServerInfo{
		Services: serviceInfo,
		Imports:  importArr,
	}

	str := ExecuteTemplate(pkg.LazyProxyServer, serverInfo)
	gf.P(str)

	for _, service := range serviceInfo {
		str := ExecuteTemplate(pkg.LazyProxyService, service)
		gf.P(str)
	}

	for _, method := range methodInformation {
		str := ExecuteTemplate(pkg.LazyProxyMethod, method)
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

func getParamPKG(in string) string {
	arr := strings.Split(in, "/")
	return strings.Trim(arr[len(arr)-1], `"`)
}
