import "./PopUp.css"
function FilePopup({ position, onClickFunction, text }) {
  return (
    <div className="popUp" style={{ position: "fixed", left: position.x, top: position.y }}>
      <button onClick={onClickFunction}>{text}</button>
    </div>
  );
}

export default FilePopup;