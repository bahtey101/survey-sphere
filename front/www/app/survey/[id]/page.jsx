"use client";

import Header from "@/components/Header";
import styles from "@/app/newsurvey/page.module.css";
import Image from "next/image";
import Input from "@/components/Input";
import noImage from "@/public/no_image.jpg";
import BlueButton from "@/components/BlueButton";
import { useState, useEffect } from "react";
import { post } from "@/utils/fething";
import { useRouter } from "next/navigation";
import Question from "./Question";

export default function Survey({ params }) {
    var image_src = image_src || noImage;

    const router = useRouter();

    const [survey, setSurvey] = useState([]);

    useEffect(() => {
        async function fetchQuestions() {
            const response = await post(
                process.env.NEXT_PUBLIC_SURVEYS +
                    "/" +
                    params.id +
                    process.env.NEXT_PUBLIC_QUESTIONS,
                {
                    token: localStorage.getItem("token"),
                }
            );
            setSurvey(response);
        }

        fetchQuestions();
    }, []);

    let topic = survey.topic;
    let questions = survey.questions;
    topic ??= "";
    questions ??= [];
    const [answers, setAnswers] = useState([Array(questions.length).fill("")]);

    const onChange = (newInput, index) => {
        let newAnswers = [...answers];
        newAnswers[index] = newInput;
        setAnswers(newAnswers);
    };

    async function handleSubmit(event) {
        event.preventDefault();

        let request = {
            answers: answers,
            token: localStorage.getItem("token"),
            survey_id: Number(params.id),
        };

        let response = await post(process.env.NEXT_PUBLIC_NEW_PASS, request);

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
                        <span className={styles.topic}>{topic}</span>
                        <div className={styles.topic_input}></div>
                    </div>

                    {questions.map((question, index) => (
                        <Question
                            key={index}
                            questionText={question.QuestionText}
                            number={index}
                            onChange={onChange}
                        />
                    ))}

                    <div className={styles.container_button}>
                        <div className={styles.button}>
                            <BlueButton text={"Сохранить ответы"} />
                        </div>
                    </div>
                </form>
            </div>
        </>
    );
}
