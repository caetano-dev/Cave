import React, {useEffect, useState} from 'react';
import './File.css';

function File({ toggle, filename, id, content, setContent}) {

  function handleTextareaChange(event){
    setContent(event.target.value);
  }

  const editContent = async (id) => {
    try {
      const response = await fetch("http://localhost:3000/fileEditContent", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ID: id,
          Content: content,
        }),
      });
      if (!response.ok){
        throw new Error("Failed to edit file content")
      }
    } catch (error) {
      console.error(error);
      //handle error here. We can display a message to the user.
    }
  };

  useEffect(() => {
    function handleSaveShortcutKeyDown(event){
      if(event.ctrlKey && event.key === 's'){
        event.preventDefault();
        editContent(id)
        setContent(content)
        console.log("content updated: "+content)
      }
    }
    document.addEventListener("keydown", handleSaveShortcutKeyDown)
    return ()=>{
      document.removeEventListener('keydown', handleSaveShortcutKeyDown)
    }
  },[content]);

  return (
    <main className={toggle}>
      <h1 className='title'>{filename}</h1>
      <textarea className='content' onChange={handleTextareaChange} value={content}/>
    </main>
  );
}

export default File;

