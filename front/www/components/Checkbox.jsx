import styles from "./styles/checkbox.module.css";

const Checkbox = ({ text }) => {
    return (
        <>
            <input type="checkbox" id={text} className={styles.custom_checkbox} />
            <label htmlFor={text}>{text}</label>
        </>
    );
};

export default Checkbox;