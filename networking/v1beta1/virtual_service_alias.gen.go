// Code generated by protoc-gen-alias. DO NOT EDIT.
package v1beta1

import "istio.io/api/networking/v1alpha3"

// Configuration affecting traffic routing.
//
// <!-- crd generation tags
// +cue-gen:VirtualService:groupName:networking.istio.io
// +cue-gen:VirtualService:versions:v1beta1,v1alpha3,v1
// +cue-gen:VirtualService:annotations:helm.sh/resource-policy=keep
// +cue-gen:VirtualService:labels:app=istio-pilot,chart=istio,heritage=Tiller,release=istio
// +cue-gen:VirtualService:subresource:status
// +cue-gen:VirtualService:scope:Namespaced
// +cue-gen:VirtualService:resource:categories=istio-io,networking-istio-io,shortNames=vs
// +cue-gen:VirtualService:printerColumn:name=Gateways,type=string,JSONPath=.spec.gateways,description="The names of gateways and sidecars
// that should apply these routes"
// +cue-gen:VirtualService:printerColumn:name=Hosts,type=string,JSONPath=.spec.hosts,description="The destination hosts to which traffic is being sent"
// +cue-gen:VirtualService:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// +cue-gen:VirtualService:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type VirtualService = v1alpha3.VirtualService

// Destination indicates the network addressable service to which the
// request/connection will be sent after processing a routing rule. The
// destination.host should unambiguously refer to a service in the service
// registry. Istio's service registry is composed of all the services found
// in the platform's service registry (e.g., Kubernetes services, Consul
// services), as well as services declared through the
// [ServiceEntry](https://istio.io/docs/reference/config/networking/service-entry/#ServiceEntry) resource.
//
// *Note for Kubernetes users*: When short names are used (e.g. "reviews"
// instead of "reviews.default.svc.cluster.local"), Istio will interpret
// the short name based on the namespace of the rule, not the service. A
// rule in the "default" namespace containing a host "reviews" will be
// interpreted as "reviews.default.svc.cluster.local", irrespective of the
// actual namespace associated with the reviews service. _To avoid potential
// misconfigurations, it is recommended to always use fully qualified
// domain names over short names._
//
// The following Kubernetes example routes all traffic by default to pods
// of the reviews service with label "version: v1" (i.e., subset v1), and
// some to subset v2, in a Kubernetes environment.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: reviews-route
//	namespace: foo
//
// spec:
//
//	hosts:
//	- reviews # interpreted as reviews.foo.svc.cluster.local
//	http:
//	- match:
//	  - uri:
//	      prefix: "/wpcatalog"
//	  - uri:
//	      prefix: "/consumercatalog"
//	  rewrite:
//	    uri: "/newcatalog"
//	  route:
//	  - destination:
//	      host: reviews # interpreted as reviews.foo.svc.cluster.local
//	      subset: v2
//	- route:
//	  - destination:
//	      host: reviews # interpreted as reviews.foo.svc.cluster.local
//	      subset: v1
//
// ```
//
// # And the associated DestinationRule
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: reviews-destination
//	namespace: foo
//
// spec:
//
//	host: reviews # interpreted as reviews.foo.svc.cluster.local
//	subsets:
//	- name: v1
//	  labels:
//	    version: v1
//	- name: v2
//	  labels:
//	    version: v2
//
// ```
//
// The following VirtualService sets a timeout of 5s for all calls to
// productpage.prod.svc.cluster.local service in Kubernetes. Notice that
// there are no subsets defined in this rule. Istio will fetch all
// instances of productpage.prod.svc.cluster.local service from the service
// registry and populate the sidecar's load balancing pool. Also, notice
// that this rule is set in the istio-system namespace but uses the fully
// qualified domain name of the productpage service,
// productpage.prod.svc.cluster.local. Therefore the rule's namespace does
// not have an impact in resolving the name of the productpage service.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: my-productpage-rule
//	namespace: istio-system
//
// spec:
//
//	hosts:
//	- productpage.prod.svc.cluster.local # ignores rule namespace
//	http:
//	- timeout: 5s
//	  route:
//	  - destination:
//	      host: productpage.prod.svc.cluster.local
//
// ```
//
// To control routing for traffic bound to services outside the mesh, external
// services must first be added to Istio's internal service registry using the
// ServiceEntry resource. VirtualServices can then be defined to control traffic
// bound to these external services. For example, the following rules define a
// Service for wikipedia.org and set a timeout of 5s for HTTP requests.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: ServiceEntry
// metadata:
//
//	name: external-svc-wikipedia
//
// spec:
//
//	hosts:
//	- wikipedia.org
//	location: MESH_EXTERNAL
//	ports:
//	- number: 80
//	  name: example-http
//	  protocol: HTTP
//	resolution: DNS
//
// ---
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: my-wiki-rule
//
// spec:
//
//	hosts:
//	- wikipedia.org
//	http:
//	- timeout: 5s
//	  route:
//	  - destination:
//	      host: wikipedia.org
//
// ```
type Destination = v1alpha3.Destination

