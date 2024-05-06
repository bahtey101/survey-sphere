import Link from "next/link";
import LoginForm from "./LoginForm";

export default function Login(children) {
    return (
        <>
            <Link href={"/"}>Домой</Link>
            <div className="fullscreen_container">
                <LoginForm />
            </div>
        </>
    );
}
