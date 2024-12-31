import { useState } from "react";
import { registerSingleFeed } from "~/lib/driver/register_single_feed";

// interface RegisterProps {
//   url: string;
// }

export const Register = () => {
  const [url, setUrl] = useState("");

  return (
    <div className="h-full w-full flex-col bg-white p-4 rounded-lg shadow-md">
      <div>
        <h1>Register feeds</h1>
        <p>Register your feeds here.</p>
      </div>
      <div className="flex flex-col w-1/3 justify-center">
        <form onSubmit={(e) => e.preventDefault()}>
          <label htmlFor="feedUrlInputField" className="mt-4">
            Feed URL
          </label>
          <input
            id="feedUrlInputField"
            type="url"
            className="border border-gray-300 mt-4 p-2 rounded-lg"
            placeholder="Enter feed URL"
            onChange={(e) => setUrl(e.target.value)}
          />
          <button
            id="registerButton"
            className="bg-indigo-400 text-white mt-4 p-2 rounded-lg"
            onClick={async () => {
              console.log("Register button clicked");
              const resJson = await registerSingleFeed(url);

              if (resJson) {
                alert("The feed has been registered successfully");
              }
            }}
          >
            Register
          </button>
        </form>
      </div>
    </div>
  );
};
