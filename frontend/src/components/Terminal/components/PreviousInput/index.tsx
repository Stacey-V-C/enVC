import { Component, } from "solid-js";

import { Colors } from "../../../../theme";

export type Indicator = "error" | "sql" | "chroma"

export type InputData = {
  odd: boolean; // we don't want a distinct index here because of garbage collection - just make sure they alternate
  text: string;
  indicators?: Indicator[];
}

type InputProps = InputData & { selected?: boolean; }

export const PreviousInput: Component<InputProps> = (props) => (
  <div
    class="p-2 m-1 text-xs"
    style={{
      "background-color": props.odd ? Colors.PaleGreen : Colors.PaleGrey,
      border: props.selected ? `1px solid blue` : "none",
    }}
  >
    {props.text}
  </div>
)
