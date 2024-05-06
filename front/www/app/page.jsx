import Link from "next/link";

export default function Home(children) {
    return (
        <main>
            <p>
                <Link href={"/login"}>Войти</Link>
            </p>
            <p>
                <Link href={"/register"}>Регистрация</Link>
            </p>
            <p>
                <Link href={"/mysurveys"}>Мои опросы</Link>
            </p>
        </main>
    );
}
