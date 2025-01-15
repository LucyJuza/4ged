const BACKEND_URL_PREFIX = "https://4ged.chacal.beer"
const POSTheaders = new Headers()
POSTheaders.append("content-type", "application/json")
const getGenres = async () => {
  return (await fetch(`${BACKEND_URL_PREFIX}/genres`)).json()
}
const getGames = async (pageId) => {
  return (await fetch(`${BACKEND_URL_PREFIX}/games?pageIndex=${pageId}`)).json()
}
const getUserData = async (id) => {
  return (await fetch(`${BACKEND_URL_PREFIX}/users/${id}`)).json()
}
const register = async (username,password,image) => {
  return (await 
    fetch(`${BACKEND_URL_PREFIX}/register`,
      {
        method: "POST", 
        headers: POSTheaders, 
        body: JSON.stringify({
          username: username,
          password: password,
          image: image
        })
      }
    )
  ).json()
}
const login = async (username,password) => {
  return (await 
    fetch(`${BACKEND_URL_PREFIX}/login`,
      {
        method: "POST", 
        headers: POSTheaders, 
        body: JSON.stringify({
          username: username,
          password: password,
        })
      }
    )
  ).json()
}
const getGame = async (id) =>{
  return (await fetch(`${BACKEND_URL_PREFIX}/games/${id}`)).json()
}
const searchGame = async (filter,pageId) =>{
  return (await fetch(`${BACKEND_URL_PREFIX}/games/search?filter=${filter}&pageIndex=${pageId}`)).json()
}
const postGame = async (userId,name,image,genres) => {
  return (await 
    fetch(`${BACKEND_URL_PREFIX}/users/${userId}/games`,
      {
        method: "POST", 
        headers: POSTheaders, 
        body: JSON.stringify({
          name: name,
          image: image,
          genres: genres
        })
      }
    )
  ).json()
}
const postPerson = async (userId, name, image) =>{
  return (await 
    fetch(`${BACKEND_URL_PREFIX}/users/${userId}/persons`,
      {
        method: "POST", 
        headers: POSTheaders, 
        body: JSON.stringify({
          name: name,
          image: image,
        })
      }
    )
  ).json()
}
const postPlay = async (userId, gameId, date, location, duration, participants, winners) => {
  return (await 
    fetch(`${BACKEND_URL_PREFIX}/users/${userId}/plays`,
      {
        method: "POST", 
        headers: POSTheaders, 
        body: JSON.stringify({
          gameId: gameId,
          date: date,
          location: location,
          duration: duration,
          participants: participants,
          winners: winners
        })
      }
    )
  ).json()
}
export {getGenres, getGames, getGame, postGame, searchGame, register, login, getUserData, postPerson, postPlay}