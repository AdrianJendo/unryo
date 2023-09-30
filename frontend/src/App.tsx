import React, { useEffect, useState } from "react";
import "./App.css";
import User from "src/types/UserType";
import UserList from "./components/UserList";
import UserForm from "./components/UserForm";
import axios from "axios";

const BASE_URL = "http://localhost:5001/api/users";

function App() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    axios.get(BASE_URL).then((resp) => {
      const newUsers: User[] = resp.data;
      setUsers(newUsers);
    });
  }, []);

  const handleSaveUser = (name: string) => {
    axios
      .post(BASE_URL, {
        name,
      })
      .then((resp) => {
        const newUsers: User[] = resp.data;
        setUsers(newUsers);
      });
  };

  const handleDeleteUser = (id: number) => {
    axios.delete(`${BASE_URL}/${id}`).then((resp) => {
      const newUsers: User[] = resp.data;
      setUsers(newUsers);
    });
  };

  return (
    <div className="App">
      <UserForm onSave={handleSaveUser} />
      <UserList users={users} handleDelete={handleDeleteUser} />
    </div>
  );
}

export default App;
