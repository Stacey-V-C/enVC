import { Component, For } from "solid-js";
import { Tooltip } from "../atoms/Tooltip";

type DocRowProps = {
  data: string[]
  columns: string[]
}

export const SQLDocVert: Component<DocRowProps> = (props) => {
  return (
    <div class="flex mx-0 my-4">
      <div class="flex flex-col w-auto mx-4">
        <For each={props.columns}>
          {column => <p class="font-bold">{column}</p>}
        </For>
      </div>
      <div class="flex flex-col mr-4 w-120">
        <For each={props.data}>
          {/* TODO figure out the ellipsis!! */}
          {value =>
            <Tooltip
              indicator={<p class="w-48 truncate">{value}</p>}
            >
              <p class="w-120">{value}</p>
            </Tooltip>
          }
        </For>
      </div>
    </div>
  )
}