using System.Net.Http;

var client = new HttpClient();

HttpResponseMessage response = await client.GetAsync (
    "https://boardgamegeek.com/xmlapi/boardgame/224517");


HttpContent responseContent = response.Content;

using (var reader = new StreamReader(await responseContent.ReadAsStreamAsync()))
{
    Console.WriteLine(await reader.ReadToEndAsync());


        // Create a file to write to.
    string createText = "Hello and Welcome" + Environment.NewLine;
    File.WriteAllText(path, createText);

    ...

    // Open the file to read from.
    string readText = File.ReadAllText(path);
}