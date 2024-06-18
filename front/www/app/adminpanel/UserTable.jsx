import styles from "./page.module.css";
import { useEffect, useState } from "react";
import { post } from "@/utils/fething";
import UserTableRow from "./UserTableRow";
import { deleteRequest } from "@/utils/fething";

const UserTable = () => {
    const [usersArray, setUsersArray] = useState([]);

    useEffect(() => {
        async function fetchUsers() {
            const response = await post(process.env.NEXT_PUBLIC_GET_ALL_USERS, {
                token: localStorage.getItem("token"),
            });
            setUsersArray(response.users);
        }

        fetchUsers();
    }, []);

    const handleDelete = (idToDelete) => {
        setUsersArray((prevUsers) =>
            prevUsers.filter((users) => users.id !== idToDelete)
        );
        let response = deleteRequest(
            process.env.NEXT_PUBLIC_GET_ALL_USERS + "/" + idToDelete,
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
                            <th>Логин</th>
                            <th>Роль</th>
                            <th>Кол-во опросов</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {usersArray.map((user) => (
                            <UserTableRow
                                key={user.id}
                                userID={user.id}
                                login={user.login}
                                role={user.role}
                                surveysAmount={user.surveys_number}
                                handleDelete={handleDelete}
                            />
                        ))}
                    </tbody>
                </table>
            </div>
        </>
    );
};

export default UserTable;
