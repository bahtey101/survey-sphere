import styles from "@/styles/my_surveys/survey.module.css";
import BlueButton from "@/components/BlueButton";
import Image from "next/image";
import noImage from "@/public/no_image.jpg";

const Survey = ({ topic, surveyID, image_src }) => {
    image_src = image_src || noImage;
    let linkToSurvey = "/mysurveys/" + surveyID;

    return (
        <div className={styles.container}>
            <div className={styles.image}>
                <Image
                    alt="Картинка не загрузилась :("
                    src={image_src}
                    fill={true}
                />
            </div>

            <div className={styles.info_container}>
                <div className={styles.topic}>{topic}</div>

                <div style={{ padding: `10px` }}>
                    <form action={linkToSurvey} method="get">
                        <BlueButton text={"Просмотреть"} />
                    </form>
                </div>
            </div>
        </div>
    );
};

export default Survey;
