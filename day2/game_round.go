package main

import (
	"strings"
)

type Move int 

const (
    Rock = 0
    Paper = 1
    Scissors = 2
)

type GameRound struct {
    myMove Move
    opponentMove Move
}

func game_round_from_str(game_round_str string, is_part_2 bool) GameRound {
    var round GameRound

    grStrSplit := strings.Split(game_round_str, " ")

    opponentMap := map[string]Move {
        "A": Rock,
        "B": Paper,
        "C": Scissors,
    }

    myMap1 := map[string]Move {
        "X": Rock,
        "Y": Paper,
        "Z": Scissors,
    }

    myMap2 := map[string]Move {
        "X": 2, // Loose
        "Y": 0, // Draw
        "Z": 1, // Win
    }

    round.opponentMove = opponentMap[grStrSplit[0]] 

    if !is_part_2 {
        round.myMove = myMap1[grStrSplit[1]] 
    } else {
        round.myMove = (round.opponentMove + myMap2[grStrSplit[1]]) % 3
    }

    return round
}

func score(round GameRound) int {
    shapeScoreMap := map[Move]int { 
        Rock: 1,
        Paper: 2,
        Scissors: 3,
    }

    scoreGrid := [][]int {
        {3, 0, 6},
        {6, 3, 0},
        {0, 6, 3},
    }

    return shapeScoreMap[round.myMove] + scoreGrid[round.myMove][round.opponentMove]
}
