# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: networking/v1beta1/sidecar.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import field_behavior_pb2 as google_dot_api_dot_field__behavior__pb2
from networking.v1beta1 import gateway_pb2 as networking_dot_v1beta1_dot_gateway__pb2
from networking.v1beta1 import virtual_service_pb2 as networking_dot_v1beta1_dot_virtual__service__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='networking/v1beta1/sidecar.proto',
  package='istio.networking.v1beta1',
  syntax='proto3',
  serialized_options=_b('Z\037istio.io/api/networking/v1beta1'),
  serialized_pb=_b('\n networking/v1beta1/sidecar.proto\x12\x18istio.networking.v1beta1\x1a\x1fgoogle/api/field_behavior.proto\x1a networking/v1beta1/gateway.proto\x1a(networking/v1beta1/virtual_service.proto\"\xf3\x02\n\x07Sidecar\x12W\n\x11workload_selector\x18\x01 \x01(\x0b\x32*.istio.networking.v1beta1.WorkloadSelectorR\x10workloadSelector\x12H\n\x07ingress\x18\x02 \x03(\x0b\x32..istio.networking.v1beta1.IstioIngressListenerR\x07ingress\x12\x45\n\x06\x65gress\x18\x03 \x03(\x0b\x32-.istio.networking.v1beta1.IstioEgressListenerR\x06\x65gress\x12g\n\x17outbound_traffic_policy\x18\x04 \x01(\x0b\x32/.istio.networking.v1beta1.OutboundTrafficPolicyR\x15outboundTrafficPolicyJ\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07R\tlocalhost\"\x81\x02\n\x14IstioIngressListener\x12\x38\n\x04port\x18\x01 \x01(\x0b\x32\x1e.istio.networking.v1beta1.PortB\x04\xe2\x41\x01\x02R\x04port\x12\x12\n\x04\x62ind\x18\x02 \x01(\tR\x04\x62ind\x12H\n\x0c\x63\x61pture_mode\x18\x03 \x01(\x0e\x32%.istio.networking.v1beta1.CaptureModeR\x0b\x63\x61ptureMode\x12/\n\x10\x64\x65\x66\x61ult_endpoint\x18\x04 \x01(\tB\x04\xe2\x41\x01\x02R\x0f\x64\x65\x66\x61ultEndpointJ\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07R\x14localhost_client_tls\"\xe5\x01\n\x13IstioEgressListener\x12\x32\n\x04port\x18\x01 \x01(\x0b\x32\x1e.istio.networking.v1beta1.PortR\x04port\x12\x12\n\x04\x62ind\x18\x02 \x01(\tR\x04\x62ind\x12H\n\x0c\x63\x61pture_mode\x18\x03 \x01(\x0e\x32%.istio.networking.v1beta1.CaptureModeR\x0b\x63\x61ptureMode\x12\x1a\n\x05hosts\x18\x04 \x03(\tB\x04\xe2\x41\x01\x02R\x05hostsJ\x04\x08\x05\x10\x06J\x04\x08\x06\x10\x07R\x14localhost_server_tls\"\xa3\x01\n\x10WorkloadSelector\x12T\n\x06labels\x18\x01 \x03(\x0b\x32\x36.istio.networking.v1beta1.WorkloadSelector.LabelsEntryB\x04\xe2\x41\x01\x02R\x06labels\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\xd5\x01\n\x15OutboundTrafficPolicy\x12H\n\x04mode\x18\x01 \x01(\x0e\x32\x34.istio.networking.v1beta1.OutboundTrafficPolicy.ModeR\x04mode\x12H\n\x0c\x65gress_proxy\x18\x02 \x01(\x0b\x32%.istio.networking.v1beta1.DestinationR\x0b\x65gressProxy\"(\n\x04Mode\x12\x11\n\rREGISTRY_ONLY\x10\x00\x12\r\n\tALLOW_ANY\x10\x01*2\n\x0b\x43\x61ptureMode\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x0c\n\x08IPTABLES\x10\x01\x12\x08\n\x04NONE\x10\x02\x42!Z\x1fistio.io/api/networking/v1beta1b\x06proto3')
  ,
  dependencies=[google_dot_api_dot_field__behavior__pb2.DESCRIPTOR,networking_dot_v1beta1_dot_gateway__pb2.DESCRIPTOR,networking_dot_v1beta1_dot_virtual__service__pb2.DESCRIPTOR,])

