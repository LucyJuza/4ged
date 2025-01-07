// Utilities
import { defineStore } from 'pinia'
import dummyUserDatas from '../../dummy-user-datas.json'
import dummyGenres from '../../dummy-genres.json'
export const useDummyStore = defineStore('app', {
  state: () => ({
    userData: dummyUserDatas,
    genres: dummyGenres
  }),
})
