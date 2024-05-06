import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";

const LoginForm = () => {
    return (
        <div className={styles.form}>
            <p className={`${styles.form_head_text} blue_text`}>
                Регистрация в Survey Sphere
            </p>
            <form action="">
                <label className={styles.form_label}>Почта</label>
                <Input type="email" placeholder="m@example.com" />
                <label className={styles.form_label}>Пароль</label>
                <Input type="password" />
                <label className={styles.form_label}>Повторите пароль</label>
                <Input type="password" />
                <div style={{ marginTop: 15 }}>
                    <BlueButton text="Зарегистрироваться"></BlueButton>
                </div>
            </form>
            <p className={styles.form_small_centered_text}>
                Уже есть аккаунт?{" "}
                <Link
                    className={styles.form_small_centered_text_link}
                    href="/login"
                >
                    Войти
                </Link>
            </p>
        </div>
    );
};

export default LoginForm;
