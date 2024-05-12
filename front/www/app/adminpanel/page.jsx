import Header from "@/components/Header";
import styles from "./page.module.css";
import Script from "next/script";

export default function AdminPanel(children) {
    return (
        <>
            <Header />
            <Script src={"/static/scripts/adminpanel/script.js"} />
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
                        <input
                            type="search"
                            className={`${styles.flex_horizontal_split} ${styles.search_bar}`}
                        ></input>
                    </div>
                </div>
                <div className={styles.panel}>
                    <div className={styles.click_button_container}>
                        <span className={styles.click_button_text}>
                            Выберите категорию отображаемого контента
                            (Пользователи или Опросы)
                        </span>
                    </div>
                </div>
                <div className={styles.table_panel}>
                    <table className={styles.table}>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Логин</th>
                                <th>Роль</th>
                                <th>Кол-во опросов</th>
                                <th>КНОПКА</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>asdfadf</td>
                                <td>asdfasdfasdf</td>
                                <td>123123123123</td>
                                <td>qrweqerwqersdfsdf</td>
                                <td>ad9ujwgi4w</td>
                            </tr>
                            <tr>
                                <td>1</td>
                                <td>2</td>
                                <td>3</td>
                                <td>4</td>
                                <td>5</td>
                            </tr>
                            <tr>
                                <td>1</td>
                                <td>2</td>
                                <td>3</td>
                                <td>4</td>
                                <td>5</td>
                            </tr>
                            <tr>
                                <td>1</td>
                                <td>2</td>
                                <td>3</td>
                                <td>4</td>
                                <td>5</td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </>
    );
}
