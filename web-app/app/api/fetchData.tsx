'use client'

import axios from 'axios'


const endpoint = process.env.NEXT_PUBLIC_BACKEND_URL;
// const port = process.env.API_PORT;
// const prefix = process.env.API_SECURE == 'TRUE' ? 'https' : 'http';

async function fetchGetData(path: string) {
    console.log(`GET : ${endpoint}/${path}`)
    return axios.get(`${endpoint}/${path}`
        ).then((response) => {
            console.log(response.data)
            return response
})
}

async function fetchPostData(path: string, data: any) {
    console.log(`POST : ${endpoint}/${path}`)
    return axios.post(`${endpoint}/${path}`, data
        ).then((response) => {
            console.log(response.data)
            return response
    })
}

export { fetchGetData, fetchPostData }