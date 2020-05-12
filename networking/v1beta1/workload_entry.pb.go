// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1beta1/workload_entry.proto

// `WorkloadEntry` enables operators to describe the properties of a
// single non-Kubernetes workload such as a VM or a bare metal server
// as it is onboarded into the mesh. A `WorkloadEntry` must be
// accompanied by an Istio `ServiceEntry` that selects the workload
// through the appropriate labels and provides the service definition
// for a `MESH_INTERNAL` service (hostnames, port properties, etc.). A
// `ServiceEntry` object can select multiple workload entries as well
// as Kubernetes pods based on the label selector specified in the
// service entry.
//
// When a workload connects to `istiod`, the status field in the
// custom resource will be updated to indicate the health of the
// workload along with other details, similar to how Kubernetes
// updates the status of a pod.
//
// The following example declares a workload entry representing a
// VM for the `details.bookinfo.com` service. This VM has
// sidecar installed and bootstrapped using the `details-legacy`
// service account. The sidecar receives HTTP traffic on port 80
// (wrapped in istio mutual TLS) and forwards it to the application on
// the localhost on the same port.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: 2.2.2.2
//   labels:
//     app: details-legacy
//     instance-id: vm1
//   # ports if not specified will be the same as service ports
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: 2.2.2.2
//   labels:
//     app: details-legacy
//     instance-id: vm1
//   # ports if not specified will be the same as service ports
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated service entry
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: STATIC
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
// {{</tabset>}}
//
//
// The following example declares the same VM workload using
// its fully qualified DNS name. The service entry's resolution
// mode should be changed to DNS to indicate that the client-side
// sidecars should dynamically resolve the DNS name at runtime before
// forwarding the request.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: vm1.vpc01.corp.net
//   labels:
//     app: details-legacy
//     instance-id: vm1
//   # ports if not specified will be the same as service ports
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: WorkloadEntry
// metadata:
//   name: details-svc
// spec:
//   # use of the service account indicates that the workload has a
//   # sidecar proxy bootstrapped with this service account. Pods with
//   # sidecars will automatically communicate with the workload using
//   # istio mutual TLS.
//   serviceAccount: details-legacy
//   address: vm1.vpc01.corp.net
//   labels:
//     app: details-legacy
//     instance-id: vm1
//   # ports if not specified will be the same as service ports
// ```
// {{</tab>}}
// {{</tabset>}}
//
// and the associated service entry
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
//
// {{<tab name="v1beta1" category-value="v1beta1">}}
// ```yaml
// apiVersion: networking.istio.io/v1beta1
// kind: ServiceEntry
// metadata:
//   name: details-svc
// spec:
//   hosts:
//   - details.bookinfo.com
//   location: MESH_INTERNAL
//   ports:
//   - number: 80
//     name: http
//     protocol: HTTP
//   resolution: DNS
//   workloadSelector:
//     labels:
//       app: details-legacy
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// WorkloadEntry enables specifying the properties of a single non-Kubernetes workload such a VM or a bare metal services that can be referred to by service entries.
//
// <!-- crd generation tags
// +cue-gen:WorkloadEntry:groupName:networking.istio.io
// +cue-gen:WorkloadEntry:version:v1beta1
// +cue-gen:WorkloadEntry:annotations:helm.sh/resource-policy=keep
// +cue-gen:WorkloadEntry:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:WorkloadEntry:subresource:status
// +cue-gen:WorkloadEntry:scope:Namespaced
// +cue-gen:WorkloadEntry:resource:categories=istio-io,networking-istio-io,shortNames=we,plural=workloadentries
// +cue-gen:WorkloadEntry:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:WorkloadEntry:printerColumn:name=Address,type=string,JSONPath=.spec.address,description="Address associated with the network endpoint."
// +cue-gen:WorkloadEntry:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1beta1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WorkloadEntry struct {
	// Address associated with the network endpoint without the
	// port.  Domain names can be used if and only if the resolution is set
	// to DNS, and must be fully-qualified without wildcards. Use the form
	// unix:///absolute/path/to/socket for Unix domain socket endpoints.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Set of ports associated with the endpoint. The ports must be
	// associated with a port name that was declared as part of the
	// service. Do not use for `unix://` addresses.
	Ports map[string]uint32 `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// One or more labels associated with the endpoint.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Network enables Istio to group endpoints resident in the same L3
	// domain/network. All endpoints in the same network are assumed to be
	// directly reachable from one another. When endpoints in different
	// networks cannot reach each other directly, an Istio Gateway can be
	// used to establish connectivity (usually using the
	// `AUTO_PASSTHROUGH` mode in a Gateway Server). This is
	// an advanced configuration used typically for spanning an Istio mesh
	// over multiple clusters.
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// The locality associated with the endpoint. A locality corresponds
	// to a failure domain (e.g., country/region/zone). Arbitrary failure
	// domain hierarchies can be represented by separating each
	// encapsulating failure domain by /. For example, the locality of an
	// an endpoint in US, in US-East-1 region, within availability zone
	// az-1, in data center rack r11 can be represented as
	// us/us-east-1/az-1/r11. Istio will configure the sidecar to route to
	// endpoints within the same locality as the sidecar. If none of the
	// endpoints in the locality are available, endpoints parent locality
	// (but within the same network ID) will be chosen. For example, if
	// there are two endpoints in same network (networkID "n1"), say e1
	// with locality us/us-east-1/az-1/r11 and e2 with locality
	// us/us-east-1/az-2/r12, a sidecar from us/us-east-1/az-1/r11 locality
	// will prefer e1 from the same locality over e2 from a different
	// locality. Endpoint e2 could be the IP associated with a gateway
	// (that bridges networks n1 and n2), or the IP associated with a
	// standard service endpoint.
	Locality string `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// The load balancing weight associated with the endpoint. Endpoints
	// with higher weights will receive proportionally higher traffic.
	Weight uint32 `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	// The service account associated with the workload if a sidecar
	// is present in the workload. The service account must be present
	// in the same namespace as the configuration ( WorkloadEntry or a
	// ServiceEntry)
	ServiceAccount       string   `protobuf:"bytes,7,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkloadEntry) Reset()         { *m = WorkloadEntry{} }
