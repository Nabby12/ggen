#!/bin/bash

latestTag=$(git describe --abbrev=0 --tags)
cmdStatus=$?

if [ $cmdStatus -ne 0 ]; then
    echo "v0.0.1"
    exit 0
fi

latestTag=$(echo ${latestTag} | awk '{print substr($0, 2)}')

IFS=. tagArray=(${latestTag})
major=${tagArray[0]}
minor=${tagArray[1]}
revision=${tagArray[2]}

if [ ${revision} -ge 9 ]; then
    if [ ${minor} -ge 9 ]; then
        major=$((major+1))
        minor=0
        revision=0
    else
        minor=$((minor+1))
        revision=0
    fi
else
  revision=$((revision+1))
fi

echo "v${major}.${minor}.${revision}"
