from pb.nvc_pb2 import SQLAction, SQLtype

def parser_by_type(type: str):
    parser_dict: dict[str, function] = {
        "number": int,
        "string": str,
        "date": str
    }
    
    return parser_dict[type]


def get_value_parser(
        schema: dict[str, dict[str, dict[str, str]]],
        columns: list[str]):
    
    parser_dict: dict[str, function] = {}
    
    for column in columns:
        col_def = schema["fields"][column]

        if col_def is not None:
            parser_dict[column] = parser_by_type(col_def)

def parse_sql_results(
        results: list[list[str]], 
        parser: dict[str, function]):
    parsed_results = []

    for result in results:
        parsed_result = {}

        for column in result:
            parsed_result[column] = parser[column](result[column])

        parsed_results.append(parsed_result)

    return parsed_results

def parse_sql_action(data_models, action: SQLAction):    
    data_model = data_models[action.dataModel]

    table_name = action.table

    schema = data_model[table_name]

    if schema is None:
        return None

    columns = action.columns

    values: list[list[str]] = action.values

    parser = get_value_parser(schema, columns)

    parsed_results = parse_sql_results(values, parser)

    return parsed_results

def parse_ids(action: SQLAction):
    id_column = action.columns.index("id")

    extract_id = lambda row: row[id_column]

    ids = [extract_id(row.values) for row in action.values]

    return ids