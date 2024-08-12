'use client';

import Link from "next/link";
import { useEffect, useState } from "react";
import { fetchGetData, fetchPostData } from "../api/fetchData";
import { HistogramSortChart } from "../components/sortChart";

export default function Page() {
    const [size, setSize] = useState(7);
    const [trackSize, setTrackSize] = useState(7);
    const [genData, setGenData] = useState<number[]>([]);
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
        const response = await fetchPostData("sortProblem", data)
        console.log(response.data)
    }

    useEffect(() => {
        fetchGen();
    }, [size]);

    // useEffect(() => {
    //     fetchSolve();
    // }, [sortAlgorithm])

    const handleSizeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setTrackSize(Number(event.target.value));
      };

    return (
    <>
        <div className="navbar bg-slate-700 text-white">
            <a className="btn btn-ghost text-3xl">Sorting</a>
            <Link href="/" className="font-normal ps-4 text-xl">Home</Link>
            <div className="flex-none">
                <ul className="menu menu-horizontal">
                    <li>
                        <details>
                            <summary className="text-xl">{sortAlgorithm}</summary>
                                <ul className="bg-slate-700 rounded-t-none w-36 font-semibold">
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Selection")}><a>Selection Sort</a></li>
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Merge")}><a>Merge Sort</a></li>
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Insertion")}><a>Insertion Sort</a></li>
                                    <li className="hover:bg-slate-600" onClick={() => setSortAlgorithm("Bubble")}><a>Bubble Sort</a></li>
                                </ul>
                        </details>
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
        <div className="flex justify-center items-center h-screen">
            <div className="h-full w-9/12">
                <HistogramSortChart dataInput={genData}/>
            </div>
        </div>

    </>
    )
}