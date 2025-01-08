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

echo "Terminé desu~ UwU! *fait un petit câlin à ton terminal* (ﾉ´ヮ\`)ﾉ*: ･ﾟ"