# Tic Tac Goe - A terminal tic tac toe game

```bash
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
```

A terminal Tic Tac Toe game implemented in GO to discover the language!

## Terminal Example

The default board size is `3x3` but supports up to `9x9` boards!

```bash
 00 | 01 | 02 
--------------
 03 | 04 | 05 
--------------
 06 | 07 | 08 
```

A `6x6` board example:

```bash
 00 | 01 | 02 | 03 | 04 | 05 
-----------------------------
 06 | 07 | 08 | 09 | 10 | 11 
-----------------------------
 12 | 13 | 14 | 15 | 16 | 17 
-----------------------------
 18 | 19 | 20 | 21 | 22 | 23 
-----------------------------
 24 | 25 | 26 | 27 | 28 | 29 
-----------------------------
 30 | 31 | 32 | 33 | 34 | 35 
```

## Code linting and formatting

This project uses Go's own formatting tool. Hence, to format the code:

```bash
go fmt ./...  # formats all `*.go` files found in the project
```

## Running the project

To run the project, run

```bash
make
```

Which will run the `Makefile` with the `go run` option. To compile the project to a binary:

```bash
make build
```

## Running tests

To run the tests, run the following command:

```bash
make test
```

Which will scan for the project's `*_test.go` files and execute them.
