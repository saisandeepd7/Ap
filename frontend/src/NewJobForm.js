import React, { useState } from 'react';

function NewJobForm() {
  const [name, setName] = useState('');
  const [duration, setDuration] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();

    const job = {
      name,
      duration: parseInt(duration),
      status: 'pending',
    };

    const response = await fetch('http://localhost:8080/jobs', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(job),
    });

    if (response.ok) {
      setName('');
      setDuration('');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label>Name</label>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
      </div>
      <div>
        <label>Duration (in milliseconds)</label>
        <input
          type="number"
          value={duration}
          onChange={(e) => setDuration(e.target.value)}
        />
      </div>
      <button type="submit">Submit</button>
    </form>
  );
}

export default NewJobForm;
