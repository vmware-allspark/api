# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mesh/v1alpha1/config.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from mesh.v1alpha1 import proxy_pb2 as mesh_dot_v1alpha1_dot_proxy__pb2
from networking.v1alpha3 import destination_rule_pb2 as networking_dot_v1alpha3_dot_destination__rule__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mesh/v1alpha1/config.proto',
  package='istio.mesh.v1alpha1',
  syntax='proto3',
  serialized_options=_b('Z\032istio.io/api/mesh/v1alpha1'),
  serialized_pb=_b('\n\x1amesh/v1alpha1/config.proto\x12\x13istio.mesh.v1alpha1\x1a\x1egoogle/protobuf/duration.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x19mesh/v1alpha1/proxy.proto\x1a*networking/v1alpha3/destination_rule.proto\"\xba\x16\n\nMeshConfig\x12\x19\n\x11proxy_listen_port\x18\x04 \x01(\x05\x12\x17\n\x0fproxy_http_port\x18\x05 \x01(\x05\x12\x32\n\x0f\x63onnect_timeout\x18\x06 \x01(\x0b\x32\x19.google.protobuf.Duration\x12=\n\x1aprotocol_detection_timeout\x18* \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x61\n\rtcp_keepalive\x18\x1c \x01(\x0b\x32J.istio.networking.v1alpha3.ConnectionPoolSettings.TCPSettings.TcpKeepalive\x12\x15\n\ringress_class\x18\x07 \x01(\t\x12\x17\n\x0fingress_service\x18\x08 \x01(\t\x12V\n\x17ingress_controller_mode\x18\t \x01(\x0e\x32\x35.istio.mesh.v1alpha1.MeshConfig.IngressControllerMode\x12\x18\n\x10ingress_selector\x18\x34 \x01(\t\x12\x43\n\x0b\x61uth_policy\x18\n \x01(\x0e\x32*.istio.mesh.v1alpha1.MeshConfig.AuthPolicyB\x02\x18\x01\x12\x38\n\x11rds_refresh_delay\x18\x0b \x01(\x0b\x32\x19.google.protobuf.DurationB\x02\x18\x01\x12\x16\n\x0e\x65nable_tracing\x18\x0c \x01(\x08\x12\x17\n\x0f\x61\x63\x63\x65ss_log_file\x18\r \x01(\t\x12\x19\n\x11\x61\x63\x63\x65ss_log_format\x18\x18 \x01(\t\x12N\n\x13\x61\x63\x63\x65ss_log_encoding\x18\x1b \x01(\x0e\x32\x31.istio.mesh.v1alpha1.MeshConfig.AccessLogEncoding\x12\'\n\x1f\x65nable_envoy_access_log_service\x18( \x01(\x08\x12\x38\n\x0e\x64\x65\x66\x61ult_config\x18\x0e \x01(\x0b\x32 .istio.mesh.v1alpha1.ProxyConfig\x12V\n\x17outbound_traffic_policy\x18\x11 \x01(\x0b\x32\x35.istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy\x12\x39\n\x0e\x63onfig_sources\x18\x16 \x03(\x0b\x32!.istio.mesh.v1alpha1.ConfigSource\x12\x34\n\x10\x65nable_auto_mtls\x18+ \x01(\x0b\x32\x1a.google.protobuf.BoolValue\x12\x14\n\x0ctrust_domain\x18\x1a \x01(\t\x12\x1c\n\x14trust_domain_aliases\x18. \x03(\t\x12!\n\x19\x64\x65\x66\x61ult_service_export_to\x18\x1f \x03(\t\x12)\n!default_virtual_service_export_to\x18  \x03(\t\x12*\n\"default_destination_rule_export_to\x18! \x03(\t\x12\x16\n\x0eroot_namespace\x18\" \x01(\t\x12S\n\x13locality_lb_setting\x18# \x01(\x0b\x32\x36.istio.networking.v1alpha3.LocalityLoadBalancerSetting\x12\x33\n\x10\x64ns_refresh_rate\x18$ \x01(\x0b\x32\x19.google.protobuf.Duration\x12J\n\x11h2_upgrade_policy\x18) \x01(\x0e\x32/.istio.mesh.v1alpha1.MeshConfig.H2UpgradePolicy\x12!\n\x19inbound_cluster_stat_name\x18, \x01(\t\x12\"\n\x1aoutbound_cluster_stat_name\x18- \x01(\t\x12\x36\n\x0c\x63\x65rtificates\x18/ \x03(\x0b\x32 .istio.mesh.v1alpha1.Certificate\x12\x43\n\rthrift_config\x18\x31 \x01(\x0b\x32,.istio.mesh.v1alpha1.MeshConfig.ThriftConfig\x12I\n\x10service_settings\x18\x32 \x03(\x0b\x32/.istio.mesh.v1alpha1.MeshConfig.ServiceSettings\x12;\n\x17\x65nable_prometheus_merge\x18\x33 \x01(\x0b\x32\x1a.google.protobuf.BoolValue\x1a\xa7\x01\n\x15OutboundTrafficPolicy\x12H\n\x04mode\x18\x01 \x01(\x0e\x32:.istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.Mode\"D\n\x04Mode\x12\x11\n\rREGISTRY_ONLY\x10\x00\x12\r\n\tALLOW_ANY\x10\x01\"\x04\x08\x02\x10\x02*\x14VIRTUAL_SERVICE_ONLY\x1a]\n\x0cThriftConfig\x12\x16\n\x0erate_limit_url\x18\x01 \x01(\t\x12\x35\n\x12rate_limit_timeout\x18\x02 \x01(\x0b\x32\x19.google.protobuf.Duration\x1a\x8f\x01\n\x0fServiceSettings\x12J\n\x08settings\x18\x01 \x01(\x0b\x32\x38.istio.mesh.v1alpha1.MeshConfig.ServiceSettings.Settings\x12\r\n\x05hosts\x18\x02 \x03(\t\x1a!\n\x08Settings\x12\x15\n\rcluster_local\x18\x01 \x01(\x08\"J\n\x15IngressControllerMode\x12\x0f\n\x0bUNSPECIFIED\x10\x00\x12\x07\n\x03OFF\x10\x01\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x02\x12\n\n\x06STRICT\x10\x03\"&\n\nAuthPolicy\x12\x08\n\x04NONE\x10\x00\x12\x0e\n\nMUTUAL_TLS\x10\x01\"\'\n\x11\x41\x63\x63\x65ssLogEncoding\x12\x08\n\x04TEXT\x10\x00\x12\x08\n\x04JSON\x10\x01\"2\n\x0fH2UpgradePolicy\x12\x12\n\x0e\x44O_NOT_UPGRADE\x10\x00\x12\x0b\n\x07UPGRADE\x10\x01J\x04\x08\x01\x10\x02J\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04J\x04\x08\x30\x10\x31J\x04\x08\x19\x10\x1aJ\x04\x08\x1e\x10\x1fJ\x04\x08\x0f\x10\x10J\x04\x08\x10\x10\x11J\x04\x08\x12\x10\x13J\x04\x08\x13\x10\x14J\x04\x08\x14\x10\x15J\x04\x08\x15\x10\x16J\x04\x08\x17\x10\x18J\x04\x08\x1d\x10\x1eJ\x04\x08%\x10&J\x04\x08&\x10\'J\x04\x08\'\x10(J\x04\x08\x35\x10\x36R\x12mixer_check_serverR\x13mixer_report_serverR\x15\x64isable_policy_checksR\x1a\x64isable_mixer_http_reportsR\x16policy_check_fail_openR%sidecar_to_telemetry_session_affinityR\rmixer_addressR\x1f\x65nable_client_side_policy_checkR\x0csds_uds_pathR\x11sds_refresh_delayR\x16\x65nable_sds_token_mountR\x12sds_use_k8s_sa_jwtR\x14\x64isable_report_batchR\x18report_batch_max_entriesR\x15report_batch_max_timeR\x1atermination_drain_duration\"\xa0\x01\n\x0c\x43onfigSource\x12\x0f\n\x07\x61\x64\x64ress\x18\x01 \x01(\t\x12\x42\n\x0ctls_settings\x18\x02 \x01(\x0b\x32,.istio.networking.v1alpha3.ClientTLSSettings\x12;\n\x14subscribed_resources\x18\x03 \x03(\x0e\x32\x1d.istio.mesh.v1alpha1.Resource\"5\n\x0b\x43\x65rtificate\x12\x13\n\x0bsecret_name\x18\x01 \x01(\t\x12\x11\n\tdns_names\x18\x02 \x03(\t* \n\x08Resource\x12\x14\n\x10SERVICE_REGISTRY\x10\x00\x42\x1cZ\x1aistio.io/api/mesh/v1alpha1b\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_wrappers__pb2.DESCRIPTOR,mesh_dot_v1alpha1_dot_proxy__pb2.DESCRIPTOR,networking_dot_v1alpha3_dot_destination__rule__pb2.DESCRIPTOR,])

_RESOURCE = _descriptor.EnumDescriptor(
  name='Resource',
  full_name='istio.mesh.v1alpha1.Resource',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SERVICE_REGISTRY', index=0, number=0,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=3281,
  serialized_end=3313,
)
_sym_db.RegisterEnumDescriptor(_RESOURCE)

Resource = enum_type_wrapper.EnumTypeWrapper(_RESOURCE)
SERVICE_REGISTRY = 0


_MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE = _descriptor.EnumDescriptor(
  name='Mode',
  full_name='istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.Mode',
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
  serialized_start=2056,
  serialized_end=2124,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE)

_MESHCONFIG_INGRESSCONTROLLERMODE = _descriptor.EnumDescriptor(
  name='IngressControllerMode',
  full_name='istio.mesh.v1alpha1.MeshConfig.IngressControllerMode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNSPECIFIED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OFF', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEFAULT', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STRICT', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2367,
  serialized_end=2441,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_INGRESSCONTROLLERMODE)

_MESHCONFIG_AUTHPOLICY = _descriptor.EnumDescriptor(
  name='AuthPolicy',
  full_name='istio.mesh.v1alpha1.MeshConfig.AuthPolicy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MUTUAL_TLS', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2443,
  serialized_end=2481,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_AUTHPOLICY)

_MESHCONFIG_ACCESSLOGENCODING = _descriptor.EnumDescriptor(
  name='AccessLogEncoding',
  full_name='istio.mesh.v1alpha1.MeshConfig.AccessLogEncoding',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TEXT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JSON', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2483,
  serialized_end=2522,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_ACCESSLOGENCODING)

_MESHCONFIG_H2UPGRADEPOLICY = _descriptor.EnumDescriptor(
  name='H2UpgradePolicy',
  full_name='istio.mesh.v1alpha1.MeshConfig.H2UpgradePolicy',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='DO_NOT_UPGRADE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UPGRADE', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=2524,
  serialized_end=2574,
)
_sym_db.RegisterEnumDescriptor(_MESHCONFIG_H2UPGRADEPOLICY)


_MESHCONFIG_OUTBOUNDTRAFFICPOLICY = _descriptor.Descriptor(
  name='OutboundTrafficPolicy',
  full_name='istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='mode', full_name='istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy.mode', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1957,
  serialized_end=2124,
)

_MESHCONFIG_THRIFTCONFIG = _descriptor.Descriptor(
  name='ThriftConfig',
  full_name='istio.mesh.v1alpha1.MeshConfig.ThriftConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='rate_limit_url', full_name='istio.mesh.v1alpha1.MeshConfig.ThriftConfig.rate_limit_url', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rate_limit_timeout', full_name='istio.mesh.v1alpha1.MeshConfig.ThriftConfig.rate_limit_timeout', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=2126,
  serialized_end=2219,
)

_MESHCONFIG_SERVICESETTINGS_SETTINGS = _descriptor.Descriptor(
  name='Settings',
  full_name='istio.mesh.v1alpha1.MeshConfig.ServiceSettings.Settings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cluster_local', full_name='istio.mesh.v1alpha1.MeshConfig.ServiceSettings.Settings.cluster_local', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=2332,
  serialized_end=2365,
)

_MESHCONFIG_SERVICESETTINGS = _descriptor.Descriptor(
  name='ServiceSettings',
  full_name='istio.mesh.v1alpha1.MeshConfig.ServiceSettings',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='settings', full_name='istio.mesh.v1alpha1.MeshConfig.ServiceSettings.settings', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='hosts', full_name='istio.mesh.v1alpha1.MeshConfig.ServiceSettings.hosts', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_MESHCONFIG_SERVICESETTINGS_SETTINGS, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=2222,
  serialized_end=2365,
)

_MESHCONFIG = _descriptor.Descriptor(
  name='MeshConfig',
  full_name='istio.mesh.v1alpha1.MeshConfig',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='proxy_listen_port', full_name='istio.mesh.v1alpha1.MeshConfig.proxy_listen_port', index=0,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='proxy_http_port', full_name='istio.mesh.v1alpha1.MeshConfig.proxy_http_port', index=1,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='connect_timeout', full_name='istio.mesh.v1alpha1.MeshConfig.connect_timeout', index=2,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='protocol_detection_timeout', full_name='istio.mesh.v1alpha1.MeshConfig.protocol_detection_timeout', index=3,
      number=42, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tcp_keepalive', full_name='istio.mesh.v1alpha1.MeshConfig.tcp_keepalive', index=4,
      number=28, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_class', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_class', index=5,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_service', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_service', index=6,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_controller_mode', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_controller_mode', index=7,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ingress_selector', full_name='istio.mesh.v1alpha1.MeshConfig.ingress_selector', index=8,
      number=52, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='auth_policy', full_name='istio.mesh.v1alpha1.MeshConfig.auth_policy', index=9,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='rds_refresh_delay', full_name='istio.mesh.v1alpha1.MeshConfig.rds_refresh_delay', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=_b('\030\001'), file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enable_tracing', full_name='istio.mesh.v1alpha1.MeshConfig.enable_tracing', index=11,
      number=12, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_log_file', full_name='istio.mesh.v1alpha1.MeshConfig.access_log_file', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_log_format', full_name='istio.mesh.v1alpha1.MeshConfig.access_log_format', index=13,
      number=24, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='access_log_encoding', full_name='istio.mesh.v1alpha1.MeshConfig.access_log_encoding', index=14,
      number=27, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enable_envoy_access_log_service', full_name='istio.mesh.v1alpha1.MeshConfig.enable_envoy_access_log_service', index=15,
      number=40, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_config', full_name='istio.mesh.v1alpha1.MeshConfig.default_config', index=16,
      number=14, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outbound_traffic_policy', full_name='istio.mesh.v1alpha1.MeshConfig.outbound_traffic_policy', index=17,
      number=17, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='config_sources', full_name='istio.mesh.v1alpha1.MeshConfig.config_sources', index=18,
      number=22, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enable_auto_mtls', full_name='istio.mesh.v1alpha1.MeshConfig.enable_auto_mtls', index=19,
      number=43, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='trust_domain', full_name='istio.mesh.v1alpha1.MeshConfig.trust_domain', index=20,
      number=26, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='trust_domain_aliases', full_name='istio.mesh.v1alpha1.MeshConfig.trust_domain_aliases', index=21,
      number=46, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_service_export_to', full_name='istio.mesh.v1alpha1.MeshConfig.default_service_export_to', index=22,
      number=31, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_virtual_service_export_to', full_name='istio.mesh.v1alpha1.MeshConfig.default_virtual_service_export_to', index=23,
      number=32, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='default_destination_rule_export_to', full_name='istio.mesh.v1alpha1.MeshConfig.default_destination_rule_export_to', index=24,
      number=33, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='root_namespace', full_name='istio.mesh.v1alpha1.MeshConfig.root_namespace', index=25,
      number=34, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='locality_lb_setting', full_name='istio.mesh.v1alpha1.MeshConfig.locality_lb_setting', index=26,
      number=35, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dns_refresh_rate', full_name='istio.mesh.v1alpha1.MeshConfig.dns_refresh_rate', index=27,
      number=36, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='h2_upgrade_policy', full_name='istio.mesh.v1alpha1.MeshConfig.h2_upgrade_policy', index=28,
      number=41, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inbound_cluster_stat_name', full_name='istio.mesh.v1alpha1.MeshConfig.inbound_cluster_stat_name', index=29,
      number=44, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outbound_cluster_stat_name', full_name='istio.mesh.v1alpha1.MeshConfig.outbound_cluster_stat_name', index=30,
      number=45, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='certificates', full_name='istio.mesh.v1alpha1.MeshConfig.certificates', index=31,
      number=47, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='thrift_config', full_name='istio.mesh.v1alpha1.MeshConfig.thrift_config', index=32,
      number=49, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='service_settings', full_name='istio.mesh.v1alpha1.MeshConfig.service_settings', index=33,
      number=50, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='enable_prometheus_merge', full_name='istio.mesh.v1alpha1.MeshConfig.enable_prometheus_merge', index=34,
      number=51, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_MESHCONFIG_OUTBOUNDTRAFFICPOLICY, _MESHCONFIG_THRIFTCONFIG, _MESHCONFIG_SERVICESETTINGS, ],
  enum_types=[
    _MESHCONFIG_INGRESSCONTROLLERMODE,
    _MESHCONFIG_AUTHPOLICY,
    _MESHCONFIG_ACCESSLOGENCODING,
    _MESHCONFIG_H2UPGRADEPOLICY,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=187,
  serialized_end=3061,
)


_CONFIGSOURCE = _descriptor.Descriptor(
  name='ConfigSource',
  full_name='istio.mesh.v1alpha1.ConfigSource',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='address', full_name='istio.mesh.v1alpha1.ConfigSource.address', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tls_settings', full_name='istio.mesh.v1alpha1.ConfigSource.tls_settings', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subscribed_resources', full_name='istio.mesh.v1alpha1.ConfigSource.subscribed_resources', index=2,
      number=3, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=3064,
  serialized_end=3224,
)


_CERTIFICATE = _descriptor.Descriptor(
  name='Certificate',
  full_name='istio.mesh.v1alpha1.Certificate',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='secret_name', full_name='istio.mesh.v1alpha1.Certificate.secret_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dns_names', full_name='istio.mesh.v1alpha1.Certificate.dns_names', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=3226,
  serialized_end=3279,
)

_MESHCONFIG_OUTBOUNDTRAFFICPOLICY.fields_by_name['mode'].enum_type = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE
_MESHCONFIG_OUTBOUNDTRAFFICPOLICY.containing_type = _MESHCONFIG
_MESHCONFIG_OUTBOUNDTRAFFICPOLICY_MODE.containing_type = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY
_MESHCONFIG_THRIFTCONFIG.fields_by_name['rate_limit_timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_MESHCONFIG_THRIFTCONFIG.containing_type = _MESHCONFIG
_MESHCONFIG_SERVICESETTINGS_SETTINGS.containing_type = _MESHCONFIG_SERVICESETTINGS
_MESHCONFIG_SERVICESETTINGS.fields_by_name['settings'].message_type = _MESHCONFIG_SERVICESETTINGS_SETTINGS
_MESHCONFIG_SERVICESETTINGS.containing_type = _MESHCONFIG
_MESHCONFIG.fields_by_name['connect_timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_MESHCONFIG.fields_by_name['protocol_detection_timeout'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_MESHCONFIG.fields_by_name['tcp_keepalive'].message_type = networking_dot_v1alpha3_dot_destination__rule__pb2._CONNECTIONPOOLSETTINGS_TCPSETTINGS_TCPKEEPALIVE
_MESHCONFIG.fields_by_name['ingress_controller_mode'].enum_type = _MESHCONFIG_INGRESSCONTROLLERMODE
_MESHCONFIG.fields_by_name['auth_policy'].enum_type = _MESHCONFIG_AUTHPOLICY
_MESHCONFIG.fields_by_name['rds_refresh_delay'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_MESHCONFIG.fields_by_name['access_log_encoding'].enum_type = _MESHCONFIG_ACCESSLOGENCODING
_MESHCONFIG.fields_by_name['default_config'].message_type = mesh_dot_v1alpha1_dot_proxy__pb2._PROXYCONFIG
_MESHCONFIG.fields_by_name['outbound_traffic_policy'].message_type = _MESHCONFIG_OUTBOUNDTRAFFICPOLICY
_MESHCONFIG.fields_by_name['config_sources'].message_type = _CONFIGSOURCE
_MESHCONFIG.fields_by_name['enable_auto_mtls'].message_type = google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE
_MESHCONFIG.fields_by_name['locality_lb_setting'].message_type = networking_dot_v1alpha3_dot_destination__rule__pb2._LOCALITYLOADBALANCERSETTING
_MESHCONFIG.fields_by_name['dns_refresh_rate'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_MESHCONFIG.fields_by_name['h2_upgrade_policy'].enum_type = _MESHCONFIG_H2UPGRADEPOLICY
_MESHCONFIG.fields_by_name['certificates'].message_type = _CERTIFICATE
_MESHCONFIG.fields_by_name['thrift_config'].message_type = _MESHCONFIG_THRIFTCONFIG
_MESHCONFIG.fields_by_name['service_settings'].message_type = _MESHCONFIG_SERVICESETTINGS
_MESHCONFIG.fields_by_name['enable_prometheus_merge'].message_type = google_dot_protobuf_dot_wrappers__pb2._BOOLVALUE
_MESHCONFIG_INGRESSCONTROLLERMODE.containing_type = _MESHCONFIG
_MESHCONFIG_AUTHPOLICY.containing_type = _MESHCONFIG
_MESHCONFIG_ACCESSLOGENCODING.containing_type = _MESHCONFIG
_MESHCONFIG_H2UPGRADEPOLICY.containing_type = _MESHCONFIG
_CONFIGSOURCE.fields_by_name['tls_settings'].message_type = networking_dot_v1alpha3_dot_destination__rule__pb2._CLIENTTLSSETTINGS
_CONFIGSOURCE.fields_by_name['subscribed_resources'].enum_type = _RESOURCE
DESCRIPTOR.message_types_by_name['MeshConfig'] = _MESHCONFIG
DESCRIPTOR.message_types_by_name['ConfigSource'] = _CONFIGSOURCE
DESCRIPTOR.message_types_by_name['Certificate'] = _CERTIFICATE
DESCRIPTOR.enum_types_by_name['Resource'] = _RESOURCE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

MeshConfig = _reflection.GeneratedProtocolMessageType('MeshConfig', (_message.Message,), {

  'OutboundTrafficPolicy' : _reflection.GeneratedProtocolMessageType('OutboundTrafficPolicy', (_message.Message,), {
    'DESCRIPTOR' : _MESHCONFIG_OUTBOUNDTRAFFICPOLICY,
    '__module__' : 'mesh.v1alpha1.config_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig.OutboundTrafficPolicy)
    })
  ,

  'ThriftConfig' : _reflection.GeneratedProtocolMessageType('ThriftConfig', (_message.Message,), {
    'DESCRIPTOR' : _MESHCONFIG_THRIFTCONFIG,
    '__module__' : 'mesh.v1alpha1.config_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig.ThriftConfig)
    })
  ,

  'ServiceSettings' : _reflection.GeneratedProtocolMessageType('ServiceSettings', (_message.Message,), {

    'Settings' : _reflection.GeneratedProtocolMessageType('Settings', (_message.Message,), {
      'DESCRIPTOR' : _MESHCONFIG_SERVICESETTINGS_SETTINGS,
      '__module__' : 'mesh.v1alpha1.config_pb2'
      # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig.ServiceSettings.Settings)
      })
    ,
    'DESCRIPTOR' : _MESHCONFIG_SERVICESETTINGS,
    '__module__' : 'mesh.v1alpha1.config_pb2'
    # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig.ServiceSettings)
    })
  ,
  'DESCRIPTOR' : _MESHCONFIG,
  '__module__' : 'mesh.v1alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.MeshConfig)
  })
_sym_db.RegisterMessage(MeshConfig)
_sym_db.RegisterMessage(MeshConfig.OutboundTrafficPolicy)
_sym_db.RegisterMessage(MeshConfig.ThriftConfig)
_sym_db.RegisterMessage(MeshConfig.ServiceSettings)
_sym_db.RegisterMessage(MeshConfig.ServiceSettings.Settings)

ConfigSource = _reflection.GeneratedProtocolMessageType('ConfigSource', (_message.Message,), {
  'DESCRIPTOR' : _CONFIGSOURCE,
  '__module__' : 'mesh.v1alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.ConfigSource)
  })
_sym_db.RegisterMessage(ConfigSource)

Certificate = _reflection.GeneratedProtocolMessageType('Certificate', (_message.Message,), {
  'DESCRIPTOR' : _CERTIFICATE,
  '__module__' : 'mesh.v1alpha1.config_pb2'
  # @@protoc_insertion_point(class_scope:istio.mesh.v1alpha1.Certificate)
  })
_sym_db.RegisterMessage(Certificate)


DESCRIPTOR._options = None
_MESHCONFIG.fields_by_name['auth_policy']._options = None
_MESHCONFIG.fields_by_name['rds_refresh_delay']._options = None
# @@protoc_insertion_point(module_scope)
