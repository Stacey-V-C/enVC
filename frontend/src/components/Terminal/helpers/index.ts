import {
  previousInputPointer,
  setPreviousInputPointer,
  previousInputs,
  setPreviousInputs,
  termInput,
  setTermInput
} from ".."

import { MODES, activeMode } from "../../../state/modes"

const storeAndResetInput = () => {
  const mode = MODES[activeMode()];

  const {
    statement = termInput(),
    leftOvers = ""
  } = mode?.validate(termInput());

  if (statement) mode.send(statement);

  setPreviousInputs((prev) =>
    [
      ...prev,
      {
        text: statement,
        odd: prev?.length
          ? !prev[prev.length - 1].odd
          : true
      }
    ]
  );

  setTermInput(leftOvers);
}

export const handleInput = (event: Event) => {
  const { value } = event.target as HTMLInputElement;

  setTermInput(value);
}

export const handleKeyPress = (event: KeyboardEvent) => {
  if (event.key === "Enter") {
    event.preventDefault();

    storeAndResetInput();
  } else if (event.key === "ArrowDown") {
    event.preventDefault();

    if (previousInputs()?.length) {
      const current = previousInputPointer();

      const newPointer =
        current === null || current === previousInputs().length - 1
          ? 0
          : current + 1;
      setPreviousInputPointer(newPointer);
    }
  } else if (event.key === "ArrowUp") {
    event.preventDefault();

    if (previousInputs()?.length) {
      const current = previousInputPointer();

      const newPointer = !current
        ? previousInputs().length - 1
        : current - 1;

      setPreviousInputPointer(newPointer);
    }
  } else if (event.key === "Escape") {
    setPreviousInputPointer(null);
  }
}