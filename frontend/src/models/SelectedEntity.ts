import type { Boat } from "./Boat";
import type { Explorer } from "./Explorer";
import type { SeaSerpent } from "./SeaSerpent";
import type { Shark } from "./Shark";
import type { Tile } from "./Tile";
import type { Whale } from "./Whale";

export interface SelectedEntity {
    type: EntityType,
    entity: Boat | Whale | Shark | SeaSerpent | Explorer,
    tile: Tile
}

export type EntityType = "boat" | "shark" | "sea_serpent" | "whale" | "explorer";