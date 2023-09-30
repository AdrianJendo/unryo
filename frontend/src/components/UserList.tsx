import React, { useEffect, useState } from "react";
import User from "src/types/UserType";
import "src/css/UserList.css";

interface UserListProps {
  users: User[];
  handleDelete: (id: number) => void;
}

const UserList: React.FC<UserListProps> = ({ users, handleDelete }) => {
  const handleEdit = (id: number) => {};

  return (
    <div className="user-list">
      <h2>User List</h2>
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            <div className="name-column">{user.name}</div>
            <div className="buttons-column">
              <button onClick={() => handleEdit(user.id)}>Edit</button>
              <button onClick={() => handleDelete(user.id)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserList;
