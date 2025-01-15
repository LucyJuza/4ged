#!/bin/bash

echo "✧･ﾟ: *✧･ﾟ:* 🌸 Hewwo World! 🌸 *:･ﾟ✧*:･ﾟ✧"
echo "I'm Cherry-Chan, your kawaii file-splitting assistant desu~! (◕‿◕✿)"
echo "Today we're gonna make your big files smol and kawaii together! *wiggles excitedly* 💕"

file="output/games.json"
miam=$(cat $file | jq length)
url=$1
lines=20000

echo "OwO what's dis?? *notices your file* The input file has $miam entries! Sugoi~! 💫"
echo "Bestieee!! Time to split $file into bite-sized chunks! Ready or not~ 🎀"
echo "*pulls out magical coding wand* We're setting each file to $lines lines! ✨"
echo "Quick maths time! *boops calculator* We'll make $((miam/lines+1)) kawaii baby files! 💕"

for i in $(seq 0 $((miam/lines))); do
    echo "*sprinkles magic coding dust on file $i* ✨"
    cat $file | jq -c '.['$((i*lines))':'$((i*lines+lines))']' > output/games_$i.json
done

echo "Yaaay~! Files split successfully! Time to send them to server-sama! (｡♥‿♥｡)"

# Send files
for i in $(seq 0 $((miam/lines))); do
    echo "Sending file $i with extra sparkles~! *nyaa* ⭐"
    curl -X POST -H "Content-Type: application/json" -o /dev/null -d @output/games_$i.json $url
done

# Delete files
rm output/games_*.json

echo "Mission accomplished! Your files have been sent with lots of luv! (◠‿◠✿)"
echo "Arigato for using Cherry-Chan's file splitter! Ja ne~! 🌸✨💖"