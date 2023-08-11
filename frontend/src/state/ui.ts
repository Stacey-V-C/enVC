import { createStore } from "solid-js/store";

export const [uiState, setUiState] = createStore({
  terminal: {
    open: false,
  },
  dataPanel: {
    open: false,
  },
});

export const initKeybindings = () => {
  document.addEventListener("keydown", (event) => {
    if (
      event.ctrlKey &&
      !event.shiftKey &&
      !event.altKey &&
      !event.metaKey &&
      keyBindings[event.key]
    ) {
      event.preventDefault();
      keyBindings[event.key]();
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

const keyBindings: Record<string, () => void> = {
  "r": toggleDataPanel,
  "t": toggleTerminal,
}