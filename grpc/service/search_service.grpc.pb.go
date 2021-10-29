// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: protos/service/search_service.grpc.proto

package service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SearchJobAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search      string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	Filter      string `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
	Page        int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Count       int32  `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	OrderBy     string `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	JobSeekerId string `protobuf:"bytes,6,opt,name=job_seeker_id,json=jobSeekerId,proto3" json:"job_seeker_id,omitempty"`
}

func (x *SearchJobAlertRequest) Reset() {
	*x = SearchJobAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_search_service_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchJobAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchJobAlertRequest) ProtoMessage() {}

func (x *SearchJobAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_search_service_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchJobAlertRequest.ProtoReflect.Descriptor instead.
func (*SearchJobAlertRequest) Descriptor() ([]byte, []int) {
	return file_protos_service_search_service_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *SearchJobAlertRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *SearchJobAlertRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

func (x *SearchJobAlertRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchJobAlertRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchJobAlertRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

func (x *SearchJobAlertRequest) GetJobSeekerId() string {
	if x != nil {
		return x.JobSeekerId
	}
	return ""
}

type SearchJobAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"@odata.count";"
	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// @inject_tag: json:"value";"
	Value []*SearchJobAlert `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *SearchJobAlertResponse) Reset() {
	*x = SearchJobAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_search_service_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchJobAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchJobAlertResponse) ProtoMessage() {}

func (x *SearchJobAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_search_service_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchJobAlertResponse.ProtoReflect.Descriptor instead.
func (*SearchJobAlertResponse) Descriptor() ([]byte, []int) {
	return file_protos_service_search_service_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *SearchJobAlertResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SearchJobAlertResponse) GetValue() []*SearchJobAlert {
	if x != nil {
		return x.Value
	}
	return nil
}

type SearchJobAlert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id";"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// @inject_tag: json:"job_title";"
	JobTitle string `protobuf:"bytes,2,opt,name=job_title,json=jobTitle,proto3" json:"job_title,omitempty"`
	// @inject_tag: json:"province";"
	Province string `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	// @inject_tag: json:"career_level";"
	CareerLevel string `protobuf:"bytes,4,opt,name=career_level,json=careerLevel,proto3" json:"career_level,omitempty"`
	// @inject_tag: json:"employment_type";"
	EmploymentType string `protobuf:"bytes,5,opt,name=employment_type,json=employmentType,proto3" json:"employment_type,omitempty"`
	// @inject_tag: json:"job_classification";"
	JobClassification string `protobuf:"bytes,6,opt,name=job_classification,json=jobClassification,proto3" json:"job_classification,omitempty"`
	// @inject_tag: json:"company";"
	Company *SearchJobEmployer `protobuf:"bytes,7,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *SearchJobAlert) Reset() {
	*x = SearchJobAlert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_search_service_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchJobAlert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchJobAlert) ProtoMessage() {}

func (x *SearchJobAlert) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_search_service_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchJobAlert.ProtoReflect.Descriptor instead.
func (*SearchJobAlert) Descriptor() ([]byte, []int) {
	return file_protos_service_search_service_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *SearchJobAlert) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SearchJobAlert) GetJobTitle() string {
	if x != nil {
		return x.JobTitle
	}
	return ""
}

func (x *SearchJobAlert) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *SearchJobAlert) GetCareerLevel() string {
	if x != nil {
		return x.CareerLevel
	}
	return ""
}

func (x *SearchJobAlert) GetEmploymentType() string {
	if x != nil {
		return x.EmploymentType
	}
	return ""
}

func (x *SearchJobAlert) GetJobClassification() string {
	if x != nil {
		return x.JobClassification
	}
	return ""
}

func (x *SearchJobAlert) GetCompany() *SearchJobEmployer {
	if x != nil {
		return x.Company
	}
	return nil
}

type SearchJobEmployer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"employer_id";"
	EmployerId string `protobuf:"bytes,1,opt,name=employer_id,json=employerId,proto3" json:"employer_id,omitempty"`
	// @inject_tag: json:"employer_name";"
	EmployerName string `protobuf:"bytes,2,opt,name=employer_name,json=employerName,proto3" json:"employer_name,omitempty"`
	// @inject_tag: json:"employer_photo_url";"
	EmployerPhotoUrl string `protobuf:"bytes,3,opt,name=employer_photo_url,json=employerPhotoUrl,proto3" json:"employer_photo_url,omitempty"`
}

