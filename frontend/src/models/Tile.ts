export interface BackendTile {
    X: number;
    Y: number;
}

export class Tile {
    x: number;
    y: number;

    constructor(backend: BackendTile) {
        this.x = backend.X;
        this.y = backend.Y;
    }
}