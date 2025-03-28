import type { AppProps } from "next/app";
import "../app/globals.css";
import { IBM_Plex_Sans_Thai_Looped, Roboto_Flex } from "next/font/google";
import type { ReactElement, ReactNode } from "react";
import type { NextPage } from "next";

const ibm_plex_sans_thai_loop = IBM_Plex_Sans_Thai_Looped({
  weight: ["100", "200", "300", "400", "500", "600", "700"],
  subsets: ["latin", "thai"],
  variable: "--font-ibm-plex-sans-thai-loop",
});

// const roboto = Roboto_Flex({ subsets: ["latin"], variable: "--font-roboto" });

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

export default function MyApp({ Component, pageProps }: AppPropsWithLayout) {
  const getLayout = Component.getLayout ?? ((page) => page);

  return getLayout(
    <div className={`h-full flex ${ibm_plex_sans_thai_loop.className}`}>
      <Component {...pageProps} />
    </div>
  );
}
