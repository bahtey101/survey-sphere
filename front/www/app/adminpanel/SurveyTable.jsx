import styles from "./page.module.css";
import { useEffect, useState } from "react";
import { post } from "@/utils/fething";
import { deleteRequest } from "@/utils/fething";

import SurveyTableRow from "./SurveyTableRow";

const SurveyTable = () => {
    const [surveysArray, setSurveysArray] = useState([]);

    useEffect(() => {
        async function fetchSurveys() {
            const response = await post(
                process.env.NEXT_PUBLIC_GET_ALL_SURVEYS,
                {
                    token: localStorage.getItem("token"),
                }
            );
            setSurveysArray(response.surveys);
        }

        fetchSurveys();
    }, []);

    const handleDelete = (idToDelete) => {
        setSurveysArray((prevSurveys) =>
            prevSurveys.filter((survey) => survey.survey_id !== idToDelete)
        );
        let response = post(
            process.env.NEXT_PUBLIC_GET_ALL_SURVEYS + "/" + idToDelete,
            { token: localStorage.getItem("token") }
        );
    };

    return (
        <>
            <div className={styles.panel}>
                <table className={styles.table}>
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Название</th>
                            <th>Логин автора</th>
                            <th>Дата создания</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {surveysArray.map((survey) => (
                            <SurveyTableRow
                                key={survey.survey_id}
                                surveyID={survey.survey_id}
                                topic={survey.topic}
                                authorLogin={survey.login}
                                creationDate={
                                    survey.creation_time.slice(8, 10) +
                                    "." +
                                    survey.creation_time.slice(5, 7) +
                                    "." +
                                    survey.creation_time.slice(0, 4)
                                }
                                handleDelete={handleDelete}
                            />
                        ))}
                    </tbody>
                </table>
            </div>
        </>
    );
};

export default SurveyTable;
