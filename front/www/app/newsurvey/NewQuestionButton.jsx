import styles from "./page.module.css";
import plusIcon from "@/public/plus.svg";
import Image from "next/image";

const NewQuestionButton = ({ onClick }) => {
    return (
        <button
            className={styles.grey_circle}
            title="Добавить вопрос"
            onClick={onClick}
            type="button"
        >
            <Image
                src={plusIcon}
                height={30}
                width={30}
                alt="Что-то пошло не так"
            />
        </button>
    );
};

export default NewQuestionButton;
