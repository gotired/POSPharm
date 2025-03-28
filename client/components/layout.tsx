import { useRouter } from "next/router";
import Navbar from "./navbar";
import SideBar from "./sidebar";
import { ReactNode } from "react";

interface LayoutProps {
  children: ReactNode;
}

export default function Layout({ children }: LayoutProps) {
  const { pathname } = useRouter();
  return (
    <div className="h-full flex">
      <SideBar pathname={pathname} />
      <div className="flex flex-col w-full">
        <Navbar />
        {children}
      </div>
    </div>
  );
}
