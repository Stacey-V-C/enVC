import { createSignal, For } from "solid-js";

import { PreviousInput } from "./components/PreviousInput";
import { Indicators } from "./components/PreviousInput/Indicator";

import type { InputData } from "./components/PreviousInput";

export const [termInput, setTermInput] = createSignal("");

export const [previousInputs, setPreviousInputs] = createSignal<InputData[]>([]);
export const [previousInputPointer, setPreviousInputPointer] = createSignal<number | null>(null);

import { handleInput, handleKeyPress } from "./helpers";


export const Terminal = () => {
  return (
    <div class="flex flex-col h-full w-[300px] justify-end bg-off-white border-pale-grey border border-solid">
      <For each={previousInputs()}>
        {(input, i) => <div>
          <div class="text-sm text-gray-400">
            <Indicators indicators={input.indicators} />
            <PreviousInput
              selected={i() === previousInputPointer()}
              {...input}
            />
          </div>
        </div>}
      </For>
      <input
        id="terminal-input"
        class="h-5 w-full min-w-144 p-2 bg-pale-blue text-sm bg-off-white border-0 outline-none"
        value={termInput()}
        onInput={handleInput}
        onKeyDown={handleKeyPress}
      />
    </div>
  )
}