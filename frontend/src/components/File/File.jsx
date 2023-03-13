import React from 'react';
import './File.css';

function File({ toggle, filename, content }) {

  return (
      <main className={toggle}>
        <h1 className='title'>{filename}</h1>
        <textarea autofocus className='content'>{content}</textarea>
      </main>
  );
}

export default File;

