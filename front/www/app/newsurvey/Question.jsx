import React, { useState } from "react";
import styles from "./page.module.css";
import Input from "@/components/Input";

const Question = ({ onInputChange, number }) => {
    const [text, setText] = useState("");

    const handleTextChange = (event) => {
        setText(event.target.value);
        onInputChange(event.target.value);
    };

    return (
        <>
            <div className={styles.question}>
                <label className={styles.question_label}>
                    Вопрос {number + 1}
                </label>
                <Input isRequired={true} onChange={handleTextChange} />
            </div>
        </>
    );
};

export default Question;
