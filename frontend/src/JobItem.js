import React from 'react';

function JobItem({ job }) {
  let statusClass = '';

  switch (job.status) {
    case 'pending':
      statusClass = 'pending';
      break;
    case 'running':
      statusClass = 'running';
      break;
    case 'completed':
      statusClass = 'completed';
      break;
    default:
      statusClass = '';
  }

  return (
    <li className={`job-item ${statusClass}`}>
      {job.name} - {job.duration}ms - {job.status}
    </li>
  );
}

export default JobItem;
