import type { Component, JSX } from "solid-js";

type TooltipProps = {
  indicator: JSX.Element;
  children: JSX.Element;
}

export const Tooltip: Component<TooltipProps> = (props) => (
  <div class="relative">
    <div class="tooltip-controller h-auto w-auto">
      {props.indicator}
    </div>
    <div
      class="tooltip absolute left-6 top-6 z-[500] border border-solid border-2 border-[#4a5568] p-2"
      style={{
        "background-color": "hsl(160, 30%, 75%)",
      }}>
      {props.children}
    </div>
  </div>
)             