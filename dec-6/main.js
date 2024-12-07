(async () => {
  const fs = require("fs");
  const frameHeight = 20;
  const frameWidth = 40;
  const framesPerSecond = 60;
  let potentialObstructionCount = 0;

  // Print the map around the guard
  function printMap(map, guardLocation) {
    // clear the console
    console.clear();
    let startY = Math.max(0, guardLocation.y - frameHeight / 2);
    let startX = Math.max(0, guardLocation.x - frameWidth / 2);
    if (startY + frameHeight > map.length) {
      startY = map.length - frameHeight;
    }
    if (startX + frameWidth > map[startY].length) {
      startX = map[startY].length - frameWidth;
    }
    console.log(guardLocation);
    for (let y = startY; y < startY + frameHeight; y++) {
      console.log(map[y].slice(startX, startX + frameWidth).join(""));
    }
  }

  function countUniqueLocations(map) {
    return map.flat().reduce((acc, cell) => {
      if (cell === "." || cell === "#") {
        return acc;
      }
      return acc + 1;
    }, 0);
  }

  function findGuardLocation(map) {
    for (let y = 0; y < map.length; y++) {
      for (let x = 0; x < map[y].length; x++) {
        if (map[y][x] === "^") {
          return { x, y };
        }
      }
    }
  }

  function rotate_right(current) {
    let rotation = ["^", ">", "v", "<"];
    let index = rotation.indexOf(current);
    return rotation[(index + 1) % rotation.length];
  }

  function progress(map, location) {
    let move_next = (location, next) => {
      map[next.y][next.x] = map[location.y][location.x];
      let symbol;
      switch (map[location.y][location.x]) {
        case "^":
        case "v":
          symbol = "|";
          break;
        case "<":
        case ">":
          symbol = "-";
          break;
      }
      map[location.y][location.x] = symbol;
      location.y = next.y;
      location.x = next.x;
    };

    // 1. Calculate next location
    let next;
    switch (map[location.y][location.x]) {
      case "^":
        next = { x: location.x, y: location.y - 1 };
        break;
      case ">":
        next = { x: location.x + 1, y: location.y };
        break;
      case "v":
        next = { x: location.x, y: location.y + 1 };
        break;
      case "<":
        next = { x: location.x - 1, y: location.y };
        break;
    }

    // 1.5 Check if we've been at the next location before
    if (map[next.y][next.x] === "-" || map[next.y][next.x] === "|") {
      potentialObstructionCount++;
    }

    // 2. Determine what to do
    if (
      next.y >= map.length ||
      next.y < 0 ||
      next.x >= map[next.y].length ||
      next.x < 0
    ) {
      return false;
    }
    switch (map[next.y][next.x]) {
      case "#":
        map[location.y][location.x] = rotate_right(map[location.y][location.x]);
        break;
      default:
        move_next(location, next);
    }
    return true;
  }

  const data = fs.readFileSync("input.txt", "utf8");
  const lines = data.split("\n");
  const map = lines.map((line) => {
    return line.split("");
  });

  const guardLocation = findGuardLocation(map);
  while (progress(map, guardLocation)) {
    printMap(map, guardLocation);
    await new Promise((resolve) => setTimeout(resolve, 1000 / framesPerSecond));
  }

  console.log(uniqueLocationsCount, countUniqueLocations(map));
})();
