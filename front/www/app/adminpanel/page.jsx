"use client";

import Header from "@/components/Header";
import styles from "./page.module.css";
import Script from "next/script";
import Start from "./Start";
import UserTable from "./UserTable";
import SurveyTable from "./SurveyTable";
import { useState } from "react";

export default function AdminPanel(children) {
    const [state, setState] = useState(0);

    const chooseTable = () => {
        if (state == 0) {
            return <Start />;
        } else if (state == 1) {
            return <UserTable />;
        } else {
            return <SurveyTable />;
        }
    };

    const setUserTableState = () => {
        setState(1);
    };

    const setSurveyTableState = () => {
        setState(2);
    };

    return (
        <>
            <Header />
            <Script src={"/static/scripts/adminpanel/script.js"} />
            <div className={styles.container}>
                <div className={styles.panel}>
                    <div className={styles.panel_flex}>
                        <div className={styles.radio_button}>
                            <input
                                name="tableType"
                                type="radio"
                                id="users"
                                className={styles.radio_field}
                                onClick={setUserTableState}
                            />
                            <label
                                htmlFor="users"
                                className={styles.radio_label}
                            >
                                Пользователи
                            </label>
                        </div>
                        <div className={styles.radio_button}>
                            <input
                                name="tableType"
                                type="radio"
                                id="surveys"
                                className={styles.radio_field}
                                onClick={setSurveyTableState}
                            />
                            <label
                                htmlFor="surveys"
                                className={styles.radio_label}
                            >
                                Опросы
                            </label>
                        </div>
                        <input
                            type="search"
                            className={`${styles.flex_horizontal_split} ${styles.search_bar}`}
                        ></input>
                    </div>
                </div>

                {chooseTable()}
            </div>
        </>
    );
}
