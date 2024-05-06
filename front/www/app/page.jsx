import Link from "next/link";

export default function Home(children) {
    return (
        <main>
            <Link href={"/login"}>Войти</Link>
        </main>
    );
}
