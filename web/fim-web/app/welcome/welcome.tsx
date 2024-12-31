import React from "react";
import "./styles.css";

export function Welcome() {
  return (
    <main className="relative">
      <div className="background-image"></div>
      <div className="ripple"></div>
      <div className="ripple"></div>
      <div className="content flex flex-col items-center justify-center min-h-screen header-text">
        <header className="flex flex-col items-center gap-9 pt-16 pb-4">
          <div className="w-[500px] max-w-[100vw] p-4 justify-items-start">
            <h1 id="fim-title" className="text-3xl">
              Feed Insights Manager
            </h1>
            <p className="text-xl">Welcome to FIM!</p>
            <ul className="mt-4 list-disc list-inside pl-4">
              <li>Manage your feed data</li>
              <li>Get insights from your feed data</li>
              <li>Visualize your feed data</li>
            </ul>
          </div>
        </header>
        <div className="absolute bottom-0 left-0 right-0 flex justify-center p-4">
          <button id="get-started">
            <a
              href="/home"
              className="bg-slate-800 text-white px-4 py-2 rounded-md"
            >
              Get Started
            </a>
          </button>
        </div>
      </div>
    </main>
  );
}
