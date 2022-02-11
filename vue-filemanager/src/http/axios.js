import axios from 'axios'
export default axios.create({
    withCredentials: true, // 跨域请求时需要使用凭证
    timeout: 3000, // 指定3秒的请求超时
});