import React from 'react';
import './File.css';

function File({ toggle, filename, content }) {

  return (
    <>
      <main className={toggle}>
        <h1 className='title'>{filename}</h1>
        <p className='content'>{content}</p>
      </main>
    </>
  );
}

export default File;