// Describes match conditions and actions for routing HTTP/1.1, HTTP2, and
// gRPC traffic. See VirtualService for usage examples.
type HTTPRoute = v1alpha3.HTTPRoute

// Describes the delegate VirtualService.
// The following routing rules forward the traffic to `/productpage` by a delegate VirtualService named `productpage`,
// forward the traffic to `/reviews` by a delegate VirtualService named `reviews`.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: bookinfo
//
// spec:
//
//	hosts:
//	- "bookinfo.com"
//	gateways:
//	- mygateway
//	http:
//	- match:
//	  - uri:
//	      prefix: "/productpage"
//	  delegate:
//	     name: productpage
//	     namespace: nsA
//	- match:
//	  - uri:
//	      prefix: "/reviews"
//	  delegate:
//	      name: reviews
//	      namespace: nsB
//
// ```
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: productpage
//	namespace: nsA
//
// spec:
//
//	http:
//	- match:
//	   - uri:
//	      prefix: "/productpage/v1/"
//	  route:
//	  - destination:
//	      host: productpage-v1.nsA.svc.cluster.local
//	- route:
//	  - destination:
//	      host: productpage.nsA.svc.cluster.local
//
// ```
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: reviews
//	namespace: nsB
//
// spec:
//
//	http:
//	- route:
//	  - destination:
//	      host: reviews.nsB.svc.cluster.local
//
// ```
type Delegate = v1alpha3.Delegate

// Message headers can be manipulated when Envoy forwards requests to,
// or responses from, a destination service. Header manipulation rules can
// be specified for a specific route destination or for all destinations.
// The following VirtualService adds a `test` header with the value `true`
// to requests that are routed to any `reviews` service destination.
// It also removes the `foo` response header, but only from responses
// coming from the `v1` subset (version) of the `reviews` service.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: reviews-route
//
// spec:
//
//	hosts:
//	- reviews.prod.svc.cluster.local
//	http:
//	- headers:
//	    request:
//	      set:
//	        test: "true"
//	  route:
//	  - destination:
//	      host: reviews.prod.svc.cluster.local
//	      subset: v2
//	    weight: 25
//	  - destination:
//	      host: reviews.prod.svc.cluster.local
//	      subset: v1
//	    headers:
//	      response:
//	        remove:
//	        - foo
//	    weight: 75
//
// ```
type Headers = v1alpha3.Headers

// HeaderOperations Describes the header manipulations to apply
type Headers_HeaderOperations = v1alpha3.Headers_HeaderOperations

// Describes match conditions and actions for routing unterminated TLS
// traffic (TLS/HTTPS) The following routing rule forwards unterminated TLS
// traffic arriving at port 443 of gateway called "mygateway" to internal
// services in the mesh based on the SNI value.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: bookinfo-sni
//
// spec:
//
//	hosts:
//	- "*.bookinfo.com"
//	gateways:
//	- mygateway
//	tls:
//	- match:
//	  - port: 443
//	    sniHosts:
//	    - login.bookinfo.com
//	  route:
//	  - destination:
//	      host: login.prod.svc.cluster.local
//	- match:
//	  - port: 443
//	    sniHosts:
//	    - reviews.bookinfo.com
//	  route:
//	  - destination:
//	      host: reviews.prod.svc.cluster.local
//
// ```
type TLSRoute = v1alpha3.TLSRoute

// Describes match conditions and actions for routing TCP traffic. The
// following routing rule forwards traffic arriving at port 27017 for
// mongo.prod.svc.cluster.local to another Mongo server on port 5555.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: bookinfo-mongo
//
// spec:
//
//	hosts:
//	- mongo.prod.svc.cluster.local
//	tcp:
//	- match:
//	  - port: 27017
//	  route:
//	  - destination:
//	      host: mongo.backup.svc.cluster.local
//	      port:
//	        number: 5555
//
// ```
type TCPRoute = v1alpha3.TCPRoute

