import styles from "@/styles/my_surveys/survey.module.css";
import plusIcon from "@/public/plus.svg";
import Image from "next/image";
import Link from "next/link";

const NewSurveyButton = () => {
    return (
        <div className={styles.container_new_survey_button}>
            <Link className={styles.grey_circle} href="/newsurvey">
                <Image
                    src={plusIcon}
                    height={50}
                    width={50}
                    alt="Что-то пошло не так"
                />
            </Link>
        </div>
    );
};

export default NewSurveyButton;
