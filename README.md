# Tetris-Optimizer
[![Made with Go](https://img.shields.io/badge/Go-1-blue?logo=go&logoColor=white)](https://golang.org "Go to Go homepage")  
Tetris-Optimizer is a program that recieves a unique argument which is a text file that contains a list of tetrominos. The program reads the file and then assembles the tetrominos in order to create the smallest square possible.
In case of an incorect format the program should return the "ERROR" message.

## Installation
To install the program :
```bash
git clone https://zone01normandie.org/git/scointin/tetris-optimizer.git
```

## Usage

```bash
go run main.go src/examples/inputfile.txt
```
to run the test file: 
```bash
./test.sh
```
If you have trouble running the test file you can use this command:
```bash
chmod +x test.sh
```

## Tech
go version 1.23.0

Only standard library is used

## Author
S. Cointin
