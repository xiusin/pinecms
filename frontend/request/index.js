import axios from 'axios'

const service = axios.create({
  baseURL: "http://localhost:2018/api/v1/",
  timeout: 10000
})

export default service
