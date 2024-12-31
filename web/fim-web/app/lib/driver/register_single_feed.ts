export async function registerSingleFeed(url: string) {
  const bffUrl = import.meta.env.VITE_BFF_URL;
  if (!bffUrl) {
    throw new Error("BFF_URL is not defined");
  }

  const res = await fetch(bffUrl, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      url: url,
    }),
  });

  if (res.ok) {
    return await res.json();
  } else {
    throw new Error("Failed to collect feed");
  }
}
