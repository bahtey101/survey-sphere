import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";

const LoginForm = () => {
    return (
        <div className={styles.form}>
            <p className={`${styles.form_head_text} blue_text`}>
                Вход в Survey Sphere
            </p>
            <form action="">
                <label className={styles.form_label}>Почта</label>
                <Input type="email" placeholder="m@example.com" />
                <label className={styles.form_label}>Пароль</label>
                <Input type="password" />
                <BlueButton text="Войти"></BlueButton>
            </form>
            <p className={styles.form_small_centered_text}>
                Впервые здесь?{" "}
                <Link
                    className={styles.form_small_centered_text_link}
                    href="/register"
                >
                    Зарегистрироваться
                </Link>
            </p>
        </div>
    );
};

export default LoginForm;
