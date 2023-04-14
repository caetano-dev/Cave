import fetchFiles from "./fetchFiles";
import uploadFile from "./uploadFile";

const syncFiles = async (setData) => {
  if (navigator.onLine) {
    const savedData = JSON.parse(localStorage.getItem("data")) || { files: [] };
    localStorage.setItem("data", JSON.stringify(savedData)); // save data instead of json
    const serverFiles = await fetchFiles();
    const unsyncedFiles = savedData.files.filter((file) => {
      return !serverFiles.files.some((serverFile) => {
        return (
          serverFile.FileInformation.id === file.FileInformation.id
        );
      });
    });

    for (const file of unsyncedFiles) {
      await uploadFile(file);
    }
  }
};

export default syncFiles;