_CAPTUREMODE = _descriptor.EnumDescriptor(
  name='CaptureMode',
  full_name='istio.networking.v1beta1.CaptureMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DEFAULT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='IPTABLES', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NONE', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1419,
  serialized_end=1469,
)
_sym_db.RegisterEnumDescriptor(_CAPTUREMODE)

CaptureMode = enum_type_wrapper.EnumTypeWrapper(_CAPTUREMODE)
DEFAULT = 0
IPTABLES = 1
NONE = 2


_OUTBOUNDTRAFFICPOLICY_MODE = _descriptor.EnumDescriptor(
  name='Mode',
  full_name='istio.networking.v1beta1.OutboundTrafficPolicy.Mode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='REGISTRY_ONLY', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ALLOW_ANY', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1377,
  serialized_end=1417,
)
_sym_db.RegisterEnumDescriptor(_OUTBOUNDTRAFFICPOLICY_MODE)


_SIDECAR = _descriptor.Descriptor(
  name='Sidecar',
  full_name='istio.networking.v1beta1.Sidecar',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='workload_selector', full_name='istio.networking.v1beta1.Sidecar.workload_selector', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='workloadSelector', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress', full_name='istio.networking.v1beta1.Sidecar.ingress', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='ingress', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='egress', full_name='istio.networking.v1beta1.Sidecar.egress', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='egress', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outbound_traffic_policy', full_name='istio.networking.v1beta1.Sidecar.outbound_traffic_policy', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='outboundTrafficPolicy', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=172,
  serialized_end=543,
)


_ISTIOINGRESSLISTENER = _descriptor.Descriptor(
  name='IstioIngressListener',
  full_name='istio.networking.v1beta1.IstioIngressListener',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='port', full_name='istio.networking.v1beta1.IstioIngressListener.port', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342A\001\002'), json_name='port', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bind', full_name='istio.networking.v1beta1.IstioIngressListener.bind', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='bind', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='capture_mode', full_name='istio.networking.v1beta1.IstioIngressListener.capture_mode', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='captureMode', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_endpoint', full_name='istio.networking.v1beta1.IstioIngressListener.default_endpoint', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342A\001\002'), json_name='defaultEndpoint', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=546,
  serialized_end=803,
)


_ISTIOEGRESSLISTENER = _descriptor.Descriptor(
  name='IstioEgressListener',
  full_name='istio.networking.v1beta1.IstioEgressListener',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='port', full_name='istio.networking.v1beta1.IstioEgressListener.port', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='port', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='bind', full_name='istio.networking.v1beta1.IstioEgressListener.bind', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='bind', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='capture_mode', full_name='istio.networking.v1beta1.IstioEgressListener.capture_mode', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='captureMode', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hosts', full_name='istio.networking.v1beta1.IstioEgressListener.hosts', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342A\001\002'), json_name='hosts', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=806,
  serialized_end=1035,
)


_WORKLOADSELECTOR_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='istio.networking.v1beta1.WorkloadSelector.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='istio.networking.v1beta1.WorkloadSelector.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='key', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='istio.networking.v1beta1.WorkloadSelector.LabelsEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1144,
  serialized_end=1201,
)

_WORKLOADSELECTOR = _descriptor.Descriptor(
  name='WorkloadSelector',
  full_name='istio.networking.v1beta1.WorkloadSelector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='labels', full_name='istio.networking.v1beta1.WorkloadSelector.labels', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\342A\001\002'), json_name='labels', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_WORKLOADSELECTOR_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1038,
  serialized_end=1201,
)


_OUTBOUNDTRAFFICPOLICY = _descriptor.Descriptor(
  name='OutboundTrafficPolicy',
  full_name='istio.networking.v1beta1.OutboundTrafficPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.networking.v1beta1.OutboundTrafficPolicy.mode', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='mode', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='egress_proxy', full_name='istio.networking.v1beta1.OutboundTrafficPolicy.egress_proxy', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='egressProxy', file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _OUTBOUNDTRAFFICPOLICY_MODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1204,
  serialized_end=1417,
)

