import styles from "./styles/blue-button.module.css";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin", "cyrillic"] });

const BlueButton = ({text}) => {
  return (
    <button className={`${styles.blue_button} ${inter.className}`}>{text}</button>
  )
}

export default BlueButton