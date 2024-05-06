import styles from "./page.module.css";
import Survey from "./Survey.jsx"
import Header from "@/components/Header";

export default function Mysurveys(children) {
    return (
        <>
            <Header />

            <div className={styles.header}>Мои опросы</div>

            <div className={styles.grid}>
                <div className={styles.cell}><Survey topic={"123"} /></div>
                <div className={styles.cell}><Survey topic={"222"} /></div>
                <div className={styles.cell}><Survey topic={"333"} /></div>
                <div className={styles.cell}><Survey topic={"5555555 555555555 5555555 555555 5555555 555555 55555555 555555555 55555555"} /></div>
            </div>
        </>
    );
}
