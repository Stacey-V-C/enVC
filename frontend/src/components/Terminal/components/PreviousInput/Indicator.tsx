import type { Indicator } from ".";

import { Component, For } from "solid-js";

const getColorByIndicator = (indicator: Indicator) => {
  switch (indicator) {
    case "error":
      return "red";
    case "sql":
      return "blue";
    case "chroma":
      return "green";
  }
}

export const Indicators: Component<{ indicators?: Indicator[] }> = (props) =>
  props?.indicators && props?.indicators?.length > 0 ? (
    <For each={props.indicators}>
      {indicator =>
        <div
          class="inline-block w-2 h-2 rounded-full"
          style={{ "background-color": getColorByIndicator(indicator) }}
        />
      }
    </For>
  ) : null;