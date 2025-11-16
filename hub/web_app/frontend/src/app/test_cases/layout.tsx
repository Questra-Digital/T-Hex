import Header from "@/components/Header/Header";
import { montserrat, outfit } from "@/fonts/fonts";

export default function TestCasesLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div className={`${montserrat.variable} ${outfit.variable}`}>
      <Header />
      {children}
    </div>
  );
}
