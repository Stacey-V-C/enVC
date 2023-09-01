import { Setter } from "solid-js";
import { createStore } from "solid-js/store";
import { Res, Result, ResultType } from "../types";
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
  Res.SQL_FORMATTED_RESULTS,
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

      const { action, payload } = result;

      if (dataResultsActions.includes(<ResultType>action)) {
        setSelectedData({
          dataType: <ResultType>action,
          ...payload,
        });
      }
    });
  }
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