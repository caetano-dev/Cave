import fetchFiles from "./fetchFiles";

const syncFiles = (setData) => {
  if (navigator.online) {
    const savedData = JSON.parse(localStorage.getItem("data")) || { files: [] };
    localStorage.setItem("data", JSON.stringify(json.files)); //similar to data
    const serverFiles = fetchFiles();
    if (savedData != serverFiles) {
      //TODO: upload savedData to server
    }
  }
};
export default syncFiles;
