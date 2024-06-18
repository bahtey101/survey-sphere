"use client";

import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";
import { post } from "@/utils/fething";
import { useState } from "react";
import { useRouter } from "next/navigation";

const RegisterForm = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");
    const [emailError, setEmailError] = useState("");
    const [passwordError, setPasswordError] = useState("");
    const [confirmPasswordError, setConfirmPasswordError] = useState("");

    const router = useRouter();

    const handleEmailChange = (event) => {
        setEmail(event.target.value);
    };

    const handlePasswordChange = (event) => {
        setPassword(event.target.value);
    };

    const handleConfirmPasswordChange = (event) => {
        setConfirmPassword(event.target.value);
    };

    async function handleSubmit(event) {
        event.preventDefault();

        setEmailError("");
        setPasswordError("");
        setConfirmPasswordError("");

        let formIsValid = true;

        if (password != confirmPassword) {
            setConfirmPasswordError("Пароли не совпадают");
            formIsValid = false;
        }

        if (!email.trim()) {
            setEmailError("Введите почту");
            formIsValid = false;
        }

        if (!password.trim()) {
            setPasswordError("Введите пароль");
            formIsValid = false;
        }

        if (!confirmPassword.trim()) {
            setConfirmPasswordError("Введите пароль повторно");
            formIsValid = false;
        }

        if (formIsValid) {
            let response = await post(process.env.NEXT_PUBLIC_SIGN_UP, {
                login: email,
                password: password,
            });

            router.push("/login");
        }
    }

    return (
        <div className={styles.form}>
            <p className={`${styles.form_head_text} blue_text`}>
                Регистрация в Survey Sphere
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
                <label className={styles.form_label}>Повторите пароль</label>
                <p className={styles.error_text}>{confirmPasswordError}</p>
                <Input
                    type="password"
                    name="confirm-password"
                    onChange={handleConfirmPasswordChange}
                />
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

export default RegisterForm;
