package pkg

import (
	"bytes"
	"sort"
	"strings"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

// Generate implements generating the lazyproxy.
func Generate(gen *protogen.Plugin) error {
	serviceInfo := make([]LazyProxyServiceInfo, 0)
	imports := map[string]string{}
	// todo for module based proxy consider having the name of this to be set by caller of plugin or pkg name.
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
			sInfo := LazyProxyServiceInfo{
				ServiceName: serviceName,
				Pkg:         pkgName,
			}

			methodInformation := make([]LazyProxyMethodInfo, 0)
			for _, method := range service.Methods {
				// todo have this condition not add info.
				if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
					break
				}

				// Note need to import req / res if it is different
				requestType := getParamPKG(method.Input.GoIdent.GoImportPath.String()) + "." + string(method.Input.Desc.Name())
				responseType := getParamPKG(method.Output.GoIdent.GoImportPath.String()) + "." + string(method.Output.Desc.Name())

				mInfo := LazyProxyMethodInfo{
					ServiceName:  serviceName,
					MethodName:   string(method.Desc.Name()),
					RequestName:  string(requestType),
					ResponseName: string(responseType),
				}

				methodInformation = append(methodInformation, mInfo)
			}
			sInfo.Methods = methodInformation
			serviceInfo = append(serviceInfo, sInfo)
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

	serverInfo := LazyProxyServerInfo{
		Services: serviceInfo,
		Imports:  importArr,
	}

	str := ExecuteTemplate(LazyProxyServer, serverInfo)
	gf.P(str)

	for _, service := range serviceInfo {
		str := ExecuteTemplate(LazyProxyService, service)
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
