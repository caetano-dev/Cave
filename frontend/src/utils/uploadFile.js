const uploadFile = async (file) => {
  if (navigator.onLine) {
    const host = "http://localhost:3000";
    const url = `${host}/files/create`;
    const formData = new FormData();
    formData.append("file", file);
    formData.append("filename", file.name);
    formData.append("tags", ""); // Update with actual tags value
    const requestParamns = {
      method: "POST",
      body: formData,
    };
    try {
      const response = await fetch(url, requestParamns);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      console.log("File upload successful:", data);
    } catch (error) {
      console.error("Error uploading file:", error);
    }
  }
};

export default uploadFile;
