import { createStore } from 'solid-js/store'
import { AsyncEvent, Events } from '../types'

type GetEventStoreFn = () => [Events, {
  markEventsNotified: () => void,
}]



export const createEventStore: GetEventStoreFn = () => {
  const [events, setEvents] = createStore<Events>({
    errors: [],
    sqlResults: [],
    chromaResults: [],
  })

  const actions = {
    markEventsNotified: () => {
      setEvents(
        'errors',
        event => event.notified === false,
        'notified',
        true
      )

      setEvents(
        'sqlResults',
        event => event.notified === false,
        'notified',
        true
      )

      setEvents(
        'chromaResults',
        event => event.notified === false,
        'notified',
        true
      )
    },
  }

  return [events, actions];
}