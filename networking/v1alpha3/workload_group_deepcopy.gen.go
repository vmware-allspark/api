// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/workload_group.proto

// `WorkloadGroup` describes a collection of workload instances.
// It provides a specification that the workload instances can use to bootstrap
// their proxies, including the metadata and identity. It is only intended to
// be used with non-k8s workloads like Virtual Machines, and is meant to mimic
// the existing sidecar injection and deployment specification model used for
// Kubernetes workloads to bootstrap Istio proxies.
//
// The following example declares a workload group representing a collection
// of workloads that will be registered under `reviews` in namespace
// `bookinfo`. The set of labels will be associated with each workload
// instance during the bootstrap process, and the ports 3550 and 8080
// will be associated with the workload group and use service account `default`.
// `app.kubernetes.io/version` is just an arbitrary example of a label.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadGroup
// metadata:
//   name: reviews
//   namespace: bookinfo
// spec:
//   metadata:
//     labels:
//       app.kubernetes.io/name: reviews
//       app.kubernetes.io/version: "1.3.4"
//   template:
//     ports:
//       grpc: 3550
//       http: 8080
//     serviceAccount: default
//   probe:
//     initialDelaySeconds: 5
//     timeoutSeconds: 3
//     periodSeconds: 4
//     successThreshold: 3
//     failureThreshold: 3
//     httpGet:
//      path: /foo/bar
//      host: 127.0.0.1
//      port: 3100
//      scheme: https
//      httpHeaders:
//      - name: Lit-Header
//        value: Im-The-Best
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using WorkloadGroup within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadGroup) DeepCopyInto(out *WorkloadGroup) {
	p := proto.Clone(in).(*WorkloadGroup)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGroup. Required by controller-gen.
func (in *WorkloadGroup) DeepCopy() *WorkloadGroup {
	if in == nil {
		return nil
	}
	out := new(WorkloadGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGroup. Required by controller-gen.
func (in *WorkloadGroup) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using WorkloadGroup_ObjectMeta within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadGroup_ObjectMeta) DeepCopyInto(out *WorkloadGroup_ObjectMeta) {
	p := proto.Clone(in).(*WorkloadGroup_ObjectMeta)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGroup_ObjectMeta. Required by controller-gen.
func (in *WorkloadGroup_ObjectMeta) DeepCopy() *WorkloadGroup_ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(WorkloadGroup_ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadGroup_ObjectMeta. Required by controller-gen.
func (in *WorkloadGroup_ObjectMeta) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ReadinessProbe within kubernetes types, where deepcopy-gen is used.
func (in *ReadinessProbe) DeepCopyInto(out *ReadinessProbe) {
	p := proto.Clone(in).(*ReadinessProbe)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbe. Required by controller-gen.
func (in *ReadinessProbe) DeepCopy() *ReadinessProbe {
	if in == nil {
		return nil
	}
	out := new(ReadinessProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbe. Required by controller-gen.
func (in *ReadinessProbe) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPHealthCheckConfig within kubernetes types, where deepcopy-gen is used.
func (in *HTTPHealthCheckConfig) DeepCopyInto(out *HTTPHealthCheckConfig) {
	p := proto.Clone(in).(*HTTPHealthCheckConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHealthCheckConfig. Required by controller-gen.
func (in *HTTPHealthCheckConfig) DeepCopy() *HTTPHealthCheckConfig {
	if in == nil {
		return nil
	}
	out := new(HTTPHealthCheckConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHealthCheckConfig. Required by controller-gen.
func (in *HTTPHealthCheckConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using HTTPHeader within kubernetes types, where deepcopy-gen is used.
func (in *HTTPHeader) DeepCopyInto(out *HTTPHeader) {
	p := proto.Clone(in).(*HTTPHeader)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader. Required by controller-gen.
func (in *HTTPHeader) DeepCopy() *HTTPHeader {
	if in == nil {
		return nil
	}
	out := new(HTTPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new HTTPHeader. Required by controller-gen.
func (in *HTTPHeader) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using TCPHealthCheckConfig within kubernetes types, where deepcopy-gen is used.
func (in *TCPHealthCheckConfig) DeepCopyInto(out *TCPHealthCheckConfig) {
	p := proto.Clone(in).(*TCPHealthCheckConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPHealthCheckConfig. Required by controller-gen.
func (in *TCPHealthCheckConfig) DeepCopy() *TCPHealthCheckConfig {
	if in == nil {
		return nil
	}
	out := new(TCPHealthCheckConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new TCPHealthCheckConfig. Required by controller-gen.
func (in *TCPHealthCheckConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ExecHealthCheckConfig within kubernetes types, where deepcopy-gen is used.
func (in *ExecHealthCheckConfig) DeepCopyInto(out *ExecHealthCheckConfig) {
	p := proto.Clone(in).(*ExecHealthCheckConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecHealthCheckConfig. Required by controller-gen.
func (in *ExecHealthCheckConfig) DeepCopy() *ExecHealthCheckConfig {
	if in == nil {
		return nil
	}
	out := new(ExecHealthCheckConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ExecHealthCheckConfig. Required by controller-gen.
func (in *ExecHealthCheckConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
