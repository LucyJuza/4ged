// Utilities
import { v4 as uuidv4 } from 'uuid';
import { defineStore } from 'pinia'
import dummyUserDatas from '../../dummy-user-datas.json'
import dummyGenres from '../../dummy-genres.json'
import dummyGames from '../../dummy-games.json'
export const useDummyStore = defineStore('app', {
  state: () => ({
    userData: {
      "id" : dummyUserDatas.id,
      "personId": dummyUserDatas.personId,
      "name": dummyUserDatas.name,
      "image": dummyUserDatas.image,
      "games": dummyUserDatas.games,
      "plays": dummyUserDatas.plays,
      "persons": dummyUserDatas.persons
    },
    genres: dummyGenres,
    games: dummyGames,
    filteredUserGames: dummyUserDatas.games,
    filteredGames: dummyGames,
    selectedGameIdForPlayCreation: undefined,
    selectedPlayIdForRepetition: undefined
  }),
  actions: {
    filterGames(searchFilter){
      searchFilter = searchFilter.toLowerCase()
      this.filteredUserGames = [...this.userData.games].filter(g => g.name.toLowerCase().includes(searchFilter))
      this.filteredGames = [...this.games].filter(g => g.name.toLowerCase().includes(searchFilter))
    },
    getGameById(id){
      return this.games.find(g => g.id == id)
    },
    getGameHistory(gameId){
      return this.userData.plays.filter(p => p.gameId == gameId)
    },
    getPlayById(id){
      return this.userData.plays.find(p => p.id == id)
    },
    getPersonById(pId){
      return this.userData.persons.find(p => p.id === pId)
    },
    setSelectedGameIdForPlayCreation(id){
      this.selectedGameIdForPlayCreation = id
    },
    setSelectedPlayIdForRepetition(id){
      this.selectedPlayIdForRepetition = id
    },
    getStatistics(){
      const playtime = this.userData.plays.reduce((acc,curr) => acc+= curr.duration,0)
      const nPlays = this.userData.plays.length
      const nGames = this.userData.games.length
      const nPersons = this.userData.persons.length

      const winners = this.userData.plays.map((play) => ({
        winners: play.winners
      }))
      const won = (pid) => winners.filter((w) => w.winners.includes(pid)).length
      const played = (pid) => this.userData.plays.filter((play) => play.participants.includes(pid)).length
      const winRateByPlayer = this.userData.persons.map(p => ({
        person: p.id,
        won: won(p.id),
        played: played(p.id),
        rate: won(p.id)/played(p.id) 
      }))
      const mostWins = [...winRateByPlayer].sort((a,b) => b.won - a.won)[0];
      const highestWinrate = [...winRateByPlayer].sort((a,b) => b.rate - a.rate)[0];
      const personalWinrate = {nGames: played(this.userData.personId), won: won(this.userData.personId), lost: played(this.userData.personId) - won(this.userData.personId)}
      return {
        playtime: playtime,
        plays: nPlays,
        games: nGames,
        persons: nPersons,
        mostWins: mostWins,
        highestWinrate: highestWinrate,
        personalWinrate: personalWinrate
      }
    },
    addGame(values){
      const id =  uuidv4()
      this.userData.games.push({
        id: id,
        name: values.name,
        genres: values.genres,
        image: values.image
      })
      this.games.push({
        id: id,
        name: values.name,
        genres: values.genres,
        image: values.image
      })
    },
    addPerson(values){
      const id =  uuidv4()
      this.userData.persons.push({
        id: id,
        name: values.name,
        image: values.image,
        winrate: 0
      })
    },
    addPlay(values){
      const id =  uuidv4()
      this.userData.plays.push({
        id: id,
        gameId: values.gameId,
        date: values.date,
        location: values.location,
        duration: values.duration,
        participants: values.participants,
        winners: values.winners
      })
    }
  }
})
