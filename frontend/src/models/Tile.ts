import type { TileType } from "./TileType";

export interface Tile {
    q: number;
    r: number;
    type: TileType;
    explorers: any[];
    boat?: any;
    shark?: any;
    whale?: any;
    seaSerpent?: any;
}