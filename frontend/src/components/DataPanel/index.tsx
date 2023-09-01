import {
  Component,
  For,
  Match,
  Switch,
  //   createEffect,
  //   createSignal
} from "solid-js";
import { selectedData } from "../../state/data";
import { SQLDocVert } from "./SQLDocVert";
import type { LabeledResult } from "../../types";

export const DataPanel: Component = () => {
  //   const [page, setPage] = createSignal<number | null>(0);

  //   const rowMaxLength = 10;

  //   createEffect(() => {
  //     if (selectedData()?.data?.length || 0 > rowMaxLength) {
  //       setPage(0);
  //     } else {
  //       setPage(null);
  //     }
  //   });

  const selected = selectedData()?.hasOwnProperty("columns")
    ? selectedData() as LabeledResult
    : null;

  return (
    <div class="flex flex-col h-full w-[300px] bg-pale-green border-off-white border border-solid">
      <Switch fallback={<NoDataMessage />}>
        <Match when={selected?.dataType === "sqlFormattedResults"}>
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
        </Match>
      </Switch>
    </div >
  );
};

const NoDataMessage = () => {
  return (
    <p class="text-center text-16 m-4">
      No data selected - Select a query or app with data and press ctrl-q or select "load data" from menu
    </p>
  )
}