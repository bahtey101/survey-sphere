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

                <div className={styles.questions}>
                    <div className={styles.row}>
                        1
                    </div>
                    <div className={styles.row}>
                        2
                    </div>
                    <div className={styles.row}>
                        3
                    </div>
                </div>
            </div>
        </>
    );
}