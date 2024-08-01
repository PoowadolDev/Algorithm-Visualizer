import Link from "next/link"

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
    <>
        <div className="navbar bg-slate-700">
            <a className="btn btn-ghost text-3xl">Sorting</a>
            <Link href="/" className="font-normal ps-4 text-xl">Home</Link>
            <a className="font-normal ps-6 text-xl">Random</a>
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
        </div>
        <div>{children}</div>
    </>
    )
}