_SIDECAR.fields_by_name['workload_selector'].message_type = _WORKLOADSELECTOR
_SIDECAR.fields_by_name['ingress'].message_type = _ISTIOINGRESSLISTENER
_SIDECAR.fields_by_name['egress'].message_type = _ISTIOEGRESSLISTENER
_SIDECAR.fields_by_name['outbound_traffic_policy'].message_type = _OUTBOUNDTRAFFICPOLICY
_ISTIOINGRESSLISTENER.fields_by_name['port'].message_type = networking_dot_v1beta1_dot_gateway__pb2._PORT
_ISTIOINGRESSLISTENER.fields_by_name['capture_mode'].enum_type = _CAPTUREMODE
_ISTIOEGRESSLISTENER.fields_by_name['port'].message_type = networking_dot_v1beta1_dot_gateway__pb2._PORT
_ISTIOEGRESSLISTENER.fields_by_name['capture_mode'].enum_type = _CAPTUREMODE
_WORKLOADSELECTOR_LABELSENTRY.containing_type = _WORKLOADSELECTOR
_WORKLOADSELECTOR.fields_by_name['labels'].message_type = _WORKLOADSELECTOR_LABELSENTRY
_OUTBOUNDTRAFFICPOLICY.fields_by_name['mode'].enum_type = _OUTBOUNDTRAFFICPOLICY_MODE
_OUTBOUNDTRAFFICPOLICY.fields_by_name['egress_proxy'].message_type = networking_dot_v1beta1_dot_virtual__service__pb2._DESTINATION
_OUTBOUNDTRAFFICPOLICY_MODE.containing_type = _OUTBOUNDTRAFFICPOLICY
DESCRIPTOR.message_types_by_name['Sidecar'] = _SIDECAR
DESCRIPTOR.message_types_by_name['IstioIngressListener'] = _ISTIOINGRESSLISTENER
DESCRIPTOR.message_types_by_name['IstioEgressListener'] = _ISTIOEGRESSLISTENER
DESCRIPTOR.message_types_by_name['WorkloadSelector'] = _WORKLOADSELECTOR
DESCRIPTOR.message_types_by_name['OutboundTrafficPolicy'] = _OUTBOUNDTRAFFICPOLICY
DESCRIPTOR.enum_types_by_name['CaptureMode'] = _CAPTUREMODE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Sidecar = _reflection.GeneratedProtocolMessageType('Sidecar', (_message.Message,), {
  'DESCRIPTOR' : _SIDECAR,
  '__module__' : 'networking.v1beta1.sidecar_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.Sidecar)
  })
_sym_db.RegisterMessage(Sidecar)

IstioIngressListener = _reflection.GeneratedProtocolMessageType('IstioIngressListener', (_message.Message,), {
  'DESCRIPTOR' : _ISTIOINGRESSLISTENER,
  '__module__' : 'networking.v1beta1.sidecar_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.IstioIngressListener)
  })
_sym_db.RegisterMessage(IstioIngressListener)

IstioEgressListener = _reflection.GeneratedProtocolMessageType('IstioEgressListener', (_message.Message,), {
  'DESCRIPTOR' : _ISTIOEGRESSLISTENER,
  '__module__' : 'networking.v1beta1.sidecar_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.IstioEgressListener)
  })
_sym_db.RegisterMessage(IstioEgressListener)

WorkloadSelector = _reflection.GeneratedProtocolMessageType('WorkloadSelector', (_message.Message,), {

  'LabelsEntry' : _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), {
    'DESCRIPTOR' : _WORKLOADSELECTOR_LABELSENTRY,
    '__module__' : 'networking.v1beta1.sidecar_pb2'
    # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.WorkloadSelector.LabelsEntry)
    })
  ,
  'DESCRIPTOR' : _WORKLOADSELECTOR,
  '__module__' : 'networking.v1beta1.sidecar_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.WorkloadSelector)
  })
_sym_db.RegisterMessage(WorkloadSelector)
_sym_db.RegisterMessage(WorkloadSelector.LabelsEntry)

OutboundTrafficPolicy = _reflection.GeneratedProtocolMessageType('OutboundTrafficPolicy', (_message.Message,), {
  'DESCRIPTOR' : _OUTBOUNDTRAFFICPOLICY,
  '__module__' : 'networking.v1beta1.sidecar_pb2'
  # @@protoc_insertion_point(class_scope:istio.networking.v1beta1.OutboundTrafficPolicy)
  })
_sym_db.RegisterMessage(OutboundTrafficPolicy)


DESCRIPTOR._options = None
_ISTIOINGRESSLISTENER.fields_by_name['port']._options = None
_ISTIOINGRESSLISTENER.fields_by_name['default_endpoint']._options = None
_ISTIOEGRESSLISTENER.fields_by_name['hosts']._options = None
_WORKLOADSELECTOR_LABELSENTRY._options = None
_WORKLOADSELECTOR.fields_by_name['labels']._options = None
# @@protoc_insertion_point(module_scope)
