import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div className="flex items-center justify-center h-screen">
      <div className="flex flex-col items-center justify-center">
        <h1 className="text-3xl font-bold pb-4">Sorting Algorithm</h1>
        <Link className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" href="/sorting">
          Go !!!
        </Link>
      </div>
    </div>
  );
}
