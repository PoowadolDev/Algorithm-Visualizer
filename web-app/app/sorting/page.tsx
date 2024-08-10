'use client';

import Link from "next/link";
import { useEffect, useState } from "react";
import { fetchGetData } from "../api/fetchData";
import { HistogramSortChart } from "../components/sortChart";

export default function Page() {
    const [size, setSize] = useState(7);
    const [trackSize, setTrackSize] = useState(7);
    const [genData, setGenData] = useState<number[]>([]);

    const fetchGen = async () => {
        const response = await fetchGetData(`generateData?size=${size}`)
        setGenData(response.data);
    }

    useEffect(() => {
        fetchGen();
    }, [size]);

    const handleSizeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setTrackSize(Number(event.target.value));
      };

    return (
    <>
        <div className="navbar bg-slate-700">
            <a className="btn btn-ghost text-3xl">Sorting</a>
            <Link href="/" className="font-normal ps-4 text-xl">Home</Link>
            <div className="flex-none">
                <ul className="menu menu-horizontal">
                    <li>
                        <details>
                            <summary className="text-xl">Algorithm</summary>
                                <ul className="bg-base-100 rounded-t-none w-36 font-semibold">
                                    <li><a>Selection Sort</a></li>
                                    <li><a>Merge Sort</a></li>
                                    <li><a>Insertion Sort</a></li>
                                    <li><a>Bubble Sort</a></li>
                                </ul>
                        </details>
                    </li>
                </ul>
            </div>
            <form>
                <button type="button" onClick={() => {
                    setSize(trackSize)
                }} className="font-normal text-xl pe-10">Generate Data</button>
                <input type="range" name="size" min={0} max={100} value={trackSize} onChange={handleSizeChange} id="size"/>
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