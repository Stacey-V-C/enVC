import type { Component } from "solid-js"
import { MODES } from "../../../state/modes"

const SIZES = {
  "sm": {
    "font-size": "8px",
    "flex-direction": "row",
  },
  "md": {
    "font-size": "14px",
    "flex-direction": "column",
    "padding": "0px 8px"
  }
}

export const InputType: Component<{ callSign: string, size?: "sm" | "md" }> = (props) => {
  return (
    <div
      class="flex flex-col justify-center h-18 w-8 absolute -left-8"
      style={{
        ...MODES[props.callSign]?.style,
        ...SIZES[props?.size || "md"]
      }}
    >
      <p>{props.callSign[0] || ""}</p>
      <p>{props.callSign[1] || ""}</p>
      <p>{props.callSign[2] || ""}</p>
    </div>
  )
}