import styles from "./styles/input.module.css"
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin", "cyrillic"] });

const Input = ({type, placeholder}) => {
  return (
    <input type={type} placeholder={placeholder} className={styles.input}></input>
  )
}

export default Input