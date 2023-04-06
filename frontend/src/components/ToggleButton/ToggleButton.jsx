import React from "react";

function ToggleButton({ toggle, toggleState }) {
  return (
    <button
      className={"sideBarButton " + toggle}
      title="Toggle sidebar"
      onClick={toggleState}
    >
      {toggle === "open" ? "«" : "»"}
    </button>
  );
}

export default ToggleButton;
