"use client";

import Header from "@/components/Header";
import styles from "./page.module.css";
import Image from "next/image";
import Input from "@/components/Input";
import noImage from "@/public/no_image.jpg";
import BlueButton from "@/components/BlueButton";
import { useState } from "react";
import QuestionContainer from "./QuestionContainer";
import { post } from "@/utils/fething";
import { useRouter } from "next/navigation";

export default function NewSurvey(children) {
    var image_src = image_src || noImage;

    const router = useRouter();

    const [questions, setQuestions] = useState([""]);

    const [topic, setTopic] = useState("");

    const handleTopicChange = (event) => {
        setTopic(event.target.value);
    };

    async function handleSubmit(event) {
        event.preventDefault();

        let request = {
            questions: questions,
            token: localStorage.getItem("token"),
            topic: topic,
        };

        let response = await post(process.env.NEXT_PUBLIC_NEW_SURVEY, request);

        router.push("/mysurveys");
    }

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

                <form
                    className={styles.container_questions}
                    onSubmit={handleSubmit}
                >
                    <div className={styles.topic_container}>
                        <span className={styles.topic}>Название опроса:</span>
                        <div className={styles.topic_input}>
                            <Input
                                isRequired={true}
                                onChange={handleTopicChange}
                            />
                        </div>
                    </div>

                    <QuestionContainer
                        questions={questions}
                        setQuestions={setQuestions}
                    />

                    <div className={styles.container_button}>
                        <div className={styles.button}>
                            <BlueButton text={"Создать опрос"} />
                        </div>
                    </div>
                </form>
            </div>
        </>
    );
}
