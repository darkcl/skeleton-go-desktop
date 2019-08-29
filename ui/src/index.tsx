import * as React from "react";
import * as ReactDOM from "react-dom";

declare var external;
const render = () =>
  ReactDOM.render(
    <div>
      <p>Hello World!</p>
    </div>,
    document.getElementById("root")
  );

window.onclick = function(e) {
  const elem = e.target as Element;
  if (elem.localName === "a") {
    e.preventDefault();
    external.invoke("openlink: " + elem.getAttribute("href"));
  }
};

render();
