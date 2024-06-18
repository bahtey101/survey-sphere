"use client";

import styles from "./page.module.css";

const SurveyTableRow = ({
    surveyID,
    topic,
    authorLogin,
    creationDate,
    handleDelete,
}) => {
    return (
        <>
            <tr>
                <td>{surveyID}</td>
                <td>{topic}</td>
                <td>{authorLogin}</td>
                <td>{creationDate}</td>
                <td>
                    <div className={styles.container_button_delete}>
                        <button
                            className={styles.button_delete}
                            onClick={() => handleDelete(surveyID)}
                        >
                            Удалить
                        </button>
                    </div>
                </td>
            </tr>
        </>
    );
};

export default SurveyTableRow;
