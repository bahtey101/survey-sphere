import Link from "next/link";

export default async function Home(children) {
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
            <p>
                <Link href={"/newsurvey"}>Новый опрос</Link>
            </p>
            <p>
                <Link href={"/adminpanel"}>Панель администратора</Link>
            </p>
        </main>
    );
}