func (m *WorkloadEntry) String() string { return proto.CompactTextString(m) }
func (*WorkloadEntry) ProtoMessage()    {}
func (*WorkloadEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f01b2f412f1f06, []int{0}
}
func (m *WorkloadEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkloadEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkloadEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkloadEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadEntry.Merge(m, src)
}
func (m *WorkloadEntry) XXX_Size() int {
	return m.Size()
}
func (m *WorkloadEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadEntry.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadEntry proto.InternalMessageInfo

func (m *WorkloadEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WorkloadEntry) GetPorts() map[string]uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *WorkloadEntry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *WorkloadEntry) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *WorkloadEntry) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *WorkloadEntry) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *WorkloadEntry) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkloadEntry)(nil), "istio.networking.v1beta1.WorkloadEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.networking.v1beta1.WorkloadEntry.LabelsEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "istio.networking.v1beta1.WorkloadEntry.PortsEntry")
}

func init() {
	proto.RegisterFile("networking/v1beta1/workload_entry.proto", fileDescriptor_82f01b2f412f1f06)
}

var fileDescriptor_82f01b2f412f1f06 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0xf2, 0x40,
	0x10, 0xc7, 0x49, 0xf2, 0x19, 0x3f, 0x47, 0xfc, 0xbe, 0xb2, 0x94, 0xb2, 0x04, 0xaa, 0xd2, 0x8b,
	0x5e, 0x9a, 0xa0, 0x5e, 0x6c, 0x6f, 0x0a, 0x85, 0x42, 0x7b, 0x28, 0xb9, 0x14, 0x7a, 0x91, 0x4d,
	0xb2, 0x8d, 0x8b, 0x4b, 0x56, 0x36, 0x6b, 0xc4, 0xc7, 0xea, 0x5b, 0xf4, 0xd8, 0x47, 0x10, 0x9f,
	0xa4, 0x64, 0x37, 0x56, 0x4b, 0x29, 0xed, 0x6d, 0xfe, 0xb3, 0xf3, 0xff, 0xed, 0xce, 0xcc, 0x42,
	0x2f, 0xa3, 0x6a, 0x2d, 0xe4, 0x82, 0x65, 0x69, 0x50, 0x0c, 0x22, 0xaa, 0xc8, 0x20, 0x28, 0x35,
	0x17, 0x24, 0x99, 0xd1, 0x4c, 0xc9, 0x8d, 0xbf, 0x94, 0x42, 0x09, 0x84, 0x59, 0xae, 0x98, 0xf0,
	0x0f, 0xe5, 0x7e, 0x55, 0xee, 0x75, 0x52, 0x21, 0x52, 0x4e, 0x03, 0xb2, 0x64, 0xc1, 0x33, 0xa3,
	0x3c, 0x99, 0x45, 0x74, 0x4e, 0x0a, 0x26, 0xa4, 0xb1, 0x5e, 0xbc, 0x38, 0xd0, 0x7a, 0xac, 0x98,
	0x37, 0x25, 0x12, 0x9d, 0x43, 0x9d, 0x24, 0x89, 0xa4, 0x79, 0x8e, 0xad, 0xae, 0xd5, 0x6f, 0x4c,
	0x9d, 0xed, 0xc4, 0x0e, 0xf7, 0x39, 0x74, 0x0b, 0xb5, 0xa5, 0x90, 0x2a, 0xc7, 0x76, 0xd7, 0xe9,
	0x37, 0x87, 0x43, 0xff, 0xbb, 0xbb, 0xfd, 0x4f, 0x58, 0xff, 0xa1, 0x34, 0xe9, 0x30, 0x34, 0x00,
	0x74, 0x07, 0x2e, 0x27, 0x11, 0xe5, 0x39, 0x76, 0x34, 0x6a, 0xf4, 0x5b, 0xd4, 0xbd, 0x76, 0x19,
	0x56, 0x85, 0x40, 0x18, 0xea, 0x95, 0x0f, 0xff, 0x29, 0x5f, 0x1d, 0xee, 0x25, 0xf2, 0xe0, 0x2f,
	0x17, 0x31, 0xe1, 0x4c, 0x6d, 0x70, 0x4d, 0x1f, 0x7d, 0x68, 0x74, 0x06, 0xee, 0x9a, 0xb2, 0x74,
	0xae, 0xb0, 0xdb, 0xb5, 0xfa, 0xad, 0xb0, 0x52, 0xa8, 0x07, 0xff, 0x73, 0x2a, 0x0b, 0x16, 0xd3,
	0x19, 0x89, 0x63, 0xb1, 0xca, 0x14, 0xae, 0x6b, 0xeb, 0xbf, 0x2a, 0x3d, 0x31, 0x59, 0x6f, 0x0c,
	0x70, 0x68, 0x0c, 0x9d, 0x80, 0xb3, 0xa0, 0x1b, 0x33, 0xb6, 0xb0, 0x0c, 0xd1, 0x29, 0xd4, 0x0a,
	0xc2, 0x57, 0x14, 0xdb, 0x9a, 0x6f, 0xc4, 0xb5, 0x3d, 0xb6, 0xbc, 0x2b, 0x68, 0x1e, 0xf5, 0xf1,
	0x93, 0xb5, 0x71, 0x64, 0x9d, 0x5e, 0xbe, 0xee, 0xda, 0xd6, 0xdb, 0xae, 0x6d, 0x6d, 0x77, 0x6d,
	0xeb, 0xa9, 0x63, 0xa6, 0xc6, 0x84, 0x5e, 0xf2, 0xd7, 0x2f, 0x13, 0xb9, 0x7a, 0xd3, 0xa3, 0xf7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x43, 0x53, 0x79, 0xcb, 0x4f, 0x02, 0x00, 0x00,
}

