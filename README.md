# Pcaphistogram

An implementation of https://github.com/joswr1ght/pcaphistogram/ to avoid having to return to Perl.

## Build

```
go build
```

## Usage

```
pcaphistogram <capture.pcap> | gnuplot
```

This will create an `output.png` file in the same directory which will have the
plot
