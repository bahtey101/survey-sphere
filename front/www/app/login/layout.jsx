import { Inter } from "next/font/google";
import "@/app/globals.css";

const inter = Inter({ subsets: ["latin", "cyrillic"] });

export default function RootLayout({ children }) {
  return (
    <div className={inter.className}>{children}</div>
  );
}