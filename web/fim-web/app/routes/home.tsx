import type { Route } from "./+types/home";
import { Welcome } from "../welcome/welcome";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Feed Insights Manager" },
    { name: "description", content: "Welcome to FIM!" },
  ];
}

export default function Home() {
  return <Welcome />;
}
