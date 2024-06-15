import Link from "next/link";

async function getData() {
    const res = await fetch("http://localhost:8080/users");
    // The return value is *not* serialized
    // You can return Date, Map, Set, etc.

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error("Failed to fetch data");
    }

    return res.json();
}

export default async function Home(children) {
    const data = await getData();

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
            <p>User ID is: {data.data.ID}</p>
        </main>
    );
}
