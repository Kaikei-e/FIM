import { Welcome } from "../welcome/welcome";

export function meta() {
  return [
    { title: "Feed Insights Manager" },
    { name: "description", content: "Welcome to FIM!" },
  ];
}

export default function Home() {
  return <Welcome />;
}
