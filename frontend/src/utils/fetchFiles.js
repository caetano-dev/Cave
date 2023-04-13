async function fetchFiles(setData) {
  const savedData = JSON.parse(localStorage.getItem("data")) || { files: [] };

  if (navigator.onLine) {
    await fetch("http://localhost:3000/files")
      .then((response) => {
        if (response.status === 200) {
          return response.json();
        } else {
          throw new Error("Server error: " + response.status);
        }
      })
      .then((json) => {
        //save to data and local storage
        //TODO: compate with local storage to see if we use it or not. if is it different from local storage, we use local storage, else, we upload
        setData(json.files);
        localStorage.setItem("data", JSON.stringify(json.files)); //similar to data
      })
      .catch((error) => {
        console.error(error);
        setData(savedData);
      });
  } else {
    setData(savedData);
  }
};
export default fetchFiles;
