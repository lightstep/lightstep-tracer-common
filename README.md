# LightStep Proto Definitions

## Requirements
* protoc

## Manual Usage
Run `protoc` in this directory, targeting the specific proto file and output type desired. For example -

Java
```
export SRC_DIR=/path/to/lightstep-tracer-common
export DST_DIR=/path/to/lightstep-tracer-java-common
protoc -I=$SRC_DIR --java_out=$DST_DIR $SRC_DIR/collector.proto
```