"use client";

import Header from "@/components/Header";
import { useState, useEffect } from "react";
import { post } from "@/utils/fething";
import Question from "./Question";
import styles from "@/app/newsurvey/page.module.css";
import BlueButton from "@/components/BlueButton";

export default function MySurvey({ params }) {
    const [questions, setQuestions] = useState([]);
    const [passes, setPasses] = useState({});

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
            setQuestions(response);
        }

        fetchQuestions();
    }, []);

    useEffect(() => {
        async function fetchPasses() {
            const response = await post(
                process.env.NEXT_PUBLIC_SURVEYS +
                    "/" +
                    params.id +
                    process.env.NEXT_PUBLIC_PASSES,
                {
                    token: localStorage.getItem("token"),
                }
            );
            setPasses(response);
        }

        fetchPasses();
    }, []);

    let questionsArray = questions.questions;
    questionsArray ??= [];
    questionsArray.sort((a, b) => a.Number - b.Number);

    let passesRef =
        "data:text/json;charset=utf-8," +
        encodeURIComponent(JSON.stringify(passes));

    return (
        <>
            <Header />

            <div className={styles.container_questions}>
                <p className={styles.topic}>{questions.topic}</p>
                {questionsArray.map((question) => (
                    <Question
                        key={question.Number}
                        topic={question.QuestionText}
                        number={question.Number}
                    />
                ))}
                <div className={styles.container_button}>
                    <div className={styles.button}>
                        <a href={passesRef} download="passes.json">
                            <BlueButton
                                text={"Загрузить данные о прохождениях"}
                            />
                        </a>
                    </div>
                </div>
            </div>
        </>
    );
}
