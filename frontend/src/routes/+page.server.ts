interface BackendTile {
    X: number;
    Y: number;
}

export function load(): any {
    return {
        tiles: fetch('http://localhost:8080/map')
            .then((response) => response.json())
            .then((tiles: BackendTile[]) => tiles.map(tile => ({ x: tile.X, y: tile.Y })))
            .catch((error) => {
                console.error(error);
                return "fail";
            })
    };
}