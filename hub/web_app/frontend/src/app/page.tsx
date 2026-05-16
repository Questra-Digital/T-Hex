// src/app/page.tsx
import { redirect } from "next/navigation";

export default function Page() {
  // Redirect to the landing page route
  redirect("/landingPage");
}
