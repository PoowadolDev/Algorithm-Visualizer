'use client';

import Link from "next/link";
import { useEffect, useState } from "react";
import { fetchGetData, fetchPostData } from "../api/fetchData";
import { HistogramSortChart } from "../components/sortChart";
import axios, { AxiosResponse } from "axios";



export default function Page() {
    const endpoint = process.env.NEXT_PUBLIC_BACKEND_URL;

    const [size, setSize] = useState(7);
    const [trackSize, setTrackSize] = useState(7);
    const [genData, setGenData] = useState<number[]>([]);
    const [solveData, setSolveData] = useState<SolveData[]>([]);
    const [sortAlgorithm, setSortAlgorithm] = useState<string>("Selection");


    const fetchGen = async () => {
        const response = await fetchGetData(`generateData?size=${size}`)
        setGenData(response.data);

    }

    const fetchSolve = async () => {
        const data = {
            "sortType": sortAlgorithm,
            "dataList": genData
        }
        try {
            const response: AxiosResponse<SolveData[]> = await axios.post<SolveData[]>(
                `${endpoint}/sortProblem`, data);
            setSolveData(response.data);
        } catch (error) {
            throw new Error(`Failed to fetch user data: ${error}`);
        }
    }

    const delay = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

    const playSolve = async () => {
        for (const item of solveData) {
            setGenData(item.DataList)
            await delay(1000);
        }
    }

    useEffect(() => {
        fetchGen();
    }, [size]);

    useEffect(() => {
        playSolve();
    }, [solveData]);

    const handleSizeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setTrackSize(Number(event.target.value));
      };

    return (
    <div className="relative">
        <div className="Static navbar bg-slate-700 text-white">
            <a className="btn btn-ghost text-3xl">Sorting</a>
            <Link href="/" className="font-normal ps-4 text-xl">Home</Link>
            <div className="flex-none">
                <ul className="menu menu-horizontal">
                    <li>
                    <div className="dropdown dropdown-bottom text-white">
                            <div tabIndex={0} role="button" className="text-xl  bg-slate-700">{sortAlgorithm}</div>
                                <ul tabIndex={0} className="dropdown-content menu  bg-slate-700 z-[1] w-52 p-2 font-semibold">
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Selection")}><a>Selection Sort</a></li>
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Merge")}><a>Merge Sort</a></li>
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Insertion")}><a>Insertion Sort</a></li>
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Bubble")}><a>Bubble Sort</a></li>
                                </ul>
                        </div>
                    </li>
                </ul>
            </div>
            <form>
                <button type="button" onClick={() => {
                    setSize(trackSize)
                }} className="font-normal text-xl pe-10">Generate Data</button>
                <input type="range" className="pe-10" name="size" min={0} max={100} value={trackSize} onChange={handleSizeChange} id="size"/>
                <button type="button" onClick={() => {
                    fetchSolve()
                }} className="font-normal text-xl ms-10 p-2 rounded hover:bg-slate-600">Solve</button>
            </form>
        </div>
        <div className="flex justify-center items-center h-96">
            <div className="h-full w-9/12">
                <HistogramSortChart dataInput={genData}/>
            </div>
        </div>

    </div>
    )
}