export type SQLResult = {
  statement: string,
  headers: string[],
  values: string[][],
}

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
  sqlResults: AsyncEvent<SQLResult>[],
  chromaResults: AsyncEvent<string>[],
}

export type ResultType = 'sqlFormattedResults' | 'sqlRawResults' | 'chromaResults'

// export type Result = {
//   dataType: ResultType,
//   data: any[],
// }

export type Result =
//   | {
//     dataType: 'sqlFormattedResults',
//     data: SQLResult,
//   }
//   | {
//     dataType: 'sqlRawResults',
//     data: SQLResult,
//   }

// export type GeneralResult =
  | {
    dataType: ResultType,
    input: string,
    values: string[]
  }
  | {
    dataType: ResultType,
    input: string,
    headers: string[],
    values: string[][],
  }

