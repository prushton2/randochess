import { Board } from "./Board";
import { Ruleset } from "./Ruleset";
import Team from "./Team";

export interface Game {
    ruleset: Ruleset,
    board: Board,
    turn: Team
}