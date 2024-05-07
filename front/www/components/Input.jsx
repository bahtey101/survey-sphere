import styles from "./styles/input.module.css";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin", "cyrillic"] });

const Input = ({ type, placeholder, id }) => {
    return (
        <input
            type={type}
            placeholder={placeholder}
            className={styles.input}
            id={id}
        ></input>
    );
};

export default Input;
