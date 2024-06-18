"use client";

import styles from "./styles/header.module.css";
import Link from "next/link";

const Header = () => {
    const LogOut = () => {
        localStorage.removeItem("token");
    };
    return (
        <header className={styles.header}>
            <nav className={styles.header_container}>
                <Link className={styles.site_name} href={"/"}>
                    Survey Sphere
                </Link>

                <Link
                    href={"/"}
                    className={styles.exit_button}
                    onClick={LogOut}
                >
                    Выйти
                </Link>
            </nav>
        </header>
    );
};

export default Header;
