#!/bin/bash

echo Testing sizes "2^$1 to 2^$2 in ascending powers of 2 (inclusive)"

cmd=${@:4}
echo Command to run "$cmd"

declare -i start end i j loops
start=$1
end=$2
loops=$3

echo Loop iterations "$loops"

for ((j=start;j<=end;j++)); do
    size=$((2 ** j))
    echo "size = $size"
    $cmd -threshold 0 -loops "$loops" -size $size -mode benchmark | \
    tee "${RESULTFOLDER}/${size}_0_${loops}.txt"

    for ((i=0;i<=j;i++)); do
        thold=$((2 ** i))
        echo -e "\tthreshold = $thold"
        $cmd -threshold $thold -loops "$loops" -size $size -mode benchmark | \
        tee "${RESULTFOLDER}/${size}_${thold}_${loops}".txt
    done
done
