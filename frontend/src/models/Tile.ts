import type { Boat } from "./Boat";
import type { Explorer } from "./Explorer";
import type { SeaSerpent } from "./SeaSerpent";
import type { Shark } from "./Shark";
import type { TileType } from "./TileType";
import type { Whale } from "./Whale";

export interface Tile {
    id: string;
    q: number;
    r: number;
    type: TileType;
    explorers: Explorer[];
    boat: Boat | undefined;
    shark: Shark | undefined;
    whale: Whale | undefined;
    seaSerpent: SeaSerpent | undefined;
}