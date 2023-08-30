from pb.nvc_pb2_grpc import \
  SQLToChromaListenerServicer, \
    add_SQLToChromaListenerServicer_to_server

from pb.nvc_pb2 import ChromaResult, SQLAction, SQLtype

from utils.parsing import parse_sql_action
from utils.fs import get_config

import grpc

from concurrent import futures
import db.calls as db

config = get_config()

data_models = config["dataModels"] 

def get_client():
    pass

class SQLChromaServicer(SQLToChromaListenerServicer):
    def LogSQLAction(self, request, context):
        action_type = request.type

        if action_type is (None or "UNSPECIFIED"):
            return ChromaResult(result="ERROR")
        
        client = get_client()
        
        if action_type == "DELETE":
            parsed_ids = parse_ids(request)

            db.delete_from_chroma(client, parsed_ids)

            return ChromaResult(result="OK") # placeholder
        else:
          parsed_results = parse_sql_action(data_models, request)
          
          match action_type:
              case 'CREATE':
                  res = db.add_to_chroma(client, parsed_results)
              case 'UPDATE':
                  res = db.update_chroma(client, parsed_results)
              case 'MIGRATE':
                  res = db.add_to_chroma_if_not_exists(client, parsed_results)

          return ChromaResult(result="OK") # placeholder
    
def start_server():
    port = config["py_chroma_port"] or 50051
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_SQLToChromaListenerServicer_to_server(SQLChromaServicer(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    print("Server started on port {port}")
    server.wait_for_termination()
    