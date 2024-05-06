import styles from "@/styles/my_surveys/survey-with-image.module.css";
import BlueButton from "@/components/BlueButton";

const SurveyWithImage = ({ survey_name, image_src }) => {
    image_src = image_src || "/";
    return (
        <div className={styles.container}>
            <div className={styles.image}>
                <img src={image_src} alt="Картинка не загрузилась :(" />
            </div>
            <div className={styles.place_for_name_and_button}>
                <div className={styles.survey_name}>{survey_name}</div>
                <div className={styles.button}>
                    <BlueButton text="Просмотреть" />
                </div>
            </div>
        </div>
    );
};

export default SurveyWithImage;
