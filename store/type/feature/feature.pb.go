// Code generated by protoc-gen-go from "src/github.com/MG-RAST/Shock/store/type/feature/feature.proto"
// DO NOT EDIT!

package feature

import proto "code.google.com/p/goprotobuf/proto"
import "math"

// Reference proto and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Index struct {
	I                *uint32 `protobuf:"varint,1,req,name=i" json:"i,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (this *Index) Reset()         { *this = Index{} }
func (this *Index) String() string { return proto.CompactTextString(this) }
func (*Index) ProtoMessage()       {}

func (this *Index) GetI() uint32 {
	if this != nil && this.I != nil {
		return *this.I
	}
	return 0
}

type Feature struct {
	Name             *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Index            []*Index `protobuf:"bytes,2,rep,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (this *Feature) Reset()         { *this = Feature{} }
func (this *Feature) String() string { return proto.CompactTextString(this) }
func (*Feature) ProtoMessage()       {}

func (this *Feature) GetName() string {
	if this != nil && this.Name != nil {
		return *this.Name
	}
	return ""
}

type FeatureList struct {
	Features         []*Feature `protobuf:"bytes,1,rep,name=features" json:"features,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (this *FeatureList) Reset()         { *this = FeatureList{} }
func (this *FeatureList) String() string { return proto.CompactTextString(this) }
func (*FeatureList) ProtoMessage()       {}

func init() {
}
