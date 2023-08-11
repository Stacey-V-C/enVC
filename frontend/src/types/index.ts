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

export type Result = any[];
