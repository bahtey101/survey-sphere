import Header from "@/components/Header";
import styles from "./page.module.css";
import Image from "next/image";
import Input from "@/components/Input";
import noImage from "@/public/no_image.jpg";
import BlueButton from "@/components/BlueButton";

export default function Newsurvey(children) {
    var image_src = image_src || noImage;
    return (
        <>
            <Header />
            <div className={styles.container}>
                <div className={styles.image}>
                    <Image
                        alt="Картинка не загрузилась :("
                        src={image_src}
                        layout="fill"
                        objectFit="cover"
                    />
                </div>
                <p className={styles.topic}>Название</p>
                <form className={styles.container_questions}>
                    <div className={styles.question}>
                        <label className={styles.question_label}>
                            Вопрос 1
                        </label>
                        <Input />
                    </div>
                    <div className={styles.question}>
                        <label className={styles.question_label}>
                            Вопрос 2
                        </label>

                        <div className={styles.container_checkbox}>
                            <input type="checkbox" id="2_1" />
                            <label htmlFor="2_1">Вариант 1</label>
                        </div>
                        <div className={styles.container_checkbox}>
                            <input type="checkbox" id="2_2" />
                            <label htmlFor="2_2"> Вариант 2</label>
                        </div>
                        <div className={styles.container_checkbox}>
                            <input type="checkbox" id="2_3" />
                            <label htmlFor="2_3"> Вариант 3</label>
                        </div>
                    </div>
                    <div className={styles.container_button}>
                        <div className={styles.button}>
                            <BlueButton text={"Отправить"} />
                        </div>
                    </div>
                </form>
            </div>
        </>
    );
}
