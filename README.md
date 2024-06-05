# version

[![Go Reference](https://pkg.go.dev/badge/github.com/bryanaustin/version/.svg)](https://pkg.go.dev/github.com/bryanaustin/version/)

Version string utility program.
Works with any version format. `1.2.3.4`, `v11.22`, `2024.05.08`, `1.2.3-4`
For more information about how parsing works, see [Parsing Section](#parsing)

## Install

Requires Go compiler: https://go.dev/dl/
```bash
go install github.com/bryanaustin/version/cmd/version@v1.1.0
```
This will install at `~/go/bin/version` by default.

## Library

This can also be used as a library if you have the need for such things.

## Usage

_Note: Examples put the options at the end for readability; they don't have to be in that order_

```bash
version [options] version
```

### --base
```bash
$ version 1.2.3-1 --base 1.2.3
1.2.3-2
```
This is the utility's main function, assuming you pass the last version as the argument and the release version for the `--base` option. If the base version matches the last version, it will increment the next highest version number. If it does not, it will adopt the `--base` version and reset all remaining numbers to zero. Separators ignored.

### --increment
```bash
$ version 1.2.3 --increment 0.5
1.7.0
```
Increment the version argument by the amount specified in the option. All smaller numbers will be reset to zero. Separators ignored.
Additional allowed values:
* `major`, same as `1`
* `minor`, same as `0.1`
* `patch`, same as `0.0.1`
* `package`, same as `0.0.0.1`

### --set
```bash
$ version 1.2.3 --set 0.5
1.5.0
```
Set the version argument to the non-zero values supplied. All smaller numbers will be reset to zero. Separators ignored.

### --minimum
```bash
$ version 1.2.3 --minimum 0.5
1.5.3
```
Enforce a minimum value for the numbers provided. Separators ignored.

### --format
```bash
$ version 1.2.3 --format v0-0_1
v1-2_3
```
Return the provided version argument in the format of the option. Numbers ignored.

### --pad
```bash
$ version 5.6.7 --pad 1.2.3
5.06.007
```
Ensure the numbers are zero-padded by the amount specified. Separators ignored.

### --greaterthan
```bash
$ version 1.10 --greaterthan 1.9
true
```
Tests to see if the version argument is greater than the option value. It cannot be used with any other arguments. Separators ignored.

### --lesserthan
```bash
$ version 1.9 --lesserthan 1.10
true
```
Tests to see if the version argument is greater than the option value. It cannot be used with any other arguments. Separators ignored.

## Operation

### Parsing

Version strings are parsed into arrays of sequential numbers and non-numbers and re-assembled for output. In code, the non-numbers are called separators. Options will often ignore the separators and operate on the provided numbers, so any separators can be used. The exception is the `--format` option, which only uses the separators and ignores the numbers.

### Order of operations

All options except for `--greaterthan` & `--lesserthan` can chained together for the combined result. The order of operations is as follows:
* `--base`
* `--increment`
* `--set`
* `--minimum`

Options cannot be defined multiple times, and the order they are provided in doesn't matter.
