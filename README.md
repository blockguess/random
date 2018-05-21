# BlockGuess Lotto Winning Numbers lib

### Input
* random via Random.org generated
* blockhash bitcoin cash blockchain block hash

## Usage

* install golang
  you can visit https://golang.org/doc/install for install golang
  
* clone code 
 ```git
 git clone https://github.com/blockguess/random.git $GOPATH/src/blockguess-random
 ```
* build code
 ```go
 go build 
 ```
 
* run
You can run blockguess-random -h to check all options：
```
blockguess-random -h
  -blockhash string
    	BCH block hash (default "000000000000000001ed86134bcee0ad3f879f88e4cc3b27138d5c738de04fa9")
  -h	Print Help Info
  -random string
    	random.org random number (default "0123")
```
Use case example:
```
blockguess-random
```
output:
```
0072
```

* custom random number and blockhash
```
blockguess-random -random 66 -blockhash 000000000000000000e7689ce450f1a9f391402f2ccd71736fb8e0c7b23de0b6
```
output
```
0783
```


