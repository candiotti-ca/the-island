import { writable, type Writable } from "svelte/store";
import { MapType } from "../models/MapType";
import type { SelectedEntity } from "../models/SelectedEntity";

export const mapType = writable(MapType.POINTY_EVEN);
export const selectedEntity: Writable<SelectedEntity | undefined> = writable(undefined)