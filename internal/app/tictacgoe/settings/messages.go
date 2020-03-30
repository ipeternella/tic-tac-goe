// messages that are displayed to the user
package settings

// Intro
const WelcomeMsg string = `
+ - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +  
|                       WELCOME :)!                         |
|    TT | II | CC                                           |
|   --------------    ____   ____    ____                   |
|    TT | AA | CC       | i c  | a c  | o e                 |
|   --------------              Terminal 1.0                |
|    GG | OO | EE                                           |
|                                                           |
|   This program is licensed under the MIT License.         |
|   >> Press ANY KEY to begin                               |
+ - - - - - - - - - - - - - - - - - - - - - - - - - - - - - +
`

// Before the match
const MatchAboutToStartMsg string = "Match's about to start in a %d x %d board. Press ANY KEY when ready!"

// During the match
const InitialBoardMsg string = "Here's a brand new board just for you..."
const AskNextMoveMsg string = "Mark the field at the desired position with a number: "

// Rejection for play moves
const RejectedtMoveBelowMinValueMsg string = "Oops! The lowest allowed spot value for your move is %d!"
const RejectedtMoveAboveMaxValueMsg string = "Oops! The highest allowed spot value for your move is %d!"
const RejectedMoveFieldPositionTaken string = "Oops! That field position %d is already taken, try another one!"

// Victory
const VictoryMsg = "[WIN]: Congrats!!! %s player has won the match!!!"
const TieMatchMsg = "[TIE]: Too bad... no one has won! Better luck next time!"

// Exiting the match
const ByeByeMsg string = "Thanks for Playing! Come again any time! The game will end in a few seconds..."
