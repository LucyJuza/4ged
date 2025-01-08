file="output/games.json"

miam=$(cat $file | jq length)
echo "OwO qu'est-ce que c'est ? *notices ton fichier fait ${miam} lignes* (｡♥‿♥｡)"

read -p "Nee nee~ On split le fichier ensemble ? ($file) [yn] uwu " yn

if [ "$yn" != "${yn#[Yy]}" ] ;then 
    echo "Yaaaaay! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧"
else
    echo "Owww... *sad gay noises* (｡•́︿•̀｡)"
    exit
fi

read -p "Combien de lignes par fichier, senpai ? (◕‿◕✿) " lines

echo "Nyaa~ Je vais split en $((miam/lines+1)) fichiers de $lines lignes chacun! (〜￣▽￣)〜"

for i in $(seq 0 $((miam/lines))); do
    echo "*split le fichier $i avec amour* ♡(◡‿◡✿)"
    cat $file | jq -c '.['$((i*lines))':'$((i*lines+lines))']' > output/games_$i.json
done

read -p "On envoie tout ça sur le serveur ? [yn] (´｡• ω •｡\`) " yn

if [ "$yn" != "${yn#[Yy]}" ] ;then 
    echo "Yaaaaay! (ﾉ◕ヮ◕)ﾉ*:･ﾟ✧"
else
    echo "Owww... *"
    exit
fi

# Get URL
read -p "Quelle est l'URL du serveur ? (´｡• ω •｡\`) " url

# Send files
for i in $(seq 0 $((miam/lines))); do
    echo "*envoie le fichier $i avec amour* ♡(◡‿◡✿)"
    curl -X POST -H "Content-Type: application/json" -o /dev/null -d @output/games_$i.json $url
done

# Delete files
rm output/games_*.json

echo "Terminé desu~ UwU! *fait un petit câlin à ton terminal* (ﾉ´ヮ\`)ﾉ*: ･ﾟ"