import React from "react";

export function Button({ message, children }) {
  return <button onClick={() => alert(message)}>{children}</button>;
}
