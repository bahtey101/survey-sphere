"use client";

import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";
import { useState } from "react";
import { post } from "@/utils/fething";

const LoginForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [emailError, setEmailError] = useState("");
    const [passwordError, setPasswordError] = useState("");

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    async function handleSubmit(event) {
        event.preventDefault();

        let formIsValid = true;

        if (!email.trim()) {
            setEmailError("Введите почту");
            formIsValid = false;
        }

        if (!password.trim()) {
            setPasswordError("Введите пароль");
            formIsValid = false;
        }

        if (formIsValid) {
            let response = await post(process.env.NEXT_PUBLIC_SIGN_IN, {
                login: email,
                password: password,
            });
        }
    }

    return (
        <div className={styles.form}>
            <p className={`${styles.form_head_text} blue_text`}>
                Вход в Survey Sphere
            </p>
            <form action="" onSubmit={handleSubmit}>
                <label className={styles.form_label}>Почта</label>
                <p className={styles.error_text}>{emailError}</p>
                <Input
                    type="email"
                    name="email"
                    placeholder="m@example.com"
                    onChange={handleEmailChange}
                />
                <label className={styles.form_label}>Пароль</label>
                <p className={styles.error_text}>{passwordError}</p>
                <Input
                    type="password"
                    name="password"
                    onChange={handlePasswordChange}
                />
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
