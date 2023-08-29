import {
  Component,
  For,
  Show,
//   createEffect,
//   createSignal
} from "solid-js";
import { selectedData } from "../../state/data";
import { SQLDocVert } from "./SQLDocVert";
import type { LabeledResult } from "../../types";

export const DataPanel: Component = () => {
//   const [page, setPage] = createSignal<number | null>(0);

//   const rowMaxLength = 10;

  const selected = selectedData()?.hasOwnProperty("columns")
    ? selectedData() as LabeledResult
    : null;

//   createEffect(() => {
//     if (selectedData()?.data?.length || 0 > rowMaxLength) {
//       setPage(0);
//     } else {
//       setPage(null);
//     }
//   });

  return (
    <div class="flex flex-col h-full w-[300px] bg-pale-green border-off-white border border-solid">
      <Show when={!!selected} fallback={<NoDataMessage />}>
        <div class="flex flex-col mx-0 items-start h-auto w-auto justify-start bg-pale-green border-off-white border border-solid">
          <For each={selected?.data}>
            {document => (
              <SQLDocVert
                data={document}
                columns={selected?.columns || []}
              />
            )}
          </For>
        </div>
      </Show>
    </div>
  );
};

const NoDataMessage = () => {
  return (
    <p class="text-center text-16 m-4">
      No data selected - Select a query or app with data and press ctrl-q or select "load data" from menu
    </p>
  )
}