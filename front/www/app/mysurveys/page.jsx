import Header from "@/components/Header";
import SurveyWithImage from "./SurveyWithImage";
import styles from "./page.module.css";

export default function Mysurveys(children) {
    return (
        <>
            <Header />
            <div style={{ margin: "100px 140px" }}>
                <p className={styles.header}>Мои опросы</p>

                <div className={styles.container}>
                    <SurveyWithImage
                        image_src={
                            "https://cdn.britannica.com/45/5645-050-B9EC0205/head-treasure-flower-disk-flowers-inflorescence-ray.jpg"
                        }
                        survey_name={"Survey 1"}
                    />
                    <SurveyWithImage
                        survey_name={
                            "Survey 2 hello my dear madness ahahahahhah"
                        }
                    />
                    <SurveyWithImage
                        survey_name={
                            "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Dolores veniam temporibus saepe sequi magnam facilis quos minus cum! Laboriosam quam nulla perspiciatis ducimus sunt amet deleniti sequi, excepturi animi error?"
                        }
                    />
                    <SurveyWithImage />
                    <SurveyWithImage />
                    <SurveyWithImage />
                    <SurveyWithImage />
                    <SurveyWithImage />
                    <SurveyWithImage />
                </div>
            </div>
        </>
    );
}
