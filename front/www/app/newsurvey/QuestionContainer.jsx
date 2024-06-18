import Question from "./Question";
import NewQuestionButton from "./NewQuestionButton";
import styles from "./page.module.css";

const QuestionContainer = ({ questions, setQuestions }) => {
    const handleInputChange = (index, newData) => {
        const updatedComponents = [...questions];
        updatedComponents[index] = newData;
        setQuestions(updatedComponents);
    };

    const addComponent = () => {
        setQuestions([...questions, ""]);
    };

    return (
        <>
            {questions.map((_, index) => (
                <Question
                    key={index}
                    onInputChange={(newData) =>
                        handleInputChange(index, newData)
                    }
                    number={index}
                />
            ))}
            <div className={styles.container_new_question_button}>
                <NewQuestionButton onClick={addComponent} />
            </div>
        </>
    );
};

export default QuestionContainer;
