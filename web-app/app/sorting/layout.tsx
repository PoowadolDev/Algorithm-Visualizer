import Link from "next/link"
import { GenData } from "./action"
import { sign } from "crypto"

export default function Layout({ children }: { children: React.ReactNode }) {
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
            <form onSubmit={GenData}>
                <button type="submit" className="font-normal text-xl pe-10">Generate Data</button>
                <input type="range" name="size" defaultValue={10} min={0} max={100} id="size"/>
            </form>
        </div>
        <div>{children}</div>
    </>
    )
}