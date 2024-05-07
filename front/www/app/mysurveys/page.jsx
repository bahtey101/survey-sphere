import styles from "./page.module.css";
import Survey from "./Survey.jsx";
import Header from "@/components/Header";

export default function Mysurveys(children) {
    return (
        <>
            <Header />

            <div className={styles.grid}>
                <Survey topic={"123"} />
                <Survey topic={"222"} />
                <Survey topic={"333"} />
                <Survey
                    topic={
                        "5555555 555555555 5555555 555555 5555555 555555 55555555 555555555 55555555"
                    }
                />
                <Survey topic={"123"} />
                <Survey topic={"222"} />
                <Survey topic={"333"} />
            </div>
        </>
    );
}
