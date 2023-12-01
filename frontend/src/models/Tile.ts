import type { TileType } from "./TileType";

export interface Tile {
    q: number;
    r: number;
    type: TileType;
}