import styles from "./register-form.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";

const LoginForm = () => {
    return (
        <div className={styles.form_reg_block}>
            <p className={styles.form_reg_block_head_text}>Регистрация в Survey Sphere</p>
            <form action="">
                <label className={styles.label}>Почта</label>
                <Input type="email" placeholder="m@example.com" />
                <label className={styles.label}>Пароль</label>
                <Input type="password" />
                <label className={styles.label}>Повторите пароль</label>
                <Input type="password" />
                <BlueButton text="Зарегистрироваться"></BlueButton>
            </form>
            <p className={styles.authorization_text}>
                Уже есть аккаунт? <Link className={styles.link} href="/login">Войти</Link>
            </p>
        </div>
    );
};

export default LoginForm;