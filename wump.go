package main

import (
	"fmt"
	"wumpus"
	"io"
)

const (
	DefRoomsInCave = 20
	DefLinksInRoom = 3
	DefNumberOfArrows = 5
	DefPitCount = 3
	DefBatCount = 3
)

func main() {
	out := os.Stdout
	in := os.Stdin

	// Register command-line flags.
	flag.BoolVar(&wumpus.Debug, "debug", false, "Run the game in debug mode")

	hardGame := flag.Bool("hard", false, "Play the game at the hard level of difficulty") // wumpus.EasyGame
	pits := flag.Int("pits", DefPitCount, "# pits in cave")
	bats := flag.Int("bats", DefBatCount, "# bats in cave")
	rooms := flag.Int("rooms", DefRoomsInCave, "# rooms in cave")
	links := flag.Int("links", DefLinksInRoom, "The number of links per room")
	arrows := flag.Int("arrows", DefNumberOfArrows, "The initial number of arrows in the inventory")

	flag.Parse()                     // Parse command-line flags.
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.

	if *hardGame {
		*bats += rand.Intn(rooms / 2) + 1
		*pits += rand.Intn(rooms / 2) + 1
	}


	// and we're OFF!  da dum, da dum, da dum, da dum...
	fmt.Fprintf(out, "\nYou're in a cave with %d rooms and %d tunnels leading from each room.\n\There are %d bat%s and %d pit%s scattered throughout the cave, and your\n\quiver holds %d custom super anti-evil Wumpus arrows.  Good luck.\n", *rooms, *links, *bats, plural(*bats), *pits, plural(*pits), *arrows);

	vic, err := Wumpus(in, out, *hardGame, *pits, *bats, *rooms, *links, *arrows)

	if err {
		fmt.Fprintln(out, err)
	}
	fmt.Printf("hello, world\n")
}
