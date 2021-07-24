#!/bin/bash

files=(*/)
if [ $1 != "" ];
then
  files=( "$1" )
  echo ${files[@]}
fi

for dir in "${files[@]}";
do
  if test -d $dir
    then
    echo "running $dir"
    dir=${dir%?};
    echo "# $dir" > $dir/README.md
    echo "\`\`\`bash" >> $dir/README.md
    go test ./$dir -v -bench=. -benchmem -cpuprofile $dir/cpu.prof -memprofile $dir/mem.prof >> $dir/README.md
    echo "\`\`\`" >> $dir/README.md
    go tool pprof -svg -output=$dir/cpu.svg $dir/cpu.prof && rm $dir/cpu.prof
    go tool pprof -svg -output=$dir/mem.svg $dir/mem.prof && rm $dir/mem.prof
    echo "## Memory profile" >> $dir/README.md
    echo "![](mem.svg)" >> $dir/README.md
    echo "## CPU profile" >> $dir/README.md
    echo "![](cpu.svg)" >> $dir/README.md
  fi
done
