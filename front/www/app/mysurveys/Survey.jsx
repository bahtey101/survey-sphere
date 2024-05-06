import styles from "@/styles/my_surveys/survey.module.css"
import BlueButton from "@/components/BlueButton";

const Survey = ({ topic, image_src }) => {
    image_src = image_src || "https://cdn11.bigcommerce.com/s-ucl2nc/images/stencil/1280x1280/c/placeholder-image__65448.original.jpg";
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