import Header from "@/components/Header/Header";


export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div>
      <header>
        <Header />
      </header>
      {children}
    </div>
  );
}
