file="output/games.json"
miam=$(cat $file | jq length)
url=$1
lines=20000

if $url = "" ; then
    echo "You need to provide an URL to send the files to"
    exit 1
fi

echo "OwO what's dis?? *notices your tiny file* >w< It's onwy $miam lines wong!! Sooo kawaii!! 🥺💕"
sleep 2

echo "Heyyy bestieee!! 🌸 Let's split dis thicc file ($file) into wittle baby files!! Kay? [yn] ✨"
sleep 2

echo "Tee-hee! 🎀 Ur opinion doesn't matter anyway bestie!! We're doing it cuz it's gonna be soooo fun! uwu"
sleep 2

echo "How many lines do u want in each file, senpai?? (*＾▽＾)／"
sleep 2

echo "Oopsie woopsie!! Can't hear uuu! >w< We're doing $lines lines cuz I said so! *giggles* 🌟"
sleep 2

echo "OMG bestie!! I'm splitting into $((miam/lines+1)) adowable files of $lines lines each! *bounces excitedly* 💖"
sleep 2

for i in $(seq 0 $((miam/lines))); do
    echo "*splits file $i with extra sparkly love* ✨💕💫"
    cat $file | jq -c '.['$((i*lines))':'$((i*lines+lines))']' > output/games_$i.json
done

echo "Should we send to server-chan?? [yn] 👉👈"
sleep 2
echo "Hehe! *evil giggles* Sending everything to server-chan anyway cuz ur opinion is irrelevant bestie!! 🎀💅✨"
sleep 2

# Send files
for i in $(seq 0 $((miam/lines))); do
    echo "*sends file $i with lots of luv and sparkles* 🌸✨💖"
    curl -X POST -H "Content-Type: application/json" -o /dev/null -d @output/games_$i.json $url
done

# Delete files
rm output/games_*.json

echo "Kay byeee! Done sending everything! Don't need u anymore bestie! *blows kiss* 💅✨💖 Get out! uwu"