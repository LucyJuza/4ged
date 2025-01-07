public class Game
{
  public string Name { get; set; }
  public string ImageUrl { get; set; }
  public List<Genre> Genres { get; set; }

  public Game()
  {
    Name = string.Empty;
    ImageUrl = string.Empty;
    Genres = [];
  }
}

public class Genre
{
  public string Name { get; set; }

  public Genre()
  {
    Name = string.Empty;
  }
}