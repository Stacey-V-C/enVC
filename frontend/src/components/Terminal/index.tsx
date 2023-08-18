import { createEffect, createSignal, For } from "solid-js";

import { PreviousInput } from "./components/PreviousInput";
import { Indicators } from "./components/PreviousInput/Indicator";

import type { InputData } from "./components/PreviousInput";

export const [termInput, setTermInput] = createSignal("");

export const [previousInputs, setPreviousInputs] = createSignal<InputData[]>([]);
export const [previousInputPointer, setPreviousInputPointer] = createSignal<number | null>(null);

import { handleInput, handleKeyPress } from "./helpers";
import { InputType } from "./components/InputType";

export const Terminal = () => {
  let inputElement: HTMLTextAreaElement;

  createEffect(() => {
    const { scrollHeight, style } = inputElement;

    if (!termInput()) {
      style.height = "auto"
      style.height = "64px"
    } else {
      style.height = "auto";
      style.height = `${scrollHeight}px`;
    }
  })

  return (
    <div class="flex flex-col h-full w-[300px] justify-end bg-off-white border-pale-grey border border-solid">
      <For each={previousInputs()}>
        {(input, i) => <div>
          <div class="text-sm text-gray-400">
            <InputType callSign="SQL" size="sm" />
            <Indicators indicators={input.indicators} />
            <PreviousInput
              selected={i() === previousInputPointer()}
              {...input}
            />
          </div>
        </div>}
      </For>
      <div class="relative min-h-64">
        <InputType callSign="SQL" />
        <textarea
          id="terminal-input"
          class="h-auto w-full min-w-144 p-2 bg-pale-blue text-sm bg-off-white border-0 outline-none"
          // @ts-expect-error
          ref={inputElement}
          value={termInput()}
          onInput={handleInput}
          onKeyDown={handleKeyPress}
        />
      </div>
    </div>
  )
}
