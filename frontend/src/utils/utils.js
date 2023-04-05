const fetchFiles = (setData) => {
  const savedData = JSON.parse(localStorage.getItem("data")) || { files: [] };
  console.log(savedData);
  if (!navigator.onLine) {
    setData(savedData.files);
    return;
  }

  fetch("http://localhost:3000/files")
    .then((response) => {
      if (response.status === 200) {
        return response.json();
      } else {
        throw new Error("Server error: " + response.status);
      }
    })
    .then((json) => {
      setData(json.files);
      localStorage.setItem("data", JSON.stringify(json));
    })
    .catch((error) => {
      console.error(error);
      setData(savedData.files);
    });
};

export default fetchFiles
