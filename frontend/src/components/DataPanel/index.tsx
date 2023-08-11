import {
  Component,
  For,
  Show,
  createEffect,
  createSignal
} from "solid-js";
import { selectedData } from "../../state/data";

export const DataPanel: Component = () => {
  const [page, setPage] = createSignal<number | null>(0);

  const rowMaxLength = 10;

  createEffect(() => {
    if (selectedData()?.length || 0 > rowMaxLength) {
      setPage(0);
    } else {
      setPage(null);
    }
  });

  return (
    <div class="flex flex-col h-full w-[300px] justify-start bg-pale-green border-off-white border border-solid">
      <Show when={selectedData()} fallback={<NoDataMessage />}>
        <For each={selectedData()}>
          {row => <SQLResultRow row={row} />}
        </For>
      </Show>
    </div>
  );
};

export const NoDataMessage = () => {
  return (
    <p class="text-center text-16">
      No data selected - Select a query or app with data and press ctrl-o or select "load data" from menu
    </p>
  )
}


export const SQLResultRow = (props: { row: any }) => {
  const content = typeof props.row === "string" ? props.row : JSON.stringify(props.row);
  return (
    <div>
      {content}
    </div>
  )
}