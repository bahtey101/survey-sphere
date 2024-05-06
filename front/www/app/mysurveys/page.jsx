import Header from "@/components/Header";
import SurveyWithImage from "./SurveyWithImage";

export default function Mysurveys(children) {
    return (
        <>
            <Header />
            <div style={{ margin: "100px 140px" }}>
                <p>
                    <b>Мои опросы</b>
                </p>
                <SurveyWithImage />
            </div>
        </>
    );
}
