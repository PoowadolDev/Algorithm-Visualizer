// Create layout for sorting page navigation bar
export default function Layout({ children }: { children: React.ReactNode }) {
    return (
    <>
        <div className="navbar bg-base-100">
            <a className="btn btn-ghost text-xl">daisyUI</a>
        </div>
        <div>{children}</div>
    </>
    )
}