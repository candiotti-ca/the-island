import { writable } from "svelte/store";
import { MapType } from "../models/MapType";

export const mapType = writable(MapType.POINTY_EVEN);