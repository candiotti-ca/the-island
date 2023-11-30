import type { TileType } from "./TileType";

export interface Tile {
    x: number;
    y: number;
    type: TileType;
}