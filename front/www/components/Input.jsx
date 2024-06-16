import styles from "./styles/input.module.css";
import { Inter } from "next/font/google";

const inter = Inter({ subsets: ["latin", "cyrillic"] });

const Input = ({ type, placeholder, id, name, onChange }) => {
    return (
        <input
            type={type}
            placeholder={placeholder}
            name={name}
            className={styles.input}
            id={id}
            onChange={onChange}
        ></input>
    );
};

export default Input;
