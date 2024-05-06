import Link from "next/link";
import LoginForm from "./LoginForm";

export default function Login(children) {
    return (
        <>
            <div className="fullscreen_container">
                <LoginForm />
            </div>
        </>
    );
}
