using System.Diagnostics;
using System.Text.Json;
using System.Xml.Linq;
using Microsoft.Extensions.Logging;

// Max retries for fetching a game from the API before giving up
const int MaxRetries = 3;
// Delay between retries in milliseconds
const int DelayBetweenRetries = 1000;
// Max degree of parallelism for fetching games (higer is better)
const int MaxDegreeOfParallelism = 5;
// Test on a smaller dataset to avoid hitting the API too hard (like 15 games instead of 160916)
const bool TestOnSmolData = false;

string line;
List<string> lines = [];
int lineCounter = 0;

using ILoggerFactory factory = LoggerFactory.Create(builder => builder.AddConsole());
ILogger logger = factory.CreateLogger("Doug");

var jsonSerializerOptions = new JsonSerializerOptions
{
    WriteIndented = true
};

try
{
    logger.LogInformation("Reading file...");
    StreamReader sr;
#pragma warning disable CS0162 // Unreachable code detected
    if (TestOnSmolData)
        sr = new("./input/boardgames_id_smol.txt");
    else
        sr = new("./input/boardgames_id.txt");
#pragma warning restore CS0162 // Unreachable code detected

    line = sr.ReadLine()!;
    while (line != null)
    {
        logger.LogDebug($"Read line {lineCounter}: {line}");
        lines.Add(line);
        line = sr.ReadLine()!;
        lineCounter++;
    }
    sr.Close();
}
catch (Exception e)
{
    logger.LogError(e, "Error reading file");
}
finally
{
    logger.LogInformation($"Read {lineCounter} lines");
}

using HttpClient client = new();
client.Timeout = TimeSpan.FromMinutes(10);
var games = new List<Game>();

var parallelOptions = new ParallelOptions
{
    MaxDegreeOfParallelism = MaxDegreeOfParallelism
};

var timer = Stopwatch.StartNew();

async Task<Game?> FetchGameWithRetry(string id, int attemptCount = 0)
{
    try
    {
        HttpResponseMessage response = await client.GetAsync($"https://boardgamegeek.com/xmlapi/boardgame/{id}");

        if (!response.IsSuccessStatusCode)
        {
            if (attemptCount < MaxRetries)
            {
                logger.LogWarning($"Retry {attemptCount + 1} for game {id}");
                await Task.Delay(DelayBetweenRetries);
                return await FetchGameWithRetry(id, attemptCount + 1);
            }
            logger.LogError($"Failed to fetch game {id} after {MaxRetries} attempts");
            return null;
        }

        string xmlContent = await response.Content.ReadAsStringAsync();
        XDocument doc = XDocument.Parse(xmlContent);
        var boardgame = doc.Descendants("boardgame").First();

        var game = new Game
        {
            Name = boardgame.Elements("name")
                .First(n => (string?)n.Attribute("primary") == "true")
                .Value,
            ImageUrl = boardgame.Element("image")?.Value ?? string.Empty,
            Genres = boardgame.Elements("boardgamecategory")
                .Select(g => g.Value)
                .ToList()
        };

        logger.LogInformation($"Successfully processed: {game.Name}");
        return game;
    }
    catch (Exception ex)
    {
        if (attemptCount < MaxRetries)
        {
            logger.LogWarning($"Error on attempt {attemptCount + 1} for game {id}: {ex.Message}");
            await Task.Delay(DelayBetweenRetries);
            return await FetchGameWithRetry(id, attemptCount + 1);
        }
        logger.LogError($"Failed to process game {id} after {MaxRetries} attempts: {ex.Message}");
        return null;
    }
}

var tasks = lines.Select(id => FetchGameWithRetry(id));
var results = await Task.WhenAll(tasks);
games.AddRange(results.Where(g => g != null).Cast<Game>());

logger.LogInformation($"Processed {games.Count} games in {timer.ElapsedMilliseconds / 1000.0 / 60.0} minutes");

// Write to file
logger.LogInformation("Writing to file...");
string json = JsonSerializer.Serialize(games, jsonSerializerOptions);
await File.WriteAllTextAsync($"./output/games.json", json);
logger.LogInformation("Done!");