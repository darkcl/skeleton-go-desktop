import * as React from "react";
import * as ReactDOM from "react-dom";
import { IPCRenderer } from "./ipc";

declare var external;

declare global {
  interface Window {
    renderer: IPCRenderer;
  }
}

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

window.renderer = new IPCRenderer();

window.renderer.on("testing", (event, value) => {
  console.log("Receive from go backend");
  console.log(event);
  console.log(value);
});

render();