// HttpMatchRequest specifies a set of criteria to be met in order for the
// rule to be applied to the HTTP request. For example, the following
// restricts the rule to match only requests where the URL path
// starts with /ratings/v2/ and the request contains a custom `end-user` header
// with value `jason`.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- match:
//	  - headers:
//	      end-user:
//	        exact: jason
//	    uri:
//	      prefix: "/ratings/v2/"
//	    ignoreUriCase: true
//	  route:
//	  - destination:
//	      host: ratings.prod.svc.cluster.local
//
// ```
//
// HTTPMatchRequest CANNOT be empty.
// **Note:**
// 1. If a root VirtualService have matched any property (path, header etc.) by regex, delegate VirtualServices should not have any other matches on the same property.
// 2. If a delegate VirtualService have matched any property (path, header etc.) by regex, root VirtualServices should not have any other matches on the same property.
type HTTPMatchRequest = v1alpha3.HTTPMatchRequest

// Each routing rule is associated with one or more service versions (see
// glossary in beginning of document). Weights associated with the version
// determine the proportion of traffic it receives. For example, the
// following rule will route 25% of traffic for the "reviews" service to
// instances with the "v2" tag and the remaining traffic (i.e., 75%) to
// "v1".
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: reviews-route
//
// spec:
//
//	hosts:
//	- reviews.prod.svc.cluster.local
//	http:
//	- route:
//	  - destination:
//	      host: reviews.prod.svc.cluster.local
//	      subset: v2
//	    weight: 25
//	  - destination:
//	      host: reviews.prod.svc.cluster.local
//	      subset: v1
//	    weight: 75
//
// ```
//
// # And the associated DestinationRule
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: DestinationRule
// metadata:
//
//	name: reviews-destination
//
// spec:
//
//	host: reviews.prod.svc.cluster.local
//	subsets:
//	- name: v1
//	  labels:
//	    version: v1
//	- name: v2
//	  labels:
//	    version: v2
//
// ```
//
// Traffic can also be split across two entirely different services without
// having to define new subsets. For example, the following rule forwards 25% of
// traffic to reviews.com to dev.reviews.com
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: reviews-route-two-domains
//
// spec:
//
//	hosts:
//	- reviews.com
//	http:
//	- route:
//	  - destination:
//	      host: dev.reviews.com
//	    weight: 25
//	  - destination:
//	      host: reviews.com
//	    weight: 75
//
// ```
type HTTPRouteDestination = v1alpha3.HTTPRouteDestination

// L4 routing rule weighted destination.
type RouteDestination = v1alpha3.RouteDestination

// L4 connection match attributes. Note that L4 connection matching support
// is incomplete.
type L4MatchAttributes = v1alpha3.L4MatchAttributes

// TLS connection match attributes.
type TLSMatchAttributes = v1alpha3.TLSMatchAttributes

// HTTPRedirect can be used to send a 301 redirect response to the caller,
// where the Authority/Host and the URI in the response can be swapped with
// the specified values. For example, the following rule redirects
// requests for /v1/getProductRatings API on the ratings service to
// /v1/bookRatings provided by the bookratings service.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- match:
//	  - uri:
//	      exact: /v1/getProductRatings
//	  redirect:
//	    uri: /v1/bookRatings
//	    authority: newratings.default.svc.cluster.local
//	...
//
// ```
type HTTPRedirect = v1alpha3.HTTPRedirect
type HTTPRedirect_RedirectPortSelection = v1alpha3.HTTPRedirect_RedirectPortSelection

const HTTPRedirect_FROM_PROTOCOL_DEFAULT HTTPRedirect_RedirectPortSelection = v1alpha3.HTTPRedirect_FROM_PROTOCOL_DEFAULT
const HTTPRedirect_FROM_REQUEST_PORT HTTPRedirect_RedirectPortSelection = v1alpha3.HTTPRedirect_FROM_REQUEST_PORT

// On a redirect, overwrite the port portion of the URL with this value.
type HTTPRedirect_Port = v1alpha3.HTTPRedirect_Port

// On a redirect, dynamically set the port:
// * FROM_PROTOCOL_DEFAULT: automatically set to 80 for HTTP and 443 for HTTPS.
// * FROM_REQUEST_PORT: automatically use the port of the request.
type HTTPRedirect_DerivePort = v1alpha3.HTTPRedirect_DerivePort

