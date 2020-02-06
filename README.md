# namegrind

This utility returns sunrise and availability data for [Handshake](https://handshake.org) names.

## Installation

1. Download the binary for your OS.
2. Move it somewhere on your `$PATH`, like this: `sude mv namegrind-darwin-x64 /usr/local/bin/namegrind`. You can replace `namegrind-darwin-x64` with the filename and location of the binary you downloaded.

## Usage

First, create an input file. The input file is simply a newline-delimited list of names, like this:

```
name1.
name2.
name3.
```

Then, run `namegrind <infile>`, where `<infile>` is replaced with the location of the input file you created above.

You'll see output that looks like this:

```
honk,15120,13,false
beep,17136,15,false
test,6048,4,false
farb,38304,36,false
1and1,49392,47,true
facebook,32256,30,true
```

These fields correspond to `<name>,<available height>,<available week>,<reserved>`. The output can be redirected to a CSV file for easy import into other programs.