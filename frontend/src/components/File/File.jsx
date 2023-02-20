import React, { useState } from 'react';
import './File.css';

function File({ toggle, filename, content }) {

    return (
      <>
        <main>
          <h1 className={toggle}>{filename}</h1>
          <p className={toggle}>{content}</p>
        </main>
      </>
    );
}

export default File;

