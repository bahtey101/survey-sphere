import styles from "./styles/header.module.css";
import Link from "next/link";

const Header = () => {
    return (
        <header className={styles.header}>
            <nav className={styles.header_container}>
                <Link className={styles.site_name} href={"/"}>
                    Survey Sphere
                </Link>



                <Link className={styles.exit_button} href={"/"}>
                    Выйти
                </Link>
            </nav>
        </header>
    );
};

export default Header;
