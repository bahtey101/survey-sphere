"use client";

import styles from "./page.module.css";

const UserTableRow = ({ userID, login, role, surveysAmount, handleDelete }) => {
    return (
        <>
            <tr>
                <td>{userID}</td>
                <td>{login}</td>
                <td>{role}</td>
                <td>{surveysAmount}</td>
                <td>
                    <div className={styles.container_button_delete}>
                        <button
                            className={styles.button_delete}
                            onClick={() => handleDelete(userID)}
                        >
                            Удалить
                        </button>
                    </div>
                </td>
            </tr>
        </>
    );
};

export default UserTableRow;
