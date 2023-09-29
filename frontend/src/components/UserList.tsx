import React, { useEffect, useState } from "react";
import User from "src/types/UserType";

const UserList: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    // Fetch users from your API and set them in the state
    // Example: fetchUsers().then((data) => setUsers(data));
  }, []);

  const handleEdit = (id: number) => {};

  const handleDelete = (id: number) => {};

  return (
    <div>
      <h2>User List</h2>
      <ul>
        {users.map((user) => (
          <li key={user.id}>
            {user.name}{" "}
            <button onClick={() => handleEdit(user.id)}>Edit</button>
            <button onClick={() => handleDelete(user.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default UserList;
