"use client";

import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";

const LoginForm = () => {
    async function handleSubmit(event) {
        event.preventDefault();
        const formData = new FormData(event.target);

        let response = await fetch("http://localhost:8080/auth/sign-in", {
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
        alert(result.token);
    }

    return (
        <div className={styles.form}>
            <p className={`${styles.form_head_text} blue_text`}>
                Вход в Survey Sphere
            </p>
            <form action="" onSubmit={handleSubmit}>
                <label className={styles.form_label}>Почта</label>
                <Input type="email" name="email" placeholder="m@example.com" />
                <label className={styles.form_label}>Пароль</label>
                <Input type="password" name="password" />
                <div style={{ marginTop: 15 }}>
                    <BlueButton text="Войти"></BlueButton>
                </div>
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
