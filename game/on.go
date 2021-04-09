package game

import (
	"context"
	"fmt"
	"io"
	"time"
)

type gamestate int

const (
	SHOW_SPLASH gamestate = iota
	GET_PLAYERS
	CHOOSE_FIRST_PLAYER
	START_GAME
	START_ROUND
	NEXT_PLAYER
	START_TURN
	ROLL_DICE
	DO_ACTION
	END_TURN
	END_ROUND
	END_GAME
)

func Next(ctx context.Context, w io.Writer) {
	t := ctx.Value("time").(time.Time)
	fmt.Fprintf(w, "Hello Talman: t -> %v", t)
}
