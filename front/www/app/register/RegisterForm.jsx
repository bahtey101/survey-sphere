"use client";

import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";

const LoginForm = () => {
    async function handleSubmit(event) {
        event.preventDefault();
        const formData = new FormData(event.target);
        if (formData.get("password") == formData.get("password-again")) {
            let response = await fetch("http://localhost:8080/users", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json;charset=utf-8",
                },
                body: JSON.stringify({
                    login: formData.get("email"),
                    password: formData.get("password"),
                }),
            });

            let result = await response.json();
            console.log(result);
            alert(result.id);
        }
    }

    return (
        <div className={styles.form}>
            <p className={`${styles.form_head_text} blue_text`}>
                Регистрация в Survey Sphere
            </p>
            <form action="" onSubmit={handleSubmit}>
                <label className={styles.form_label}>Почта</label>
                <Input type="email" name="email" placeholder="m@example.com" />
                <label className={styles.form_label}>Пароль</label>
                <Input type="password" name="password" />
                <label className={styles.form_label}>Повторите пароль</label>
                <Input type="password" name="password-again" />
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
