package template

import (
	"fmt"
	"strings"
	"text/template"

	"google.golang.org/protobuf/types/descriptorpb"
)

var TemplateFuncs template.FuncMap

func init() {
	TemplateFuncs = make(map[string]interface{})
	TemplateFuncs["toLower"] = strings.ToLower
	TemplateFuncs["privateName"] = PrivateName
	TemplateFuncs["publicName"] = PublicName
	TemplateFuncs["removeAll"] = RemoveAll
	TemplateFuncs["deref"] = Deref
	TemplateFuncs["goType"] = GoType
}

func NewTemplate() *template.Template {
	return template.New("protoc-gen-cli").Funcs(TemplateFuncs)
}

func PrivateName(s string) string {
	return fmt.Sprintf("%s%s", strings.ToLower(s[0:1]), s[1:])
}

func PublicName(s string) string {
	return fmt.Sprintf("%s%s", strings.ToUpper(s[0:1]), s[1:])
}

func RemoveAll(s string, remove string) string {
	return strings.ReplaceAll(s, remove, "")
}

func Deref(s *string) string {
	return *s
}

// TODO: Revize and remove
func GoType(s *descriptorpb.FieldDescriptorProto_Type) string {
	if s == nil {
		return ""
	}
	switch *s {
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		return "float64"
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		return "float64"
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		return "int64"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		return "uint64"
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		return "int32"
	// case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
	// 	return ""
	// case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
	// 	return ""
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		return "bool"
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		return "string"
	// case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
	// 	return ""
	// case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
	// 	return ""
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		return "[]byte"
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		return "uint32"
	// case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
	// 	return ""
	// case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
	// 	return ""
	// case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
	// 	return ""
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		return "int32"
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		return "int64"
	}
	return ""
}
