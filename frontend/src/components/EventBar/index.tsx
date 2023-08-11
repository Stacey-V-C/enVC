import { Component, For, createEffect, createSignal } from "solid-js";
import type { Events, EventType } from "../../types";
// import type { Accessor } from "solid-js";

type EventNotification = {
  eventType: EventType;
  id: string;
};

export const EventBar: Component<{
  events: Events;
//   setEvents: (events: Events) => void;
}> = (props) => {
//   const [newEvents, setNewEvents] = createSignal<EventNotification[]>([]);

//   createEffect(() => {
//     let newEvents = Object.values(props.events())
//       .flat()
//       .filter((event) => !event.notified);

//     // TODO the events need to be made not new in the event store after they are read
//     if (newEvents.length > 0) {
//       const newNotifications = newEvents.map((event) => {
//         return {
//           eventType: event.type,
//           id: Math.random().toString(36).substring(7),
//         };
//       });
//       setNewEvents((existing) => [...existing, ...newNotifications]);

//       setTimeout(() => {
//         const idsToRemove = newNotifications.map((event) => event.id);
//         setNewEvents((existing) =>
//           existing.filter((event) => !idsToRemove.includes(event.id))
//         );
//       }, 5000);
//     }
//   });

  return (
    <div class="w-screen h-5 bg-[black] border-off-white border border-solid">
      {/* <For each={newEvents()}>
        {(event) => <NotificationIndicator event={event} />}
      </For> */}
    </div>
  );
};

const colorByEventType = (eventType: EventType) => {
  switch (eventType) {
    case "error":
      return "red";
    case "sql":
      return "green";
    case "chroma":
      return "blue";
  }
};
/*
Note: Care should be taken when using a transition immediately after:

adding the element to the DOM using .appendChild()
removing an element's display: none; property.
This is treated as if the initial state had never occurred and the element 
was always in its final state. The easy way to overcome this limitation is 
to apply a setTimeout() of a handful of milliseconds before changing the CSS 
property you intend to transition to. 
*/

const fadeOpacityInOut = (element: HTMLElement) => {
  element.style.opacity = "1";
  setTimeout(() => {
    element.style.opacity = "0";
  }, 1000);
};

const NotificationIndicator: Component<{
  event: EventNotification;
}> = (props) => (
  <div
    ref={(indicator) => fadeOpacityInOut(indicator!)}
    class="w-5 h-5 rounded-full transition-fade opacity-0"
    style={{
      "background-color": colorByEventType(props.event.eventType),
    }}
  />
);
