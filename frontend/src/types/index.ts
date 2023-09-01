export type EventType = 'error' | 'sql' | 'chroma'

export type AsyncEvent<T> = {
  notified: boolean,
  read: boolean,
  received: Date,
  data: T,
  type: 'error' | 'sql' | 'chroma' // TODO would be nice if these were dependent?
}

export type Events = {
  errors: AsyncEvent<string>[],
  sqlResults: AsyncEvent<LabeledResult>[],
  chromaResults: AsyncEvent<string>[],
}

export enum Res {
  SQL_FORMATTED_RESULTS = 'sqlFormattedResults',
  SQL_RAW_RESULTS = 'sqlRawResults',
  CHROMA_RESULTS = 'chromaResults',
}

export type LabeledResultType =
  | Res.SQL_FORMATTED_RESULTS
  | Res.SQL_RAW_RESULTS
  | Res.CHROMA_RESULTS

export type UnlabeledResultType = 'TODO'

export type ResultType = LabeledResultType | UnlabeledResultType

export type UnlabeledResult = {
  dataType: UnlabeledResultType,
  input?: string,
  data: string[]
}

export type LabeledResult = {
  dataType: LabeledResultType,
  input?: string,
  columns: string[],
  data: string[][],
}

export type Result = UnlabeledResult | LabeledResult