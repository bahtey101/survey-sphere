import styles from "./page.module.css";
import RegisterForm from "./RegisterForm"

export default function Login(children) {
    return (
        <div className={styles.box}>
            <RegisterForm />
        </div>
    )
}