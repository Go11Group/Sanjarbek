// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: weather.proto

package weatherService

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

type Place struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Place) Reset() {
	*x = Place{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Place) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Place) ProtoMessage() {}

func (x *Place) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Place.ProtoReflect.Descriptor instead.
func (*Place) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{0}
}

func (x *Place) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Forecast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Day  string `protobuf:"bytes,2,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Forecast) Reset() {
	*x = Forecast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forecast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forecast) ProtoMessage() {}

func (x *Forecast) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forecast.ProtoReflect.Descriptor instead.
func (*Forecast) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{1}
}

func (x *Forecast) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Forecast) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

type CurrentWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature int32   `protobuf:"varint,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    float32 `protobuf:"fixed32,2,opt,name=humidity,proto3" json:"humidity,omitempty"`
	WindSpeed   int32   `protobuf:"varint,3,opt,name=windSpeed,proto3" json:"windSpeed,omitempty"`
}

func (x *CurrentWeatherResponse) Reset() {
	*x = CurrentWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentWeatherResponse) ProtoMessage() {}

func (x *CurrentWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentWeatherResponse.ProtoReflect.Descriptor instead.
func (*CurrentWeatherResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{2}
}

func (x *CurrentWeatherResponse) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *CurrentWeatherResponse) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *CurrentWeatherResponse) GetWindSpeed() int32 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

type WeatherForecastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature []int32   `protobuf:"varint,1,rep,packed,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    []float32 `protobuf:"fixed32,2,rep,packed,name=humidity,proto3" json:"humidity,omitempty"`
	WindSpeed   []int32   `protobuf:"varint,3,rep,packed,name=windSpeed,proto3" json:"windSpeed,omitempty"`
	Date        int32     `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *WeatherForecastResponse) Reset() {
	*x = WeatherForecastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherForecastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherForecastResponse) ProtoMessage() {}

func (x *WeatherForecastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherForecastResponse.ProtoReflect.Descriptor instead.
func (*WeatherForecastResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{3}
}

func (x *WeatherForecastResponse) GetTemperature() []int32 {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *WeatherForecastResponse) GetHumidity() []float32 {
	if x != nil {
		return x.Humidity
	}
	return nil
}

func (x *WeatherForecastResponse) GetWindSpeed() []int32 {
	if x != nil {
		return x.WindSpeed
	}
	return nil
}

func (x *WeatherForecastResponse) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

type WeatherConditionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place       *Place  `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
	Temperature int32   `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    float32 `protobuf:"fixed32,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	WindSpeed   int32   `protobuf:"varint,4,opt,name=windSpeed,proto3" json:"windSpeed,omitempty"`
	Condition   string  `protobuf:"bytes,5,opt,name=condition,proto3" json:"condition,omitempty"`
	Date        int32   `protobuf:"varint,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *WeatherConditionResponse) Reset() {
	*x = WeatherConditionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_weather_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeatherConditionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherConditionResponse) ProtoMessage() {}

func (x *WeatherConditionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_weather_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherConditionResponse.ProtoReflect.Descriptor instead.
func (*WeatherConditionResponse) Descriptor() ([]byte, []int) {
	return file_weather_proto_rawDescGZIP(), []int{4}
}

func (x *WeatherConditionResponse) GetPlace() *Place {
	if x != nil {
		return x.Place
	}
	return nil
}

func (x *WeatherConditionResponse) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *WeatherConditionResponse) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *WeatherConditionResponse) GetWindSpeed() int32 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

func (x *WeatherConditionResponse) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

func (x *WeatherConditionResponse) GetDate() int32 {
	if x != nil {
		return x.Date
	}
	return 0
}

var File_weather_proto protoreflect.FileDescriptor

var file_weather_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x1b, 0x0a, 0x05, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x08,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x74,
	0x0a, 0x16, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x17, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x09, 0x77, 0x69, 0x6e, 0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x22, 0xd5, 0x01, 0x0a, 0x18, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65,
	0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x69, 0x6e, 0x64,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x77, 0x69, 0x6e,
	0x64, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32, 0x9e, 0x02, 0x0a, 0x0e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x15, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x26, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46,
	0x6f, 0x72, 0x65, 0x63, 0x61, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x65, 0x63, 0x61,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x16,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x28, 0x2e,
	0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_weather_proto_rawDescOnce sync.Once
	file_weather_proto_rawDescData = file_weather_proto_rawDesc
)

func file_weather_proto_rawDescGZIP() []byte {
	file_weather_proto_rawDescOnce.Do(func() {
		file_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_weather_proto_rawDescData)
	})
	return file_weather_proto_rawDescData
}

var file_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_weather_proto_goTypes = []any{
	(*Place)(nil),                    // 0: weatherService.Place
	(*Forecast)(nil),                 // 1: weatherService.Forecast
	(*CurrentWeatherResponse)(nil),   // 2: weatherService.CurrentWeatherResponse
	(*WeatherForecastResponse)(nil),  // 3: weatherService.WeatherForecastResponse
	(*WeatherConditionResponse)(nil), // 4: weatherService.WeatherConditionResponse
}
var file_weather_proto_depIdxs = []int32{
	0, // 0: weatherService.WeatherConditionResponse.place:type_name -> weatherService.Place
	0, // 1: weatherService.WeatherService.GetCurrentWeather:input_type -> weatherService.Place
	1, // 2: weatherService.WeatherService.GetWeatherForecast:input_type -> weatherService.Forecast
	0, // 3: weatherService.WeatherService.ReportWeatherCondition:input_type -> weatherService.Place
	2, // 4: weatherService.WeatherService.GetCurrentWeather:output_type -> weatherService.CurrentWeatherResponse
	3, // 5: weatherService.WeatherService.GetWeatherForecast:output_type -> weatherService.WeatherForecastResponse
	4, // 6: weatherService.WeatherService.ReportWeatherCondition:output_type -> weatherService.WeatherConditionResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_weather_proto_init() }
func file_weather_proto_init() {
	if File_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_weather_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Place); i {
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
		file_weather_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Forecast); i {
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
		file_weather_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CurrentWeatherResponse); i {
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
		file_weather_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*WeatherForecastResponse); i {
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
		file_weather_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WeatherConditionResponse); i {
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
			RawDescriptor: file_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_weather_proto_goTypes,
		DependencyIndexes: file_weather_proto_depIdxs,
		MessageInfos:      file_weather_proto_msgTypes,
	}.Build()
	File_weather_proto = out.File
	file_weather_proto_rawDesc = nil
	file_weather_proto_goTypes = nil
	file_weather_proto_depIdxs = nil
}
