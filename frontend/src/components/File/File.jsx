import React, {useEffect, useState} from 'react';
import './File.css';

function File({ toggle, filename, id, content, setContent}) {
  const [textareaValue, setTextareaValue] = useState(content);

  function handleTextareaChange(event){
    setTextareaValue(event.target.value);
  }

  const editContent = async (id) => {
    console.log("making request to the server.")
    try {
      const response = await fetch("http://localhost:3000/fileEditContent", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ID: id,
          Content: textareaValue,
        }),
      });
    } catch (error) {
      console.error(error);
    }
  };


  useEffect(() => {
    function handleKeyDown(event){
      if(event.ctrlKey && event.key === 's'){
        event.preventDefault();
        setContent(textareaValue)
        editContent(id)
      }
    }
    document.addEventListener("keydown", handleKeyDown)
    return ()=>{
      document.removeEventListener('keydown', handleKeyDown)
    }
  },[textareaValue]);


  return (
    <main className={toggle}>
      <h1 className='title'>{filename}</h1>
      <textarea autofocus className='content' value={textareaValue} onChange={handleTextareaChange}>{content}</textarea>
    </main>
  );
}

export default File;

