// Copyright 2011 The Go Authors.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Generating random text: a Markov chain algorithm

Based on the program presented in the "Design and Implementation" chapter
of The Practice of Programming (Kernighan and Pike, Addison-Wesley 1999).
See also Computer Recreations, Scientific American 260, 122 - 125 (1989).

A Markov chain algorithm generates text by creating a statistical model of
potential textual suffixes for a given prefix. Consider this text:

	I am not a number! I am a free man!

Our Markov chain algorithm would arrange this text into this set of prefixes
and suffixes, or "chain": (This table assumes a prefix length of two words.)

	Prefix    Suffix

	"" ""     I
	"" I      am
	I am      a
	I am      not
	a free <span class="Apple-converted-space">      </span>man!
	am a <span class="Apple-converted-space">        </span>free
	am not <span class="Apple-converted-space">      </span>a
	a number!<span class="Apple-converted-space">    </span>I
	number! I<span class="Apple-converted-space">    </span>am
	not a<span class="Apple-converted-space">        </span>number!

To generate text using this table we select an initial prefix ("I am", for
example), choose one of the suffixes associated with that prefix at random
with probability determined by the input statistics ("a"),
and then create a new prefix by removing the first word from the prefix
and appending the suffix (making the new prefix is "am a"). Repeat this process
until we can't find any suffixes for the current prefix or we exceed the word
limit. (The word limit is necessary as the chain table may contain cycles.)

Our version of this program reads text from standard input, parsing it into a
Markov chain, and writes generated text to standard output.
The prefix and output lengths can be specified using the -prefix and -words
flags on the command-line.
*/
package wump

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
// Shift removes the first word from the Prefix and appends the given word.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Build reads text from the provided Reader and
// parses it into prefixes and suffixes that are stored in Chain.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

// Generate returns a string of at most n words generated from Chain.
func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return strings.Join(words, " ")
}
*/

const (
	MaxArrowShotDistance = 6		// +1 for '0' stopper
	MAXLinksInRoom = 25		// a complex cave
	MinLinksInRoom = 2
	MaxRoomsInCave = 250
	MinRoomsInCave = 10
)

const (		// levels of play
	EasyGame = iota
	HardGame = iota
)

// Debug is true when debugging information is output.
var Debug bool

// Chain contains a map ("chain") of prefixes to a list of suffixes.
// A prefix is a string of prefixLen words joined with spaces.
// A suffix is a single word. A prefix can have multiple suffixes.
type Room struct {
	tunnel []int
	hasPit bool
	hasBat bool
}

// Cave is a collecion of rooms.
type Cave []Room

// String returns the Prefix as a string (for use as a map key).
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewRoom(prefixLen int) *Room {
	return &Chain{make(map[string][]string), prefixLen}
}
	instructions();
	cave_init();

	for (;;) {
		initialize_things_in_cave();
		arrows_left = arrow_num;
		do {
			display_room_stats();
			(void)printf("Move or shoot? (m-s) ");
			(void)fflush(stdout);
			if (!fgets(answer, sizeof(answer), stdin))
				break;
		} while (!take_action());

    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
		if (!getans("\nCare to play another game? (y-n) "))
			exit(0);
		if (getans("In the same cave? (y-n) "))
			clear_things_in_cave();
		else
			cave_init();
	}

func plural(n int) string {
	if s < 1 {
		return "s"
	}
	return ""
}

// NewChain returns a new Chain with prefixes of prefixLen words.
func NewCave(pits int, bats int, rooms int, links int, arrows int) (Cave, error) {
	switch {
	case rooms < MinRoomsInCave:
		return nil, Errorf("No self-respecting wumpus would live in such a small cave!")
	case rooms > MaxRoomsInCave:
		return nil, Errorf("Even wumpii can't furnish caves that large!")
	case links < MinLinksInRoom:
		return nil, Errorf("Wumpii like extra doors in their caves!")
	case links > MAXLinksInRoom || links > rooms - (rooms / 4):
		return nil, Errorf("Too many tunnels!  The cave collapsed!\n(Fortunately, the wumpus escaped!)")
	case bats > rooms / 2:
		return nil, Errorf("The wumpus refused to enter the cave, claiming it was too crowded!")
	case pit_num > room_num / 2:
		return nil, Errorf("The wumpus refused to enter the cave, claiming it was too dangerous!")
	}
//	default:
//	return &Chain{make(map[string][]string), prefixLen}
}

func Wumpus(in io.Reader, out io.Writer) (bool, error) {
    




}
