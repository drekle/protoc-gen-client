package option

import (
	"fmt"

	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"google.golang.org/protobuf/types/descriptorpb"
)

type LocationMessage struct {
	Location        *descriptor.SourceCodeInfo_Location
	Message         *descriptor.DescriptorProto
	LeadingComments []string
}

type Procedure struct {
	*descriptorpb.MethodDescriptorProto
	InputStreaming  bool
	InputMessage    *descriptorpb.DescriptorProto
	OutputStreaming bool
	OutputMessage   *descriptorpb.DescriptorProto
}

func (p *Procedure) IsValid() bool {
	return p.InputMessage != nil && p.OutputMessage != nil
}

type Service struct {
	*descriptorpb.ServiceDescriptorProto
	Procedures []*Procedure
	Proto      *descriptorpb.FileDescriptorProto
}

func FindMessageType(protos []*descriptorpb.FileDescriptorProto, name string) *descriptorpb.DescriptorProto {
	for _, proto := range protos {
		desc := proto.GetSourceCodeInfo()
		locations := desc.GetLocation()
		for _, location := range locations {
			// I would encourage developers to read the documentation about paths as I might have misunderstood this
			// I am trying to process message types which I understand to be `4` and only at the root level which I understand
			// to be path len == 2
			if len(location.GetPath()) > 2 {
				continue
			}

			if len(location.GetPath()) > 1 && location.GetPath()[0] == int32(4) {
				message := proto.GetMessageType()[location.GetPath()[1]]
				fqn := fmt.Sprintf(".%s.%s", proto.GetPackage(), message.GetName())
				if fqn == name {
					return message
				}
			}
		}
	}
	return nil
}

/*
GetServices will map the required information to make a single command in cobra. This will require the service itself as we will create a client to this service, a procedure which reflects the cobra command, and the input and output message types to be sent to and recieved by the service respectively.
Appended service procedures in the returned list may still be invalid.
*/
func GetServices(req *plugin.CodeGeneratorRequest, proto *descriptorpb.FileDescriptorProto) []*Service {

	ret := make([]*Service, 0)
	for _, service := range proto.GetService() {
		s := &Service{
			ServiceDescriptorProto: service,
			Procedures:             make([]*Procedure, len(service.GetMethod())),
			Proto:                  proto,
		}
		for i, method := range service.GetMethod() {
			procedure := &Procedure{
				MethodDescriptorProto: method,
				InputStreaming:        method.GetClientStreaming(),
				InputMessage:          FindMessageType(req.ProtoFile, method.GetInputType()),
				OutputStreaming:       method.GetServerStreaming(),
				OutputMessage:         FindMessageType(req.ProtoFile, method.GetOutputType()),
			}
			s.Procedures[i] = procedure
		}
		ret = append(ret, s)
	}
	return ret
}
