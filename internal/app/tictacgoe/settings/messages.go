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
var MatchAboutToStartMsg string = "Match's about to start in a %d x %d board. Press ANY KEY when ready!"

// During the match
var InitialBoardMsg string = "Here's a brand new board just for you..."
var AskNextMoveMsg string = "Mark the field at the desired position with a number: "

// Rejection for play moves
var RejectedtMoveBelowMinValueMsg string = "Oops! The lowest allowed spot value for your move is %d!"
var RejectedtMoveAboveMaxValueMsg string = "Oops! The highest allowed spot value for your move is %d!"
var RejectedMoveFieldPositionTaken string = "Oops! That field position %d is already taken, try another one!"

// Exiting the match
var ByeByeMsg string = "Thanks for Playing! Come again any time! The game will end in a few seconds..."
