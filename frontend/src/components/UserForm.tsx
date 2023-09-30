import React, { useEffect, useState } from "react";

interface UserFormProps {
  onSave: (user: string) => void;
}

const UserForm: React.FC<UserFormProps> = ({ onSave }) => {
  const [user, setUser] = useState("");
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUser(e.target.value);
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave(user);
    setUser("");
  };

  useEffect(() => {}, []);

  return (
    <div>
      <h2>User Form</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Name:
          <input type="text" name="name" value={user} onChange={handleChange} />
        </label>
        {/* Add other input fields for user properties here */}
        <button type="submit">Save</button>
      </form>
    </div>
  );
};

export default UserForm;
