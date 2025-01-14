// ⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️
// CECI EST UNE IMPLEMENTATION NON SECURISÉE (purement frontend)
// ELLE EST LA POUR L'EXEMPLE
// ET NE DEVRAIT PAS "VRAIMENT" ÊTRE UTILISÉE
// ⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️
import { v4 as uuidv4 } from 'uuid';
const auths = [{
  username: "Raptiste Boland",
  password: "Pa$$w0rd",
  userId: "test_player"
}]

const login = (username, password) => {
  const result = auths.find( u => (u.username === username && u.password === password) )?.userId
  return result
}
const register = (username,password) => {
  const id = uuidv4()
  auths.push({
    userId: id,
    username: username,
    password: password
  })
  return id
}
export { login, register } 