import Header from "@/components/Header";
import styles from "./page.module.css";
import BlueButton from "@/components/BlueButton";

export default function Stats(children) {
    return (
        <>
            <Header />
            <div className={styles.container}>

                <div className={styles.panel}>
                    <div className={styles.large_text}>Название</div>
                    <div style={{ width: `150px` }}>
                        <BlueButton text={"Удалить опрос"} />
                    </div>
                </div>

                <div className={styles.grid}>

                    <div className={styles.cell}>
                        <div className={styles.text}>Опрошено</div>
                        <div className={styles.large_text}>1100</div>
                        <div className={styles.small_text}>пользователей</div>
                    </div>

                    <div className={styles.cell}>
                        <div className={styles.text}>Ответили на все вопросы</div>
                        <div className={styles.large_text}>89%</div>
                        <div className={styles.small_text}>пользователей</div>
                    </div>

                    <div className={styles.cell}>
                        <div className={styles.text}>Средняя оценка</div>
                        <div className={styles.large_text}>4.7</div>
                        <div className={styles.small_text}></div>
                    </div>

                </div>

                <p>Вопросы</p>

                <table className={styles.table}>
                    <thead>
                        <tr>
                            <th>Номер вопроса</th>
                            <th>Текст</th>
                            <th>Количество ответов</th>
                            <th>.json</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>1</td>
                            <td>Вопрос 1</td>
                            <td>100</td>
                            <td>
                                <div className={styles.button_container}>
                                    <div className={styles.button_download} />
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>2</td>
                            <td>Вопрос 2</td>
                            <td>80</td>
                            <td>
                                <div className={styles.button_container}>
                                    <div className={styles.button_download} />
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>3</td>
                            <td>Вопрос 3</td>
                            <td>120</td>
                            <td>
                                <div className={styles.button_container}>
                                    <div className={styles.button_download} />
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>4</td>
                            <td>Вопрос 4</td>
                            <td>90</td>
                            <td>
                                <div className={styles.button_container}>
                                    <div className={styles.button_download} />
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td>5</td>
                            <td>Вопрос 5</td>
                            <td>30</td>
                            <td>
                                <div className={styles.button_container}>
                                    <div className={styles.button_download} />
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div >
        </>
    );
}