import styles from "./LoginForm.module.css";
import BlueButton from "@/components/BlueButton";
import Input from "@/components/Input";
import Link from "next/link";

const LoginForm = () => {
  return (
    <div className={styles.form_auth_block}>
      <p className={styles.form_auth_block_head_text}>Вход в Survey Sphere</p>
      {/* <p>Введите свою электронную почту и пароль для входа в аккаунт</p> */}
      <form action="">
        <label className={styles.label}>Почта</label>
        <Input type="email" placeholder="m@example.com" />
        <label className={styles.label}>Пароль</label>
        <Input type="password" />
        <BlueButton text="Войти"></BlueButton>
      </form>
      <p className={styles.registration_text}>Впервые здесь? <Link href="/register">Зарегистрироваться</Link></p>
    </div>
  );
};

export default LoginForm;
