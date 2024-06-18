"use client";

import styles from "./page.module.css";
import Survey from "./Survey.jsx";
import Header from "@/components/Header";
import NewSurveyButton from "./NewSurveyButton";
import { post } from "@/utils/fething";
import { useEffect, useState } from "react";

export default function Mysurveys(children) {
    const [surveys, setSurveys] = useState([]);
    let surveysArray = surveys.surveys;
    surveysArray ??= [];

    useEffect(() => {
        async function fetchSurveys() {
            const response = await post(process.env.NEXT_PUBLIC_SURVEYS, {
                token: localStorage.getItem("token"),
            });
            setSurveys(response);
        }

        fetchSurveys();
    }, []);

    return (
        <>
            <Header />

            <div className={styles.grid} id="grid">
                {surveysArray.map((survey) => (
                    <Survey
                        key={survey.ID}
                        surveyID={survey.ID}
                        topic={survey.Topic}
                    />
                ))}
                <NewSurveyButton />
            </div>
        </>
    );
}
