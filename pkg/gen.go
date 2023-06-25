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
		for _, service := range file.Services {

			pkgName := getParamPKG(file.GoDescriptorIdent.GoImportPath.String())
			connectPkgName := "/" + pkgName + "connect"
			protoIdent := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: file.GoDescriptorIdent.GoImportPath})
			gf.P("// proto ident " + protoIdent)

			connectIdent := gf.QualifiedGoIdent(protogen.GoIdent{GoImportPath: file.GoDescriptorIdent.GoImportPath + protogen.GoImportPath(connectPkgName)})
			gf.P("// connectImport ident " + connectIdent)

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

				reqIdent := gf.QualifiedGoIdent(method.Input.GoIdent)
				resIdent := gf.QualifiedGoIdent(method.Output.GoIdent)
				gf.P("// req ident " + reqIdent)
				gf.P("// res ident " + resIdent)

				//imports[method.Output.GoIdent.GoImportPath.String()] = method.Output.GoIdent.GoImportPath.String()
				requestType := reqIdent  //+ "." + string(method.Input.Desc.Name())
				responseType := resIdent //+ "." + string(method.Output.Desc.Name())

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
