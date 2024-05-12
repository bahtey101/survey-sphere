import Header from "@/components/Header";
import styles from "./page.module.css";

export default function AdminPanel(children) {
    return (
        <>
            <Header />
            <div className={styles.container}>
                <div className={styles.panel}>
                    <div className={styles.panel_flex}>
                        <div className={styles.radio_button}>
                            <input
                                name="tableType"
                                type="radio"
                                id="users"
                                className={styles.radio_field}
                            />
                            <label
                                htmlFor="users"
                                className={styles.radio_label}
                            >
                                Пользователи
                            </label>
                        </div>
                        <div className={styles.radio_button}>
                            <input
                                name="tableType"
                                type="radio"
                                id="surveys"
                                className={styles.radio_field}
                            />
                            <label
                                htmlFor="surveys"
                                className={styles.radio_label}
                            >
                                Опросы
                            </label>
                        </div>
                        <textarea
                            className={styles.flex_horizontal_split}
                        ></textarea>
                    </div>
                </div>
            </div>
        </>
    );
}
