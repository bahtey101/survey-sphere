import styles from "@/styles/my_surveys/survey-with-image.module.css";
import Image from "next/image";
import BlueButton from "@/components/BlueButton";

const SurveyWithImage = () => {
    return (
        <div className={styles.container}>
            {/* <Image
                // width={300}
                // height={500}
                fill
                src="/public/Triangles-4k.png"
                alt="Картинка не загрузилась :("
            /> */}
            <div className={styles.image}>
                <img
                    src="/public/Triangles-4k.png"
                    alt="Картинка не загрузилась :("
                />
            </div>
            <div className={styles.place_for_name_and_button}>
                <div className={styles.survey_name}>SurveyWithImage</div>
                <div>
                    <BlueButton text="Просмотреть" />
                </div>
            </div>
        </div>
    );
};

export default SurveyWithImage;
