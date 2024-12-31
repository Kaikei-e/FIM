import { Register } from "~/components/islands/register";

export default function Home() {
  return (
    <div className="h-screen w-screen p-8 bg-slate-200 flex flex-col">
      <div className="flex flex-col">
        <h1 id="fim-home-title" className="text-2xl font-semibold">
          Islands
        </h1>
        <p>
          Islands is based on the multiple components that have different
          functinalities.
        </p>
        <p>This is the dashboard for your insights.</p>
      </div>
      <div className="h-full w-full p-4 grid grid-cols-1 gap-4 sm:grid-cols-2">
        <div>
          <p>Search feeds</p>
        </div>
        <div className="h-full w-full">
          <Register />
        </div>
        <div>
          <p>Latest feeds</p>
        </div>
        <div>
          <p>Today&apos;s topics</p>
        </div>
      </div>
    </div>
  );
}
