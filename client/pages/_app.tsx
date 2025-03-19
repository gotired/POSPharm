import type { AppProps } from "next/app";
import "../app/globals.css";
import SideBar from "@/components/sidebar";
import { useRouter } from "next/router";
import NavBar from "@/components/navbar";
import { IBM_Plex_Sans_Thai_Looped, Roboto_Flex } from "next/font/google";

const ibm_plex_sans_thai_loop = IBM_Plex_Sans_Thai_Looped({
  weight: ["100", "200", "300", "400", "500", "600", "700"],
  subsets: ["latin", "thai"],
  variable: "--font-ibm-plex-sans-thai-loop",
});

// const roboto = Roboto_Flex({ subsets: ["latin"], variable: "--font-roboto" });

export default function MyApp({ Component, pageProps }: AppProps) {
  const path = useRouter();
  return (
    <div className={`h-full flex ${ibm_plex_sans_thai_loop.className}`}>
      <SideBar pathname={path.pathname} />
      <div className="flex flex-col w-full">
        <NavBar />
        <Component {...pageProps} />
      </div>
    </div>
  );
}
