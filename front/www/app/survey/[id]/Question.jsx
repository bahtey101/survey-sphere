import React, { useState } from "react";
import styles from "@/app/newsurvey/page.module.css";
import Input from "@/components/Input";

const Question = ({ onChange, number, questionText }) => {
    const [text, setText] = useState("");

    const handleTextChange = (event) => {
        setText(event.target.value);
        onChange(event.target.value, number);
    };

    return (
        <>
            <div className={styles.question}>
                <label className={styles.question_label}>
                    {number + 1}. {questionText}
                </label>

                <Input isRequired={true} onChange={handleTextChange} />
            </div>
        </>
    );
};

export default Question;
