import styles from "@/styles/my_surveys/survey.module.css"
import BlueButton from "@/components/BlueButton";

const Survey = ({ topic, image_src }) => {
    image_src = image_src || "https://www.al2uno.com/wp-content/uploads/2023/02/no-image-1536x1536.jpg";
    return (
        <div className={styles.container}>

            <div className={styles.image}>
                <img src={image_src} />
            </div>

            <div className={styles.info_container}>
                <div className={styles.topic}>{topic}</div>

                <div style={{ padding: `10px` }}>
                    <BlueButton text={"Просмотреть"} />
                </div>
            </div>
        </div>
    )
}

export default Survey;