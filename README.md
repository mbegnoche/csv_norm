# CSV Normalization Example Work

## Build
- Clone the repository into a subdirectory of your GO Workspace
- Build using the GO Compiler v1.10.2 
```bash
go build
```

## Run
To run the way that was suggested in the example problem:
```bash
cat sample.csv | ./csv
```
To run using a command line argument for the input:
```bash
./csv -input sample.csv
```
To run using command line arguments for input and write to an output file:
```bash
./csv -input sample.csv -output norm.csv
```
