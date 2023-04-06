import React from "react";
import createFile from "../../utils/createFile"

const CreateFileButton = ({ setData }) => {
  return <button className="createFileButton" onClick={() => createFile(setData)}>+</button>;
};

export default CreateFileButton;
