"use client";

import styles from "@/styles/auth_forms/main.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";
import { post } from "@/utils/fething";

const RegisterForm = () => {
    async function handleSubmit(event) {
        event.preventDefault();
        const formData = new FormData(event.target);
        for (var [_, value] of formData.entries()) {
            if (value == "") {
                return;
            }
        }
        let not_same_passwords = document.getElementById("not-same");
        if (formData.get("password") == formData.get("password-again")) {
            if (not_same_passwords != null) {
                not_same_passwords.remove();
            }
            let response = post(process.env.NEXT_PUBLIC_SIGN_UP, {
                login: formData.get("email"),
                password: formData.get("password"),
            });
        } else {
            if (not_same_passwords == null) {
                let p = document.createElement("p");
                p.className = styles.error_text;
                p.id = "not-same";
                p.innerText = "Пароли не совпадают";
                let password_again = document.getElementById("password-again");
                password_again.before(p);
            }
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
                <Input
                    type="password"
                    name="password-again"
                    id="password-again"
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
