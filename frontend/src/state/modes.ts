import { createSignal } from "solid-js"
import { SendSQL } from "../../wailsjs/go/controllers/Controller"

type ModeObject = {
  style: any,
  validate: (q: string) => {
    statement: string,
    leftOvers: string,
  },
  send: (s: string) => Promise<any>
}

export const MODES: Record<string, ModeObject> = {
  // when uploading these from plugins we should be able to run a spread operator on a default function
  // to fill in things like validation function I think?

  "SQL": {
    style: {
      "background-color": "hsl(45, 75%, 85%)",
      "color": "hsl(110, 60%, 45%)",
      border: "1.5px ridge hsl(110, 60%, 45%)"
    },
    validate: (query: string) => {
      const completeQuery = (query?.includes(";"))

      const [statement, ...leftOvers] = completeQuery
        ? query?.split(";")
        : ["", [query]]

      return {
        statement: statement ? `${statement};` : "",
        leftOvers: leftOvers.join(";")
      }
    },
    send: SendSQL
  },
}

export type Mode = keyof typeof MODES

export const [activeMode, setActiveMode] = createSignal<Mode>("SQL")
