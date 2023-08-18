import { Setter } from "solid-js";
import { createStore } from "solid-js/store";
import { Result, ResultType } from "../types";
import * as UIReceiver from "../../wailsjs/go/services/UIReceiver";

export const [uiState, setUiState] = createStore({
  terminal: {
    open: false,
  },
  dataPanel: {
    open: false,
  },
  activeQuery: null,
});

type InitKeybindingsArgs = {
  setSelectedData: Setter<Result | null>;
}

export const initKeybindings = (context: InitKeybindingsArgs) => {
  document.addEventListener("keydown", (event) => {
    if (
      event.ctrlKey &&
      !event.shiftKey &&
      !event.altKey &&
      !event.metaKey &&
      keyBindings[event.key]
    ) {
      event.preventDefault();
      keyBindings[event.key](context);
    }
  });
};

const toggleTerminal = () => {
  setUiState("terminal", { open: !uiState.terminal.open });

  if (uiState.terminal.open) {
    document.getElementById("terminal-input")?.focus();
  }
}

const toggleDataPanel = () => {
  setUiState("dataPanel", { open: !uiState.dataPanel.open });

  if (uiState.dataPanel.open) {
    document.getElementById("data-panel-input")?.focus();
  }
}

const dataResultsActions: ResultType[] = [
  "sqlFormattedResults",
]

// think none of these have to be async because the results get posted elsewhere?
// or they might need a new ID returned if we want to be able to track them
// separately from the original query
const loadActiveQueryData = ({
  setSelectedData,
}: InitKeybindingsArgs) => {
  if (uiState.activeQuery) {
    UIReceiver.PullData(uiState.activeQuery).then(result => {
      if (!result) return;

      const [payloadString, action] = result;

      const payload = extractPayload(payloadString);

      if (dataResultsActions.includes(<ResultType>action)) {
        setSelectedData({
          dataType: <ResultType>action,
          ...payload,
        });
      }
    });
  }
}

const extractPayload = (raw: string) => {
  const obj = JSON.parse(raw);

  for (const key in obj) {
    if (!["values", "headers", "input"].includes(key)) {
      delete obj[key];
    }
  }

  return obj;
}

const rerunActiveQuery = () => {

}

const copyActiveQueryToTerminal = () => {

}

// Do these live here or inside the terminal?  
// What's nice about having them here is that they can be used
// for multiple similar cases, i.e. if the terminal and a query in it are selected
// it pulls from there, but if i.e. an app component is selected
// we can target the query that caused that component to render

const keyBindings: Record<string, (context: InitKeybindingsArgs) => void> = {
  "r": toggleDataPanel,
  "t": toggleTerminal,
  "q": loadActiveQueryData,
  "w": rerunActiveQuery,
  "a": copyActiveQueryToTerminal,
}