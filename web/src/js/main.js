import React from "react";
import { createRoot } from "react-dom/client";
import { Button } from "./test";

function Test() {
  return (
    <div>
      <Button message="testing">Play test</Button>
      <Button message="puta">Play puta</Button>
    </div>
  );
}

const domNode = document.getElementById("button");
const root = createRoot(domNode);
root.render(<Test />);
