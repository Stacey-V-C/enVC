# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import nvc_pb2 as nvc__pb2


class SQLToChromaListenerStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.LogSQLAction = channel.unary_unary(
                '/nvc.SQLToChromaListener/LogSQLAction',
                request_serializer=nvc__pb2.SQLAction.SerializeToString,
                response_deserializer=nvc__pb2.ChromaResult.FromString,
                )


class SQLToChromaListenerServicer(object):
    """Missing associated documentation comment in .proto file."""

    def LogSQLAction(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SQLToChromaListenerServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'LogSQLAction': grpc.unary_unary_rpc_method_handler(
                    servicer.LogSQLAction,
                    request_deserializer=nvc__pb2.SQLAction.FromString,
                    response_serializer=nvc__pb2.ChromaResult.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'nvc.SQLToChromaListener', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class SQLToChromaListener(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def LogSQLAction(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/nvc.SQLToChromaListener/LogSQLAction',
            nvc__pb2.SQLAction.SerializeToString,
            nvc__pb2.ChromaResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