func (x *SearchJobEmployer) Reset() {
	*x = SearchJobEmployer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_service_search_service_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchJobEmployer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchJobEmployer) ProtoMessage() {}

func (x *SearchJobEmployer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_service_search_service_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchJobEmployer.ProtoReflect.Descriptor instead.
func (*SearchJobEmployer) Descriptor() ([]byte, []int) {
	return file_protos_service_search_service_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *SearchJobEmployer) GetEmployerId() string {
	if x != nil {
		return x.EmployerId
	}
	return ""
}

func (x *SearchJobEmployer) GetEmployerName() string {
	if x != nil {
		return x.EmployerName
	}
	return ""
}

func (x *SearchJobEmployer) GetEmployerPhotoUrl() string {
	if x != nil {
		return x.EmployerPhotoUrl
	}
	return ""
}

var File_protos_service_search_service_grpc_proto protoreflect.FileDescriptor

var file_protos_service_search_service_grpc_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6a, 0x6f, 0x62,
	0x5f, 0x73, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6a, 0x6f, 0x62, 0x53, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x22, 0x64, 0x0a,
	0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f,
	0x62, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x6f, 0x62, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x65, 0x65, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6a,
	0x6f, 0x62, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6a, 0x6f, 0x62, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x87, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4a, 0x6f, 0x62, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x5f,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72,
	0x6c, 0x32, 0x70, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4a, 0x6f, 0x62, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x61, 0x6e, 0x70, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protos_service_search_service_grpc_proto_rawDescOnce sync.Once
	file_protos_service_search_service_grpc_proto_rawDescData = file_protos_service_search_service_grpc_proto_rawDesc
)

func file_protos_service_search_service_grpc_proto_rawDescGZIP() []byte {
	file_protos_service_search_service_grpc_proto_rawDescOnce.Do(func() {
		file_protos_service_search_service_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_service_search_service_grpc_proto_rawDescData)
	})
	return file_protos_service_search_service_grpc_proto_rawDescData
}

var file_protos_service_search_service_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_service_search_service_grpc_proto_goTypes = []interface{}{
	(*SearchJobAlertRequest)(nil),  // 0: protos.service.SearchJobAlertRequest
	(*SearchJobAlertResponse)(nil), // 1: protos.service.SearchJobAlertResponse
	(*SearchJobAlert)(nil),         // 2: protos.service.SearchJobAlert
	(*SearchJobEmployer)(nil),      // 3: protos.service.SearchJobEmployer
}
var file_protos_service_search_service_grpc_proto_depIdxs = []int32{
	2, // 0: protos.service.SearchJobAlertResponse.value:type_name -> protos.service.SearchJobAlert
	3, // 1: protos.service.SearchJobAlert.company:type_name -> protos.service.SearchJobEmployer
	0, // 2: protos.service.SearchService.SearchJobAlert:input_type -> protos.service.SearchJobAlertRequest
	1, // 3: protos.service.SearchService.SearchJobAlert:output_type -> protos.service.SearchJobAlertResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_service_search_service_grpc_proto_init() }
func file_protos_service_search_service_grpc_proto_init() {
	if File_protos_service_search_service_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_service_search_service_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchJobAlertRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_service_search_service_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchJobAlertResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_service_search_service_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchJobAlert); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_service_search_service_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchJobEmployer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_service_search_service_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_service_search_service_grpc_proto_goTypes,
		DependencyIndexes: file_protos_service_search_service_grpc_proto_depIdxs,
		MessageInfos:      file_protos_service_search_service_grpc_proto_msgTypes,
	}.Build()
	File_protos_service_search_service_grpc_proto = out.File
	file_protos_service_search_service_grpc_proto_rawDesc = nil
	file_protos_service_search_service_grpc_proto_goTypes = nil
	file_protos_service_search_service_grpc_proto_depIdxs = nil
}
