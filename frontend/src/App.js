import React, { useEffect, useState } from 'react';
import JobList from './JobList';
import NewJobForm from './NewJobForm';

function App() {
  const [jobs, setJobs] = useState([]);

  useEffect(() => {
    const fetchJobs = async () => {
      const response = await fetch('http://localhost:8080/jobs');
      const data = await response.json();
      setJobs(data);
    };
    
    fetchJobs();

    const ws = new WebSocket('ws://localhost:8080/ws');
    ws.onmessage = (event) => {
      const job = JSON.parse(event.data);
      setJobs((prevJobs) => {
        const updatedJobs = prevJobs.filter(j => j.name !== job.name);
        return [...updatedJobs, job];
      });
    };

    return () => ws.close();
  }, []);

  return (
    <div className="App">
      <h1>Job Scheduler</h1>
      <NewJobForm />
      <JobList jobs={jobs} />
    </div>
  );
}

export default App;
