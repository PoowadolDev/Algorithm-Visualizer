'use client'

import { useState, useEffect, FormEvent } from "react";
import { fetchGetData } from "../api/fetchData";

function GenData (event: FormEvent<HTMLFormElement>) {

    const formData = new FormData(event.currentTarget);
    const size = formData.get('size');
    console.log(size)
    const [genData, setGenData] = useState([]);
    useEffect(() => {
        fetchGetData(`generateData?size=${size}`).then((data) => setGenData(data))
    }, []);

    return genData
}

export { GenData }