// HTTPDirectResponse can be used to send a fixed response to clients.
// For example, the following rule returns a fixed 503 status with a body
// to requests for /v1/getProductRatings API.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- match:
//	  - uri:
//	      exact: /v1/getProductRatings
//	  directResponse:
//	    status: 503
//	    body:
//	      string: "unknown error"
//	...
//
// ```
//
// It is also possible to specify a binary response body.
// This is mostly useful for non text-based protocols such as gRPC.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- match:
//	  - uri:
//	      exact: /v1/getProductRatings
//	  directResponse:
//	    status: 503
//	    body:
//	      bytes: "dW5rbm93biBlcnJvcg==" # "unknown error" in base64
//	...
//
// ```
//
// It is good practice to add headers in the HTTPRoute
// as well as the direct_response, for example to specify
// the returned Content-Type.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- match:
//	  - uri:
//	      exact: /v1/getProductRatings
//	  directResponse:
//	    status: 503
//	    body:
//	      string: "{\"error\": \"unknown error\"}"
//	  headers:
//	    response:
//	      set:
//	        content-type: "text/plain"
//	...
//
// ```
type HTTPDirectResponse = v1alpha3.HTTPDirectResponse
type HTTPBody = v1alpha3.HTTPBody

// response body as a string
type HTTPBody_String_ = v1alpha3.HTTPBody_String_

// response body as base64 encoded bytes.
type HTTPBody_Bytes = v1alpha3.HTTPBody_Bytes

// HTTPRewrite can be used to rewrite specific parts of a HTTP request
// before forwarding the request to the destination. Rewrite primitive can
// be used only with HTTPRouteDestination. The following example
// demonstrates how to rewrite the URL prefix for api call (/ratings) to
// ratings service before making the actual API call.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- match:
//	  - uri:
//	      prefix: /ratings
//	  rewrite:
//	    uri: /v1/bookRatings
//	  route:
//	  - destination:
//	      host: ratings.prod.svc.cluster.local
//	      subset: v1
//
// ```
type HTTPRewrite = v1alpha3.HTTPRewrite
type RegexRewrite = v1alpha3.RegexRewrite

// Describes how to match a given string in HTTP headers. Match is
// case-sensitive.
type StringMatch = v1alpha3.StringMatch

// exact string match
type StringMatch_Exact = v1alpha3.StringMatch_Exact

// prefix-based match
type StringMatch_Prefix = v1alpha3.StringMatch_Prefix

// RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).
type StringMatch_Regex = v1alpha3.StringMatch_Regex

// Describes the retry policy to use when a HTTP request fails. For
// example, the following rule sets the maximum number of retries to 3 when
// calling ratings:v1 service, with a 2s timeout per retry attempt.
// A retry will be attempted if there is a connect-failure, refused_stream
// or when the upstream server responds with Service Unavailable(503).
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- route:
//	  - destination:
//	      host: ratings.prod.svc.cluster.local
//	      subset: v1
//	  retries:
//	    attempts: 3
//	    perTryTimeout: 2s
//	    retryOn: gateway-error,connect-failure,refused-stream
//
// ```
type HTTPRetry = v1alpha3.HTTPRetry

// Describes the Cross-Origin Resource Sharing (CORS) policy, for a given
// service. Refer to [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/Access_control_CORS)
// for further details about cross origin resource sharing. For example,
// the following rule restricts cross origin requests to those originating
// from example.com domain using HTTP POST/GET, and sets the
// `Access-Control-Allow-Credentials` header to false. In addition, it only
// exposes `X-Foo-bar` header and sets an expiry period of 1 day.
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- route:
//	  - destination:
//	      host: ratings.prod.svc.cluster.local
//	      subset: v1
//	  corsPolicy:
//	    allowOrigins:
//	    - exact: https://example.com
//	    allowMethods:
//	    - POST
//	    - GET
//	    allowCredentials: false
//	    allowHeaders:
//	    - X-Foo-Bar
//	    maxAge: "24h"
//
// ```
type CorsPolicy = v1alpha3.CorsPolicy
type CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_UnmatchedPreflights

// Default to FORWARD
const CorsPolicy_UNSPECIFIED CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_UNSPECIFIED

// Preflight requests not matching the configured allowed origin
// will be forwarded to the upstream.
const CorsPolicy_FORWARD CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_FORWARD

