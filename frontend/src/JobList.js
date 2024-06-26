import React from 'react';
import JobItem from './JobItem';

function JobList({ jobs }) {
  return (
    <div>
      <h2>Job List</h2>
      <ul>
        {jobs.map((job, index) => (
          <JobItem key={index} job={job} />
        ))}
      </ul>
    </div>
  );
}

export default JobList;
