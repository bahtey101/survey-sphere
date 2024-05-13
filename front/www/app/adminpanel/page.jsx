import Header from "@/components/Header";
import styles from "./page.module.css";
import Script from "next/script";
import BlueButton from "@/components/BlueButton";

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
                <div className={styles.panel}>
                    <table className={styles.table}>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Логин</th>
                                <th>Роль</th>
                                <th>Кол-во опросов</th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>asdfadf</td>
                                <td>asdfasdfasdf</td>
                                <td>123123123123</td>
                                <td>qrweqerwqersdfsdf</td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr>
                                <td>1</td>
                                <td>2</td>
                                <td>3</td>
                                <td>4</td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr>
                                <td>1</td>
                                <td>2</td>
                                <td>3</td>
                                <td>4</td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr>
                                <td>
                                    Lorem ipsum dolor sit, amet consectetur
                                    adipisicing elit. Corrupti esse, itaque iste
                                    sit beatae doloribus perferendis et, dicta
                                    mollitia laboriosam atque voluptates
                                    consequatur, maxime dolore corporis id.
                                    Pariatur, aut tempore!
                                </td>
                                <td>
                                    Lorem ipsum dolor sit amet consectetur
                                    adipisicing elit. Explicabo mollitia alias
                                    consectetur laboriosam culpa reiciendis,
                                    officia iusto ex vitae porro quod molestias
                                    error id a nihil amet tenetur cumque illum?
                                </td>
                                <td>
                                    Lorem ipsum dolor sit amet consectetur
                                    adipisicing elit. Amet, assumenda? Autem,
                                    eveniet corporis. Repudiandae esse cumque
                                    eaque? Natus cumque accusantium soluta
                                    aperiam. Sapiente cum, repellendus ullam
                                    deserunt eveniet nemo aut.
                                </td>
                                <td>
                                    {" "}
                                    Lorem ipsum dolor, sit amet consectetur
                                    adipisicing elit. Unde corrupti excepturi
                                    perspiciatis eos sint cumque possimus
                                    temporibus dolore, earum ab nesciunt totam
                                    repellat, aperiam, sit dolor! Blanditiis
                                    dignissimos magnam debitis.
                                </td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr>
                                <td>1</td>
                                <td>2</td>
                                <td>3</td>
                                <td>4</td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div className={styles.panel}>
                    <table className={styles.table}>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Название</th>
                                <th>Логин автора</th>
                                <th>Дата создания</th>
                                <th></th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>1</td>
                                <td>Какая ты собака?</td>
                                <td>dima@mail.ru</td>
                                <td>01.01.2024</td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                            <tr>
                                <td>2</td>
                                <td>Какая ты кошка?</td>
                                <td>nekit@mail.ru</td>
                                <td>02.02.2024</td>
                                <td>
                                    <div
                                        className={
                                            styles.container_button_delete
                                        }
                                    >
                                        <button
                                            className={styles.button_delete}
                                        >
                                            Удалить
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <div className={styles.container_button_load_more_flex}>
                        <div className={styles.container_button_load_more}>
                            <BlueButton text="Загрузить ещё" />
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
