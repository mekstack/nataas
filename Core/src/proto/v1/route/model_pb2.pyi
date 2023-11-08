from proto.v1.subdomain import model_pb2 as _model_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, \
    Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor


class Route(_message.Message):
    __slots__ = ["target_ip", "subdomains"]
    TARGET_IP_FIELD_NUMBER: _ClassVar[int]
    SUBDOMAINS_FIELD_NUMBER: _ClassVar[int]
    target_ip: str
    subdomains: _containers.RepeatedCompositeFieldContainer[_model_pb2.Subdomain]
    
    def __init__(self, target_ip: _Optional[str] = ...,
                 subdomains: _Optional[_Iterable[_Union[_model_pb2.Subdomain, _Mapping]]] = ...) -> None: ...
