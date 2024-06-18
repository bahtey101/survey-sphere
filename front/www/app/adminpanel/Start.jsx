import styles from "./page.module.css";

const Start = () => {
    return (
        <>
            <div className={styles.panel}>
                <div className={styles.click_button_container}>
                    <span className={styles.click_button_text}>
                        Выберите категорию отображаемого контента (Пользователи
                        или Опросы)
                    </span>
                </div>
            </div>
        </>
    );
};

export default Start;
