import React from "react";
import "./App.css";
import User from "src/types/UserType";
import UserList from "./components/UserList";
import UserForm from "./components/UserForm";

function App() {
  const handleSaveUser = (user: User) => {
    // Send a POST request to create a new user or a PUT request to update an existing user
    // Example: saveUser(user).then(() => refreshUserList());
  };

  return (
    <div className="App">
      <UserForm onSave={handleSaveUser} />
      <UserList />
    </div>
  );
}

export default App;
