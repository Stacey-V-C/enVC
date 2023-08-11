import { createSignal } from "solid-js";
import type { Result } from "../types";


export const [selectedData, setSelectedData] = createSignal<Result | null>([1, 2]);
