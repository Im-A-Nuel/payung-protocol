import type { Metadata, Viewport } from "next";
import { Plus_Jakarta_Sans } from "next/font/google";

import { DaftarSW } from "@/components/daftar-sw";
import { Nav } from "@/components/nav";
import { Providers } from "./providers";
import "./globals.css";

const jakarta = Plus_Jakarta_Sans({
  variable: "--font-jakarta",
  subsets: ["latin"],
  display: "swap",
});

export const metadata: Metadata = {
  title: "Payung",
  description:
    "Asuransi hujan untuk driver ojol. Hujan lewat 20 mm di zonamu, Rp 25.000 masuk ke dompet. Tanpa klaim, tanpa foto.",
  manifest: "/manifest.webmanifest",
  appleWebApp: { capable: true, statusBarStyle: "default", title: "Payung" },
  icons: { icon: "/icon.svg", apple: "/apple-touch-icon.png" },
};

export const viewport: Viewport = {
  themeColor: "#0b5fd0",
  width: "device-width",
  initialScale: 1,
  viewportFit: "cover",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="id">
      <body className={`${jakarta.variable} antialiased`}>
        <Providers>
          <div className="mx-auto flex min-h-dvh w-full max-w-[520px] flex-col">
            <main className="flex-1 px-4 pt-5 pb-28">{children}</main>
            <Nav />
          </div>
          <DaftarSW />
        </Providers>
      </body>
    </html>
  );
}
