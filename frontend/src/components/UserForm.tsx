import React, { useEffect, useState } from "react";
import User from "src/types/UserType";
import axios from "axios";

interface UserFormProps {
  onSave: (user: User) => void;
}

const UserForm: React.FC<UserFormProps> = ({ onSave }) => {
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    // setUser({ ...user, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // onSave(user);
    // setUser({ name: "" });
    axios.get("http://localhost:5001/api/users");
  };

  useEffect(() => {}, []);

  return (
    <div>
      <h2>User Form</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input
            type="text"
            name="name"
            value={"test"}
            onChange={handleChange}
          />
        </label>
        {/* Add other input fields for user properties here */}
        <button type="submit">Save</button>
      </form>
    </div>
  );
};

export default UserForm;
