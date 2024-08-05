'use server'

import axios from 'axios'


const endpoint = process.env.API_ENDPOINT
const port = process.env.API_PORT
const prefix = process.env.API_SECURE == 'TRUE' ? 'https' : 'http'

async function fetchGetData(path: string) {
    const response = await axios.get(`${prefix}://${endpoint}:${port}/${path}`)

    return response.data
}

export { fetchGetData }