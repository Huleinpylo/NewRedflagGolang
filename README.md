# NewRedflag
Basic Python script and go version to download/create wordlist from https://red.flag.domains/
Ref: https://red.flag.domains/details/ 
## Install
```python3
git clone https://github.com/lil-doudou/NewRedflag.git
cd NewRedFlag
pip3 install -r requirements.txt
```
``` Go version
For the go version
git clone https://github.com/Huleinpylo/NewRedflagGolang.git
cd NewRedflagGolang
go build -ldflags "-s -w" .\main.go

```
## Usage go
```
go  main.go
```
```
usage: main.go [-h] [--latest] [--day DAY] [--all]

options:
  -all              All list
  -day int          Day of list
  -latest           latest list
  -output string    Output file
  -verbose          Verbose output
```
## Building go
``` For compiling go 
go build -ldflags "-s -w" .\main.go
```

## Usage python
```
python3 newflag.py
```


```
usage: newflag.py [-h] [--latest] [--day DAY] [--all]

options:
  -h, --help         show this help message and exit
  --latest, -l       latest list
  --day DAY, -d DAY  Day of list
  --all, -A          All list
```




## Support
https://www.buymeacoffee.com/0xlildoudou
