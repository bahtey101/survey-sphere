import styles from "@/styles/my_surveys/survey-with-image.module.css";
import Image from "next/image";
import BlueButton from "@/components/BlueButton";

const SurveyWithImage = () => {
    return (
        <div className={styles.container}>
            <div className={styles.image}>
                <img
                    src="/public/Triangles-4k.png"
                    alt="Картинка не загрузилась :("
                />
            </div>
            <div className={styles.place_for_name_and_button}>
                <div className={styles.survey_name}>
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                    Facilis unde ab, in, quia, enim minima nihil alias
                    exercitationem animi temporibus quae cupiditate commodi
                    saepe rerum. Sint ipsam quo facilis fugiat!
                </div>
                <div className={styles.button}>
                    <BlueButton text="Просмотреть" />
                </div>
            </div>
        </div>
    );
};

export default SurveyWithImage;
