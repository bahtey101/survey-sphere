import styles from "@/app/newsurvey/page.module.css";

const Question = ({ topic, number }) => {
    return (
        <>
            <div className={styles.question}>
                <label className={styles.question_label}>Вопрос {number}</label>
                <p>{topic}</p>
            </div>
        </>
    );
};

export default Question;
