import { Show, createEffect, createSignal } from "solid-js";
import { Terminal } from "./components/Terminal";
import { DataPanel } from "./components/DataPanel";
import { EventBar } from "./components/EventBar";
import { createEventStore } from "./state/events";
import { uiState, initKeybindings } from "./state/ui";
import { setSelectedData } from "./state/data";
import { PullData } from "../wailsjs/go/services/UIReceiver";

initKeybindings({
  setSelectedData,
});

const [events, setEvents] = createEventStore();

export const [id, setId] = createSignal("");

createEffect(() => {
  if (id()) PullData(id()).then(setSelectedData);
});

function App() {
  return (
    <div class="w-screen h-screen overflow-hidden flex flex-col align-stretch justify-stretch">
      <EventBar events={events} />
      {/* <button class="w-10 h-10 bg-red-500" onClick={handleClick}>Click me</button> */}
      <div class="h-full w-full relative">
        <div class="h-full absolute left-0">
          <Show when={uiState.dataPanel.open}>
            <DataPanel />
          </Show>
        </div>
        <div class="h-full absolute right-0">
          <Show when={uiState.terminal.open}>
            <Terminal />
          </Show>
        </div>
      </div>
    </div>
  );
}

export default App;
