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
            <h1 className="text-3xl">Feed Insights Manager</h1>
            <p className="text-xl">Welcome to FIM!</p>
            <ul className="mt-4 list-disc list-inside pl-4">
              <li>Manage your feed data</li>
              <li>Get insights from your feed data</li>
              <li>Visualize your feed data</li>
            </ul>
          </div>
        </header>
      </div>
    </main>
  );
}
