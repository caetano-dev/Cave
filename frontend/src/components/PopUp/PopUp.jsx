function FilePopup({ position, onDelete }) {
  return (
    <div style={{ position: "fixed", left: position.x, top: position.y }}>
      <button onClick={onDelete}>Delete file</button>
    </div>
  );
}

export default FilePopup;