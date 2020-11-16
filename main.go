package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/proto"

	"github.com/drekle/protoc-gen-client/pkg/option"
	protochelpers "github.com/drekle/protoc-gen-client/pkg/option"
	"github.com/drekle/protoc-gen-client/pkg/template"
)

const (
	Version = "0.1.0"
)

type Config struct {
	Version   string   `json:"version"`
	GitRepo   string   `json:"gitRepository"`
	SkipFiles []string `json:"skipFiles"`
}

func GenConfig() ([]byte, error) {
	return json.MarshalIndent(&Config{
		Version: Version,
	}, " ", "\t")
}

type clientGen struct {
	Request    *plugin.CodeGeneratorRequest
	Response   *plugin.CodeGeneratorResponse
	Parameters map[string]string
	Config     Config
}

func (runner *clientGen) PrintParameters(w io.Writer) {
	const padding = 3
	tw := tabwriter.NewWriter(w, 0, 0, padding, ' ', tabwriter.TabIndent)
	fmt.Fprintf(tw, "Parameters:\n")
	for k, v := range runner.Parameters {
		fmt.Fprintf(tw, "%s:\t%s\n", k, v)
	}
	fmt.Fprintln(tw, "")
	tw.Flush()
}

func (runner *clientGen) GenerateCobra(
	cliOpt *protochelpers.ClientOptions,
	cmdOpt *protochelpers.CommandOptions,
	svc *protochelpers.Service,
	procedure *protochelpers.Procedure) error {

	goPackage := strings.ReplaceAll(svc.Proto.GetPackage(), ".", "/")
	outfileName := fmt.Sprintf("cmd/%s/%s/%s.go", goPackage, *svc.Name, strings.ToLower(*procedure.Name))
	var buf bytes.Buffer
	t := template.NewTemplate()
	ex, err := t.Parse(template.CobraCommandTemplate)
	if err != nil {
		return err
	}
	err = ex.Execute(&buf, &template.CobraCommandInput{
		Service:   svc,
		GoPBPath:  fmt.Sprintf("gen/%s", goPackage),
		RepoURL:   cliOpt.Repository,
		Procedure: procedure,
	})
	if err != nil {
		return err
	}
	var outFile plugin.CodeGeneratorResponse_File
	outFile.Name = &outfileName
	content := buf.String()
	outFile.Content = &content
	runner.Response.File = append(runner.Response.File, &outFile)

	return nil
}

func (runner *clientGen) generateCode() error {
	// Initialize the output file slice
	files := make([]*plugin.CodeGeneratorResponse_File, 0)
	runner.Response.File = files
	{
		for _, protofile := range runner.Request.ProtoFile {

			var clientOpt *protochelpers.ClientOptions
			if proto.HasExtension(protofile.Options, protochelpers.E_ClientOptions) {
				e := proto.GetExtension(protofile.Options, protochelpers.E_ClientOptions)
				if cmd, ok := e.(*option.ClientOptions); ok {
					clientOpt = cmd
				}
			} else {
				continue
			}
			// Create a file and append it to the output files
			for _, svc := range protochelpers.GetServices(runner.Request, protofile) {
				for _, procedure := range svc.Procedures {
					if proto.HasExtension(procedure.Options, protochelpers.E_CommandOptions) {
						e := proto.GetExtension(procedure.Options, protochelpers.E_CommandOptions)
						if cmdOpt, ok := e.(*protochelpers.CommandOptions); ok {
							err := runner.GenerateCobra(clientOpt, cmdOpt, svc, procedure)
							if err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func main() {

	stdInFile := flag.String("stdinFile", "", "A file to use for stdin")
	flag.Parse()

	// os.Stdin will contain data which will unmarshal into the following object:
	// https://godoc.org/github.com/golang/protobuf/protoc-gen-go/plugin#CodeGeneratorRequest
	req := &plugin.CodeGeneratorRequest{}
	resp := &plugin.CodeGeneratorResponse{}

	// Debugging Support
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	if *stdInFile != "" {
		data, err = ioutil.ReadFile(*stdInFile)
		if err != nil {
			panic(err)
		}
	}

	if err := proto.Unmarshal(data, req); err != nil {
		panic(err)
	}

	// You may require more data than what is in the proto files alone.  There are a couple ways in which to do this.
	// The first is by parameters.  Another may be using leading comments in the proto files which I will cover in generateCode.
	parameters := req.GetParameter()
	// =grpc,import_path=mypackage:.
	clientRunner := &clientGen{
		Request:    req,
		Response:   resp,
		Parameters: make(map[string]string),
	}
	groupkv := strings.Split(parameters, ",")
	for _, element := range groupkv {
		kv := strings.Split(element, "=")
		if len(kv) > 1 {
			clientRunner.Parameters[kv[0]] = kv[1]
		}
	}

	err = clientRunner.generateCode()
	if err != nil {
		panic(err)
	}

	marshalled, err := proto.Marshal(resp)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(marshalled)
}
