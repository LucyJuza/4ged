public class Game
{
    public string Name { get; set; }
    public string ImageUrl { get; set; }
    public List<string> Genres { get; set; }

    public Game()
    {
        Name = string.Empty;
        ImageUrl = string.Empty;
        Genres = [];
    }
}