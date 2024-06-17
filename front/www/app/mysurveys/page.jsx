"use client";

import styles from "./page.module.css";
import Survey from "./Survey.jsx";
import Header from "@/components/Header";
import NewSurveyButton from "./NewSurveyButton";
import { post } from "@/utils/fething";
import { useEffect, useState } from "react";

export default function Mysurveys(children) {
    const [surveys, setSurveys] = useState({ surveys: [{ Topic: "" }] });

    useEffect(() => {
        async function fetchSurveys() {
            const response = await post(process.env.NEXT_PUBLIC_SURVEYS, {
                token: localStorage.getItem("token"),
                topic: "asd",
            });
            setSurveys(response);
        }

        fetchSurveys();
    }, []);

    return (
        <>
            <Header />

            <div className={styles.grid} id="grid">
                {surveys.surveys.map((survey) => (
                    <Survey topic={survey.Topic} />
                ))}
                <NewSurveyButton />
            </div>
        </>
    );
}
