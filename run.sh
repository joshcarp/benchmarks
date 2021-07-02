#!/bin/bash

for dir in */
do
  echo "$dir"
  find $dir  -type f -exec grep "func Benchmark" {} \; | sed 's/func //' | sed 's/(b \*testing\.B) {//' | while read test; do
    echo "Processing file '$dir $test'"
    go test ./$dir -bench=. -run BenchmarkOptionMap
  done
done