// Preflight requests not matching the configured allowed origin
// will not be forwarded to the upstream.
const CorsPolicy_IGNORE CorsPolicy_UnmatchedPreflights = v1alpha3.CorsPolicy_IGNORE

// HTTPFaultInjection can be used to specify one or more faults to inject
// while forwarding HTTP requests to the destination specified in a route.
// Fault specification is part of a VirtualService rule. Faults include
// aborting the Http request from downstream service, and/or delaying
// proxying of requests. A fault rule MUST HAVE delay or abort or both.
//
// *Note:* Delay and abort faults are independent of one another, even if
// both are specified simultaneously.
type HTTPFaultInjection = v1alpha3.HTTPFaultInjection

// Delay specification is used to inject latency into the request
// forwarding path. The following example will introduce a 5 second delay
// in 1 out of every 1000 requests to the "v1" version of the "reviews"
// service from all pods with label env: prod
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: reviews-route
//
// spec:
//
//	hosts:
//	- reviews.prod.svc.cluster.local
//	http:
//	- match:
//	  - sourceLabels:
//	      env: prod
//	  route:
//	  - destination:
//	      host: reviews.prod.svc.cluster.local
//	      subset: v1
//	  fault:
//	    delay:
//	      percentage:
//	        value: 0.1
//	      fixedDelay: 5s
//
// ```
//
// The _fixedDelay_ field is used to indicate the amount of delay in seconds.
// The optional _percentage_ field can be used to only delay a certain
// percentage of requests. If left unspecified, no request will be delayed.
type HTTPFaultInjection_Delay = v1alpha3.HTTPFaultInjection_Delay

// Add a fixed delay before forwarding the request. Format:
// 1h/1m/1s/1ms. MUST be >=1ms.
type HTTPFaultInjection_Delay_FixedDelay = v1alpha3.HTTPFaultInjection_Delay_FixedDelay

// $hide_from_docs
type HTTPFaultInjection_Delay_ExponentialDelay = v1alpha3.HTTPFaultInjection_Delay_ExponentialDelay

// Abort specification is used to prematurely abort a request with a
// pre-specified error code. The following example will return an HTTP 400
// error code for 1 out of every 1000 requests to the "ratings" service "v1".
//
// ```yaml
// apiVersion: networking.istio.io/v1
// kind: VirtualService
// metadata:
//
//	name: ratings-route
//
// spec:
//
//	hosts:
//	- ratings.prod.svc.cluster.local
//	http:
//	- route:
//	  - destination:
//	      host: ratings.prod.svc.cluster.local
//	      subset: v1
//	  fault:
//	    abort:
//	      percentage:
//	        value: 0.1
//	      httpStatus: 400
//
// ```
//
// The _httpStatus_ field is used to indicate the HTTP status code to
// return to the caller. The optional _percentage_ field can be used to only
// abort a certain percentage of requests. If not specified, no request will be
// aborted.
type HTTPFaultInjection_Abort = v1alpha3.HTTPFaultInjection_Abort

// HTTP status code to use to abort the Http request.
type HTTPFaultInjection_Abort_HttpStatus = v1alpha3.HTTPFaultInjection_Abort_HttpStatus

// GRPC status code to use to abort the request. The supported
// codes are documented in https://github.com/grpc/grpc/blob/master/doc/statuscodes.md
// Note: If you want to return the status "Unavailable", then you should
// specify the code as `UNAVAILABLE`(all caps), but not `14`.
type HTTPFaultInjection_Abort_GrpcStatus = v1alpha3.HTTPFaultInjection_Abort_GrpcStatus

// $hide_from_docs
type HTTPFaultInjection_Abort_Http2Error = v1alpha3.HTTPFaultInjection_Abort_Http2Error

// HTTPMirrorPolicy can be used to specify the destinations to mirror HTTP traffic in addition
// to the original destination. Mirrored traffic is on a
// best effort basis where the sidecar/gateway will not wait for the
// mirrored destinations to respond before returning the response from the
// original destination. Statistics will be generated for the mirrored
// destination.
type HTTPMirrorPolicy = v1alpha3.HTTPMirrorPolicy

// PortSelector specifies the number of a port to be used for
// matching or selection for final routing.
type PortSelector = v1alpha3.PortSelector

// Percent specifies a percentage in the range of [0.0, 100.0].
type Percent = v1alpha3.Percent
