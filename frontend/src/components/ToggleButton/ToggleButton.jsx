import React from "react";

function ToggleButton({toggle, toggleState}) {
  return (
    <button
      className={"sideBarButton " + toggle}
      title="Toggle sidebar"
      onClick={toggleState}
    >
      ☰
    </button>
  );
}

export default ToggleButton;
