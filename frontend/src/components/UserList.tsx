import React, { useEffect, useState } from "react";
import User from "src/types/UserType";
import "src/css/UserList.css";

interface UserListProps {
  users: User[];
  handleDelete: (id: number) => void;
  handleEdit: (id: number, name: string) => void;
}

const UserList: React.FC<UserListProps> = ({
  users,
  handleDelete,
  handleEdit,
}) => {
  const [editUser, setEditUser] = useState<number | null>(null);
  const [editUserValue, setEditUserValue] = useState<string>("");

  const toggleEdit = (user: User): void => {
    setEditUser(user.id);
    setEditUserValue(user.name);
  };
  const handleEditRequest = (): void => {
    if (editUser) {
      handleEdit(editUser, editUserValue);
      setEditUser(null);
      setEditUserValue("");
    }
  };

  return (
    <div className="user-list">
      <h2>User List</h2>
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            {user.id == editUser ? (
              <input
                className="name-column"
                type="text"
                value={editUserValue}
                onChange={(e) => setEditUserValue(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key == "Enter") {
                    handleEditRequest();
                  }
                }}
              />
            ) : (
              <div className="name-column">{user.name}</div>
            )}
            <div className="buttons-column">
              {user.id == editUser && (
                <button onClick={() => handleEditRequest()}>Submit</button>
              )}
              <button onClick={() => toggleEdit(user)}>
                {user.id == editUser ? "Cancel" : "Edit"}
              </button>
              <button onClick={() => handleDelete(user.id)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserList;
