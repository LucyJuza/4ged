// Utilities
import { v4 as uuidv4 } from 'uuid';
import { defineStore } from 'pinia'
import dummyUserDatas from '../../dummy-user-datas.json'
import dummyGames from '../../dummy-games.json'
import { getGame, getGames, getGenres, getUserData, login, postGame, postPerson, postPlay, register, searchGame } from '@/utils/requests';

export const useBackendStore = defineStore('backend', {
  state: () => ({
    userData: undefined,
    genres: [],
    games: [],
    filter: "",
    filteredUserGames: undefined,
    filteredGames: [],
    selectedGameIdForPlayCreation: undefined,
    selectedPlayIdForRepetition: undefined,
    userId: undefined,
    lastFilteredGamePageId: 0,
  }),
  actions: {
    async filterGames(searchFilter) {
      searchFilter = searchFilter.toLowerCase()
      this.lastFilteredGamePageId = 0;
      this.filter = searchFilter;
      this.filteredUserGames = [...this.userData.games].filter(g => g.name.toLowerCase().includes(searchFilter))
      this.filteredGames = await searchGame(searchFilter,0)
    },
    async getGameById(id) {
      return await getGame(id)
    },
    getGameHistory(gameId) {
      return this.userData.plays.filter(p => p.gameId == gameId)
    },
    getPlayById(id) {
      return this.userData.plays.find(p => p.id == id)
    },
    getPersonById(pId) {
      return this.userData.persons.find(p => p.id === pId)
    },
    setSelectedGameIdForPlayCreation(id) {
      this.selectedGameIdForPlayCreation = id
    },
    setSelectedPlayIdForRepetition(id) {
      this.selectedPlayIdForRepetition = id
    },
    getStatistics() {
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
    async addGame(values) {
      await postGame(this.userId,values.name,values.image,values.genres)
      await this.refreshDatas()
    },
    async addPerson(values) {
      await postPerson(this.userId,values.name,values.image)
      await this.refreshDatas()
    },
    async addPlay(values) {
      await postPlay(
        this.userId,
        values.gameId,
        values.date,
        values.location,
        values.duration,
        values.participants,
        values.winners
      )
      await this.refreshDatas()
    },
    async login(username,password) {
      const userId = await login(username,password)
      console.log("found user: " + userId)
      if (userId) {
        this.userId = userId
        const userData = await getUserData(userId)
        this.userData = userData;
        this.filteredUserGames = userData.games
      }
    },
    async register(username,password,image) {
      const userId = await register(username,password,image)
      const userData = await getUserData(userId)
      this.userId = userId
      this.userData = userData;
      this.filteredUserGames = userData.games
    },
    async pushNextFilteredGamesPage() {
      this.lastFilteredGamePageId += 1;
      const res = await searchGame(this.filter,this.lastFilteredGamePageId)
      this.filteredGames.push(...res)
    },
    async refreshDatas(){
      console.log("refreshing datas")
      this.userData = await getUserData(this.userId)
      this.lastFilteredGamePageId = 0;
      this.filteredGames = await searchGame(this.filter,this.lastFilteredGamePageId)
      this.filteredUserGames = [...this.userData.games].filter(g => g.name.toLowerCase().includes(this.filter))
    }
  }
})
// populate values that depend on backend (games and genres)
const populate = async() =>{
  const backendStore = useBackendStore()
  backendStore.genres = await getGenres()
  backendStore.games = await getGames(0)
  backendStore.filteredGames = backendStore.games
}
populate();