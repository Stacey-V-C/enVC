from pb.nvc_pb2_grpc import \
  SQLToChromaListenerServicer, \
    add_SQLToChromaListenerServicer_to_server

from pb.nvc_pb2 import ChromaResult, SQLAction

from utils.parsing import get_value_parser, parse_sql_results
from utils.fs import get_config

import grpc

from concurrent import futures

config = get_config()

data_models = config["dataModels"] 

class SQLChromaServicer(SQLToChromaListenerServicer):
    def LogSQLAction(self, request, context):
        CreateNewSQLEntities(request)
        
        return ChromaResult(result="OK") # placeholder
    
def start_server():
    port = config["py_chroma_port"] or 50051
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_SQLToChromaListenerServicer_to_server(SQLChromaServicer(), server)
    server.add_insecure_port(f'[::]:{port}')
    server.start()
    print("Server started on port {port}")
    server.wait_for_termination()


    
def CreateNewSQLEntities(action: SQLAction):
    data_model = data_models[action.dataModel]

    table_name = action.table

    schema = data_model[table_name]

    if schema is None:
        return None

    columns = action.columns

    values: list[list[str]] = action.values

    parser = get_value_parser(schema, columns)

    parsed_results = parse_sql_results(values, parser)

    AddToChroma(parsed_results)
        
    
def AddToChroma(results):
    pass

