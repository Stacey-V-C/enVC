import {
  For,
  Show,
  createSignal,
} from "solid-js";
import type { Component } from "solid-js";

import { Tooltip } from "../atoms/Tooltip";

/* 

This is really the shape of the thing but I'm not sure there's a useful way to import it into typescript...
We would need to be defining a component type for each new relation from the db, and because typescript ultimately
compiles to javascript, we're not actually getting any benefit for that new compilation (since we're also not
actually writing the code ourselves if auto generated so we're not getting any type safety from it either)

Keep this in back of mind - if somebody wanted to write CUSTOM COMPONENTS for each document type in a data model,
having these type definitions would be a really nice starting point...

type DefinitionSet<T, D extends string> = {
  document: T;
  definitionSet: Record<D, keyof T[]>;
}

type TextExpressionGroups = "main" | "additional" | "metadata";

type TextExpressionProps<T> = DefinitionSet<T, TextExpressionGroups>;

*/

type DefinitionSet<D extends string> = {
  document: any;
  definitionSet: Record<D, string[]>;
}

type TextExpressionProps = DefinitionSet<"main" | "additional" | "metadata">;


export const TextExpression: Component<TextExpressionProps> = (props) => {
  const [expanded, setExpanded] = createSignal(false);

  const MetadataTooltip = () => (
    <Tooltip
      indicator={
        <div
          class="border-radius-50px white text-16 text-center w-24 h-24 flex justify-center items-center"
          style={{ color: "var (--color-pale-green)" }}
        >
          M
        </div >
      }
    >
      <For each={props.definitionSet.metadata}>
        {metadata => (
          <p class="text-12">
            <p class="text-12 font-bold">{metadata}</p>
            {props.document[metadata]}
          </p>
        )}
      </For>
    </Tooltip >
  )

  return (
    <div class="flex flex-col">
      <For each={props.definitionSet.main}>
        {main => (
          <p class="text-16">{props.document[main]}</p>
        )}
      </For>
      <p
        class="text-24"
        onClick={() => setExpanded(e => !e)}
      >
        {expanded() ? "-" : "+"}
      </p>
      <Show when={expanded()}>
        <div class="flex flex-col">
          <For each={props.definitionSet.additional}>
            {additional => (
              <p class="text-12">
                <p class="text-12 font-bold">{additional}</p>
                {props.document[additional]}
              </p>
            )}
          </For>
          <MetadataTooltip />
        </div>
      </Show>
    </div>
  )
}