func (m *WorkloadEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkloadEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkloadEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ServiceAccount) > 0 {
		i -= len(m.ServiceAccount)
		copy(dAtA[i:], m.ServiceAccount)
		i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(m.ServiceAccount)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Weight != 0 {
		i = encodeVarintWorkloadEntry(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Locality) > 0 {
		i -= len(m.Locality)
		copy(dAtA[i:], m.Locality)
		i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(m.Locality)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWorkloadEntry(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Ports) > 0 {
		for k := range m.Ports {
			v := m.Ports[k]
			baseI := i
			i = encodeVarintWorkloadEntry(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWorkloadEntry(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWorkloadEntry(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkloadEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkloadEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkloadEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWorkloadEntry(uint64(l))
	}
	if len(m.Ports) > 0 {
		for k, v := range m.Ports {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWorkloadEntry(uint64(len(k))) + 1 + sovWorkloadEntry(uint64(v))
			n += mapEntrySize + 1 + sovWorkloadEntry(uint64(mapEntrySize))
		}
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWorkloadEntry(uint64(len(k))) + 1 + len(v) + sovWorkloadEntry(uint64(len(v)))
			n += mapEntrySize + 1 + sovWorkloadEntry(uint64(mapEntrySize))
		}
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovWorkloadEntry(uint64(l))
	}
	l = len(m.Locality)
	if l > 0 {
		n += 1 + l + sovWorkloadEntry(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovWorkloadEntry(uint64(m.Weight))
	}
	l = len(m.ServiceAccount)
	if l > 0 {
		n += 1 + l + sovWorkloadEntry(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkloadEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkloadEntry(x uint64) (n int) {
	return sovWorkloadEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkloadEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkloadEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkloadEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkloadEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ports", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ports == nil {
				m.Ports = make(map[string]uint32)
			}
			var mapkey string
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorkloadEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWorkloadEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Ports[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorkloadEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWorkloadEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWorkloadEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Locality", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Locality = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkloadEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkloadEntry
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorkloadEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkloadEntry
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkloadEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWorkloadEntry
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWorkloadEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkloadEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWorkloadEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWorkloadEntry
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWorkloadEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkloadEntry   = fmt.Errorf("proto: integer overflow")
)
